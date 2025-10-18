// Package register provides utilities for dynamically registering
// Cobra commands based on API client structs using reflection.
package register

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/equinix/cli/internal/parser"
	"github.com/spf13/cobra"
)

// APIClientInterface represents any API client that can be used for command registration
type APIClientInterface interface{}

// ServiceInfo holds information about a registered API service
type ServiceInfo struct {
	Name        string
	ServiceType reflect.Type
	Description string
}

// descriptionsData holds the loaded SDK descriptions
var descriptionsData *parser.SDKDescriptions

// LoadDescriptions loads SDK descriptions from embedded JSON data
func LoadDescriptions(jsonData []byte) error {
	var descriptions parser.SDKDescriptions
	if err := json.Unmarshal(jsonData, &descriptions); err != nil {
		return fmt.Errorf("failed to unmarshal descriptions: %w", err)
	}
	descriptionsData = &descriptions
	return nil
}

// ServiceCommands dynamically creates Cobra commands for each API service
// in the provided API client using reflection. It creates a subcommand for each
// service field that ends with "Api" or "ApiService".
func ServiceCommands(parentCmd *cobra.Command, client APIClientInterface, _ string) error {
	if client == nil {
		return fmt.Errorf("client cannot be nil")
	}

	clientValue := reflect.ValueOf(client)
	if clientValue.Kind() == reflect.Ptr {
		clientValue = clientValue.Elem()
	}

	clientType := clientValue.Type()
	if clientType.Kind() != reflect.Struct {
		return fmt.Errorf("client must be a struct or pointer to struct")
	}

	// Iterate through all fields in the client struct
	for i := 0; i < clientType.NumField(); i++ {
		field := clientType.Field(i)
		fieldValue := clientValue.Field(i)

		// Skip unexported fields
		if !field.IsExported() {
			continue
		}

		// Check if the field name ends with "Api" or "ApiService"
		fieldName := field.Name
		if !strings.HasSuffix(fieldName, "Api") && !strings.HasSuffix(fieldName, "ApiService") {
			continue
		}

		// Skip nil fields
		if fieldValue.IsNil() {
			continue
		}

		// Create a command for this service
		serviceName := extractServiceName(fieldName)
		serviceCmd := createServiceCommand(serviceName, fieldValue.Interface())

		parentCmd.AddCommand(serviceCmd)
	}

	return nil
}

// extractServiceName converts an API service field name to a CLI-friendly name
// e.g., "ConnectionsApi" -> "connections", "ServiceProfilesApiService" -> "service-profiles"
func extractServiceName(fieldName string) string {
	// Remove "Api" or "ApiService" suffix
	name := strings.TrimSuffix(fieldName, "ApiService")
	name = strings.TrimSuffix(name, "Api")

	// Convert camelCase to kebab-case
	var result strings.Builder
	for i, r := range name {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteRune('-')
		}
		result.WriteRune(r)
	}

	return strings.ToLower(result.String())
}

// createServiceCommand creates a Cobra command for a specific API service
func createServiceCommand(serviceName string, service interface{}) *cobra.Command {
	cmd := &cobra.Command{
		Use:   serviceName,
		Short: fmt.Sprintf("Manage %s resources", serviceName),
		Long:  fmt.Sprintf("Commands for managing %s resources in the API", serviceName),
	}

	// Register methods as subcommands
	registerServiceMethods(cmd, service, serviceName)

	return cmd
}

// registerServiceMethods inspects the service and registers its methods as subcommands
func registerServiceMethods(parentCmd *cobra.Command, service interface{}, _ string) {
	serviceType := reflect.TypeOf(service)

	// Get all methods of the service pointer
	for i := 0; i < serviceType.NumMethod(); i++ {
		method := serviceType.Method(i)

		// Skip unexported methods
		if !method.IsExported() {
			continue
		}

		// Only register methods that end with "Execute" as these are the actual API execution methods
		// Methods without "Execute" are request builders that return ApiRequest structs
		if !strings.HasSuffix(method.Name, "Execute") {
			continue
		}

		// Get the corresponding builder method (without Execute suffix)
		builderMethodName := strings.TrimSuffix(method.Name, "Execute")
		builderMethod, hasBuilder := serviceType.MethodByName(builderMethodName)
		if !hasBuilder {
			continue
		}

		// Create a command for this method
		methodName := extractMethodName(method.Name)
		methodCmd := createMethodCommand(methodName, method.Name, service, method, builderMethod)

		parentCmd.AddCommand(methodCmd)
	}
}

// extractMethodName converts a method name to a CLI-friendly name
// e.g., "GetConnectionByUuidExecute" -> "get-connection-by-uuid"
func extractMethodName(methodName string) string {
	// Remove the "Execute" suffix first
	methodName = strings.TrimSuffix(methodName, "Execute")

	// Convert the whole method name to kebab-case for uniqueness
	var result strings.Builder
	for i, r := range methodName {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteRune('-')
		}
		result.WriteRune(r)
	}

	return strings.ToLower(result.String())
}

// createMethodCommand creates a Cobra command for a specific API method
func createMethodCommand(methodName, originalMethodName string, service interface{}, method reflect.Method, builderMethod reflect.Method) *cobra.Command {
	// Extract description from loaded SDK descriptions or generate default
	description := extractMethodDescription(service, originalMethodName, builderMethod)

	cmd := &cobra.Command{
		Use:   methodName,
		Short: description.Short,
		Long:  description.Long,
		RunE: func(cmd *cobra.Command, args []string) error {
			return executeMethod(cmd, service, method, builderMethod, args)
		},
	}

	// Generate flags based on builder method parameters (not Execute method)
	// The builder method has the actual request parameters we need
	registerBuilderMethodFlags(cmd, builderMethod, description.ParameterDescriptions)

	return cmd
}

// MethodDescription holds the description extracted from SDK documentation
type MethodDescription struct {
	Short                 string
	Long                  string
	ParameterDescriptions map[string]string
}

// extractMethodDescription attempts to extract description from loaded SDK descriptions
func extractMethodDescription(service interface{}, methodName string, _ reflect.Method) MethodDescription {
	// Remove "Execute" suffix to get the builder method name
	builderMethodName := strings.TrimSuffix(methodName, "Execute")

	// Try to get description from loaded SDK descriptions
	if descriptionsData != nil {
		// Get service name from the service type
		serviceName := getServiceNameFromType(service)

		if serviceDesc, exists := descriptionsData.Services[serviceName]; exists {
			if methodDesc, exists := serviceDesc.Methods[builderMethodName]; exists {
				cmdName := extractMethodName(builderMethodName)

				shortDesc := methodDesc.ShortDescription
				if shortDesc == "" {
					shortDesc = fmt.Sprintf("Execute %s operation", cmdName)
				}

				longDesc := methodDesc.LongDescription
				if longDesc == "" {
					longDesc = shortDesc
				}

				// Add usage hint to long description
				longDesc = fmt.Sprintf("%s\n\nUse --request flag to provide optional JSON payload fields.", longDesc)

				return MethodDescription{
					Short:                 shortDesc,
					Long:                  longDesc,
					ParameterDescriptions: methodDesc.Parameters,
				}
			}
		}
	}

	// Fallback to generating description
	cmdName := extractMethodName(builderMethodName)
	return MethodDescription{
		Short: fmt.Sprintf("Execute %s operation", cmdName),
		Long: fmt.Sprintf(`Execute the %s operation on this service.

Use --request flag to provide a JSON payload for the request body.
Example: --request '{"field":"value"}'

The command accepts parameters based on the SDK method signature.`, cmdName),
		ParameterDescriptions: nil,
	}
}

// getServiceNameFromType extracts the service name from the service interface type
func getServiceNameFromType(service interface{}) string {
	serviceType := reflect.TypeOf(service)
	typeName := serviceType.String()

	// Remove package path and pointer indicator
	parts := strings.Split(typeName, ".")
	if len(parts) > 1 {
		typeName = parts[len(parts)-1]
	}
	typeName = strings.TrimPrefix(typeName, "*")

	// Remove "ApiService" suffix
	typeName = strings.TrimSuffix(typeName, "ApiService")

	// Convert to kebab-case
	var result strings.Builder
	for i, r := range typeName {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteRune('-')
		}
		result.WriteRune(r)
	}

	return strings.ToLower(result.String())
}

// executeMethod performs the actual API call using reflection
func executeMethod(cmd *cobra.Command, service interface{}, _ reflect.Method, builderMethod reflect.Method, args []string) error {
	ctx := context.Background()
	serviceValue := reflect.ValueOf(service)

	// Build the arguments for the builder method
	builderArgs := []reflect.Value{reflect.ValueOf(ctx)}

	// Get builder method parameters (skip receiver at index 0 and context at index 1)
	builderType := builderMethod.Type
	for i := 2; i < builderType.NumIn(); i++ {
		paramType := builderType.In(i)
		paramName := getParamNameFromPosition(builderMethod, i)

		// Try to get the value from flags or args
		var paramValue reflect.Value
		if cmd.Flags().Changed(paramName) {
			// Get value from flag
			paramValue = getValueFromFlag(cmd, paramName, paramType)
		} else if i-2 < len(args) {
			// Get value from positional argument
			paramValue = convertStringToType(args[i-2], paramType)
		} else {
			return fmt.Errorf("missing required parameter: %s", paramName)
		}

		builderArgs = append(builderArgs, paramValue)
	}

	// Call the builder method with arguments
	builderResult := builderMethod.Func.Call(append([]reflect.Value{serviceValue}, builderArgs...))
	if len(builderResult) == 0 {
		return fmt.Errorf("builder method returned no value")
	}

	requestBuilder := builderResult[0]

	// Check if --request flag was provided for JSON payload to set optional fields
	requestJSON, _ := cmd.Flags().GetString("request")
	if requestJSON != "" {
		// Parse JSON and populate request struct fields using setter methods
		var jsonData map[string]interface{}
		if err := json.Unmarshal([]byte(requestJSON), &jsonData); err != nil {
			return fmt.Errorf("failed to parse JSON request: %w", err)
		}

		// Try to call setter methods on the request builder
		requestBuilder = applyJSONToRequestBuilder(requestBuilder, jsonData)
	}

	// Call the Execute method on the request builder
	executeMethod := requestBuilder.MethodByName("Execute")
	if !executeMethod.IsValid() {
		return fmt.Errorf("could not find Execute method on request builder")
	}

	result := executeMethod.Call([]reflect.Value{})
	if len(result) < 3 {
		return fmt.Errorf("unexpected result from Execute method")
	}

	// Check for errors (third return value)
	if !result[2].IsNil() {
		err := result[2].Interface().(error)
		return fmt.Errorf("API error: %w", err)
	}

	// Format and display the result (first return value)
	if !result[0].IsNil() {
		response := result[0].Interface()
		jsonBytes, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to format response: %w", err)
		}
		fmt.Println(string(jsonBytes))
	}

	return nil
}

// getParamNameFromPosition extracts parameter name from method signature
func getParamNameFromPosition(method reflect.Method, position int) string {
	// For now, use a simple naming convention based on position
	// In a real implementation, we might parse the method name or use Go AST
	paramType := method.Type.In(position)
	typeName := paramType.String()

	// Remove package paths
	parts := strings.Split(typeName, ".")
	if len(parts) > 1 {
		typeName = parts[len(parts)-1]
	}

	// Remove pointer indicators
	typeName = strings.TrimPrefix(typeName, "*")

	// Convert to kebab-case
	var result strings.Builder
	for i, r := range typeName {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteRune('-')
		}
		result.WriteRune(r)
	}

	name := strings.ToLower(result.String())

	// Handle common parameter names by examining the method name
	methodName := method.Name
	lowerMethod := strings.ToLower(methodName)

	// For string types, try to infer from method name
	if name == "string" || strings.Contains(typeName, "string") {
		// Try to extract resource type from method name
		if strings.Contains(lowerMethod, "connection") {
			return "connection-id"
		}
		if strings.Contains(lowerMethod, "port") {
			return "port-id"
		}
		if strings.Contains(lowerMethod, "router") {
			return "router-id"
		}
		if strings.Contains(lowerMethod, "network") {
			return "network-id"
		}
		if strings.Contains(lowerMethod, "profile") {
			return "profile-id"
		}
		if strings.Contains(lowerMethod, "token") {
			return "token-id"
		}
		if strings.Contains(lowerMethod, "subscription") {
			return "subscription-id"
		}
		if strings.Contains(lowerMethod, "stream") {
			return "stream-id"
		}
		if strings.Contains(lowerMethod, "rule") {
			return "rule-id"
		}
		if strings.Contains(lowerMethod, "filter") {
			return "filter-id"
		}
		if strings.Contains(lowerMethod, "protocol") {
			return "protocol-id"
		}
		if strings.Contains(lowerMethod, "aggregation") {
			return "aggregation-id"
		}

		// Generic UUID/ID handling
		if strings.Contains(methodName, "ByUuid") || strings.Contains(methodName, "ByUUID") {
			return "uuid"
		}
		if strings.Contains(methodName, "ById") || strings.Contains(methodName, "ByID") {
			return "id"
		}

		// Default for string
		return fmt.Sprintf("param-%d", position-1)
	}

	return name
}

// getValueFromFlag retrieves a flag value and converts it to the target type
func getValueFromFlag(cmd *cobra.Command, flagName string, targetType reflect.Type) reflect.Value {
	// Handle pointer types
	if targetType.Kind() == reflect.Ptr {
		targetType = targetType.Elem()
	}

	switch targetType.Kind() {
	case reflect.String:
		val, _ := cmd.Flags().GetString(flagName)
		return reflect.ValueOf(val)
	case reflect.Int, reflect.Int32, reflect.Int64:
		val, _ := cmd.Flags().GetInt(flagName)
		return reflect.ValueOf(val).Convert(targetType)
	case reflect.Bool:
		val, _ := cmd.Flags().GetBool(flagName)
		return reflect.ValueOf(val)
	default:
		// For complex types, try to get as string and parse
		val, _ := cmd.Flags().GetString(flagName)
		return convertStringToType(val, targetType)
	}
}

// convertStringToType converts a string to the target type
func convertStringToType(value string, targetType reflect.Type) reflect.Value {
	// Handle pointer types
	if targetType.Kind() == reflect.Ptr {
		targetType = targetType.Elem()
	}

	switch targetType.Kind() {
	case reflect.String:
		return reflect.ValueOf(value)
	case reflect.Int, reflect.Int32, reflect.Int64:
		// Try to parse as int
		if val, err := strconv.ParseInt(value, 10, 64); err == nil {
			return reflect.ValueOf(val).Convert(targetType)
		}
	case reflect.Bool:
		if val, err := strconv.ParseBool(value); err == nil {
			return reflect.ValueOf(val)
		}
	}

	// Default: return zero value
	return reflect.Zero(targetType)
}

// applyJSONToRequestBuilder applies JSON data to request builder using setter methods
func applyJSONToRequestBuilder(requestBuilder reflect.Value, jsonData map[string]interface{}) reflect.Value {
	builderType := requestBuilder.Type()

	// Iterate through all methods to find setters
	for i := 0; i < builderType.NumMethod(); i++ {
		method := builderType.Method(i)

		// Skip Execute and other non-setter methods
		if method.Name == "Execute" || !method.IsExported() {
			continue
		}

		// Check if we have a value for this setter in the JSON
		// Try both exact match and lowercase match
		for jsonKey, jsonValue := range jsonData {
			if strings.EqualFold(jsonKey, method.Name) {
				// Try to call the setter method
				methodFunc := requestBuilder.MethodByName(method.Name)
				if methodFunc.IsValid() && methodFunc.Type().NumIn() == 1 {
					// Convert JSON value to the expected type
					paramType := methodFunc.Type().In(0)
					paramValue := convertJSONValueToType(jsonValue, paramType)
					if paramValue.IsValid() {
						// Call the setter method
						result := methodFunc.Call([]reflect.Value{paramValue})
						if len(result) > 0 {
							requestBuilder = result[0]
						}
					}
				}
			}
		}
	}

	return requestBuilder
}

// convertJSONValueToType converts a JSON value to a reflect.Value of the target type
func convertJSONValueToType(jsonValue interface{}, targetType reflect.Type) reflect.Value {
	// Handle pointer types
	if targetType.Kind() == reflect.Ptr {
		targetType = targetType.Elem()
	}

	// Try direct type assertion first
	jsonValueType := reflect.TypeOf(jsonValue)
	if jsonValueType != nil && jsonValueType.AssignableTo(targetType) {
		return reflect.ValueOf(jsonValue)
	}

	// Convert based on target type
	switch targetType.Kind() {
	case reflect.String:
		return reflect.ValueOf(fmt.Sprintf("%v", jsonValue))
	case reflect.Int, reflect.Int32, reflect.Int64:
		if val, ok := jsonValue.(float64); ok {
			return reflect.ValueOf(int64(val)).Convert(targetType)
		}
	case reflect.Bool:
		if val, ok := jsonValue.(bool); ok {
			return reflect.ValueOf(val)
		}
	case reflect.Struct, reflect.Map, reflect.Slice:
		// For complex types, marshal to JSON and unmarshal to target type
		jsonBytes, err := json.Marshal(jsonValue)
		if err == nil {
			newValue := reflect.New(targetType)
			if err := json.Unmarshal(jsonBytes, newValue.Interface()); err == nil {
				return newValue.Elem()
			}
		}
	}

	return reflect.Value{}
}

// registerBuilderMethodFlags creates flags based on the builder method parameters
func registerBuilderMethodFlags(cmd *cobra.Command, builderMethod reflect.Method, paramDescriptions map[string]string) {
	builderType := builderMethod.Type

	// Always add --request flag for JSON payload
	cmd.Flags().String("request", "", "Raw JSON payload for optional request fields")

	// Iterate through builder method parameters (skip receiver at index 0, context at index 1)
	for i := 2; i < builderType.NumIn(); i++ {
		paramType := builderType.In(i)
		paramName := getParamNameFromPosition(builderMethod, i)

		// Check if flag already exists (to avoid redefinition)
		if cmd.Flags().Lookup(paramName) != nil {
			continue
		}

		// Get description from SDK if available
		var paramDesc string
		// Try to find a matching description
		for key, desc := range paramDescriptions {
			if strings.Contains(strings.ToLower(key), strings.ToLower(paramName)) ||
				strings.Contains(strings.ToLower(desc), strings.ToLower(paramName)) {
				paramDesc = desc
				break
			}
		}

		// Add flags based on parameter type
		addBuilderFlagForType(cmd, paramName, paramType, builderMethod, paramDesc)
	}
}

// addBuilderFlagForType adds an appropriate flag for a builder method parameter
func addBuilderFlagForType(cmd *cobra.Command, paramName string, paramType reflect.Type, _ reflect.Method, sdkDescription string) {
	// Handle pointer types
	isOptional := paramType.Kind() == reflect.Ptr
	if isOptional {
		paramType = paramType.Elem()
	}

	// Create description based on SDK description or default
	var description string
	if sdkDescription != "" {
		description = sdkDescription
		if !isOptional {
			description = fmt.Sprintf("%s (required)", description)
		}
	} else {
		description = fmt.Sprintf("%s (required)", paramName)
		if isOptional {
			description = fmt.Sprintf("%s (optional)", paramName)
		}
	}

	// Add flags based on type kind
	switch paramType.Kind() {
	case reflect.String:
		cmd.Flags().String(paramName, "", description)
		if !isOptional {
			_ = cmd.MarkFlagRequired(paramName)
		}
	case reflect.Int, reflect.Int32, reflect.Int64:
		cmd.Flags().Int(paramName, 0, description)
		if !isOptional {
			_ = cmd.MarkFlagRequired(paramName)
		}
	case reflect.Bool:
		cmd.Flags().Bool(paramName, false, description)
	case reflect.Float32, reflect.Float64:
		cmd.Flags().Float64(paramName, 0.0, description)
	default:
		// For complex types, accept as string (will be parsed later)
		cmd.Flags().String(paramName, "", fmt.Sprintf("%s (JSON or string)", description))
		if !isOptional {
			_ = cmd.MarkFlagRequired(paramName)
		}
	}
}

// GetServiceList returns a list of all services available in the client
func GetServiceList(client APIClientInterface) ([]ServiceInfo, error) {
	if client == nil {
		return nil, fmt.Errorf("client cannot be nil")
	}

	clientValue := reflect.ValueOf(client)
	if clientValue.Kind() == reflect.Ptr {
		clientValue = clientValue.Elem()
	}

	clientType := clientValue.Type()
	if clientType.Kind() != reflect.Struct {
		return nil, fmt.Errorf("client must be a struct or pointer to struct")
	}

	var services []ServiceInfo

	// Iterate through all fields in the client struct
	for i := 0; i < clientType.NumField(); i++ {
		field := clientType.Field(i)
		fieldValue := clientValue.Field(i)

		// Skip unexported fields
		if !field.IsExported() {
			continue
		}

		// Check if the field name ends with "Api" or "ApiService"
		fieldName := field.Name
		if !strings.HasSuffix(fieldName, "Api") && !strings.HasSuffix(fieldName, "ApiService") {
			continue
		}

		// Skip nil fields
		if fieldValue.IsNil() {
			continue
		}

		serviceName := extractServiceName(fieldName)
		services = append(services, ServiceInfo{
			Name:        serviceName,
			ServiceType: field.Type,
			Description: fmt.Sprintf("%s API service", serviceName),
		})
	}

	return services, nil
}

// contextKey is used for context values
type contextKey string

// ContextWithClient adds a client to the context
func ContextWithClient(ctx context.Context, client interface{}) context.Context {
	return context.WithValue(ctx, contextKey("client"), client)
}

// ClientFromContext retrieves a client from the context
func ClientFromContext(ctx context.Context) (interface{}, bool) {
	client := ctx.Value(contextKey("client"))
	if client == nil {
		return nil, false
	}
	return client, true
}
