// Package register provides utilities for dynamically registering
// Cobra commands based on API client structs using reflection.
package register

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/equinix/cli/internal/parser"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
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
			return executeMethod(cmd, service, method, builderMethod, args, description.ParameterDescriptions)
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
	ParameterDescriptions []parser.ParameterDescription
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
func executeMethod(cmd *cobra.Command, service interface{}, _ reflect.Method, builderMethod reflect.Method, args []string, paramDescriptions []parser.ParameterDescription) error {
	ctx := context.Background()
	serviceValue := reflect.ValueOf(service)

	// Build the arguments for the builder method
	builderArgs := []reflect.Value{reflect.ValueOf(ctx)}

	// Get builder method parameters (skip receiver at index 0 and context at index 1)
	builderType := builderMethod.Type
	for i := 2; i < builderType.NumIn(); i++ {
		paramType := builderType.In(i)
		paramName := getParamNameFromPosition(builderMethod, i, paramDescriptions)

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

	// Apply setter method values from expanded flags
	requestBuilder = applyFlagsToRequestBuilder(cmd, requestBuilder)

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

	// Get HTTP response (second return value) for additional context
	var httpResp *http.Response
	if !result[1].IsNil() {
		httpResp = result[1].Interface().(*http.Response)
	}

	// Check for errors (third return value)
	if !result[2].IsNil() {
		err := result[2].Interface().(error)
		errMsg := err.Error()

		// Add HTTP status code if available
		statusInfo := ""
		if httpResp != nil {
			statusInfo = fmt.Sprintf(" (HTTP %d %s)", httpResp.StatusCode, httpResp.Status)
		}

		// Provide helpful hints for common errors
		if strings.Contains(errMsg, "is required") && strings.Contains(errMsg, "must be specified") {
			// Extract field name if possible
			hint := "\n\nHint: This command requires additional fields in the request body."
			hint += "\nUse the --request flag with a JSON payload. For example:"
			hint += "\n  --request '{\"fieldName\": \"value\"}'"
			hint += "\n\nRun with --help to see available request body fields."
			return fmt.Errorf("API error%s: %w%s", statusInfo, err, hint)
		}

		// Handle unmarshal errors - these typically indicate the API returned an error response
		// in a different format than expected (e.g., error object instead of success array)
		if strings.Contains(errMsg, "cannot unmarshal") {
			// If we have an HTTP response with an error status code, this is an API error
			// not a parsing error with a successful response
			if httpResp != nil && httpResp.StatusCode >= 400 {
				// Extract just the status from the error message if it's cleaner
				cleanMsg := errMsg
				// The SDK sometimes includes "json: cannot unmarshal..." which isn't useful
				// when we know it's an API error response
				if httpResp.StatusCode == 401 {
					cleanMsg = "Authentication failed - invalid or missing credentials"
				} else if httpResp.StatusCode == 403 {
					cleanMsg = "Access forbidden - insufficient permissions"
				} else if httpResp.StatusCode == 404 {
					cleanMsg = "Resource not found"
				}
				
				hint := "\n\nHint: The API returned an error response."
				hint += "\nTry running with --debug to see the full HTTP request and response."
				return fmt.Errorf("API error%s: %s%s", statusInfo, cleanMsg, hint)
			}
			
			// Otherwise it's a real unmarshal error with a successful response
			hint := "\n\nHint: Failed to parse API response."
			if httpResp != nil {
				hint += fmt.Sprintf("\nHTTP Status: %d %s", httpResp.StatusCode, httpResp.Status)
			}
			hint += "\nThis might indicate an unexpected response format."
			hint += "\nTry running with --debug to see the raw HTTP request and response."
			return fmt.Errorf("API error%s: %w%s", statusInfo, err, hint)
		}

		return fmt.Errorf("API error%s: %w", statusInfo, err)
	}

	// Format and display the result (first return value)
	if !result[0].IsNil() {
		response := result[0].Interface()
		jsonBytes, err := json.MarshalIndent(response, "", "  ")
		if err != nil {
			return fmt.Errorf("failed to format response: %w", err)
		}
		fmt.Println(string(jsonBytes))
	} else {
		// Some endpoints might return no content (e.g., DELETE operations)
		if httpResp != nil && httpResp.StatusCode >= 200 && httpResp.StatusCode < 300 {
			fmt.Printf("Success (HTTP %d %s)\n", httpResp.StatusCode, httpResp.Status)
		}
	}

	return nil
}

// getParamNameFromPosition extracts parameter name from method signature
// It uses SDK descriptions when available, falling back to heuristics
func getParamNameFromPosition(method reflect.Method, position int, paramDescriptions []parser.ParameterDescription) string {
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

	// First, try to get parameter name from SDK descriptions
	// The descriptions are now in order, so we can match by position
	if len(paramDescriptions) > 0 {
		// Position 2+ are actual parameters (skip receiver at 0, ctx at 1)
		paramIndex := position - 2

		// Find the non-ctx parameter at this index
		actualParamIndex := 0
		for _, param := range paramDescriptions {
			if param.Name == "ctx" {
				continue
			}
			if actualParamIndex == paramIndex {
				// Found the matching parameter
				return camelToKebab(param.Name)
			}
			actualParamIndex++
		}
	}

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

// camelToKebab converts camelCase to kebab-case
func camelToKebab(s string) string {
	var result strings.Builder
	for i, r := range s {
		if i > 0 && r >= 'A' && r <= 'Z' {
			result.WriteRune('-')
		}
		result.WriteRune(r)
	}
	return strings.ToLower(result.String())
}

// getValueFromFlag retrieves a flag value and converts it to the target type
func getValueFromFlag(cmd *cobra.Command, flagName string, targetType reflect.Type) reflect.Value {
	// Handle pointer types
	isPtr := targetType.Kind() == reflect.Ptr
	if isPtr {
		targetType = targetType.Elem()
	}

	var result reflect.Value

	switch targetType.Kind() {
	case reflect.String:
		val, _ := cmd.Flags().GetString(flagName)
		// For string-based types (including enums), create a value of the target type
		newValue := reflect.New(targetType).Elem()
		newValue.SetString(val)
		result = newValue
	case reflect.Int, reflect.Int32, reflect.Int64:
		val, _ := cmd.Flags().GetInt(flagName)
		result = reflect.ValueOf(val).Convert(targetType)
	case reflect.Bool:
		val, _ := cmd.Flags().GetBool(flagName)
		result = reflect.ValueOf(val)
	case reflect.Struct:
		// Build struct from expanded flags
		result = buildStructFromFlags(cmd, flagName, targetType)
	default:
		// For complex types, try to get as string and parse
		val, _ := cmd.Flags().GetString(flagName)
		result = convertStringToType(val, targetType)
	}

	// If the original type was a pointer, wrap the result
	if isPtr && result.IsValid() {
		ptrValue := reflect.New(result.Type())
		ptrValue.Elem().Set(result)
		return ptrValue
	}

	return result
}

// buildStructFromFlags constructs a struct value from CLI flags
func buildStructFromFlags(cmd *cobra.Command, prefix string, structType reflect.Type) reflect.Value {
	structValue := reflect.New(structType).Elem()

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		
		// Skip unexported fields
		if !field.IsExported() {
			continue
		}

		// Get JSON tag for field name
		fieldName := field.Name
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" && jsonTag != "-" {
			parts := strings.Split(jsonTag, ",")
			if parts[0] != "" {
				fieldName = parts[0]
			}
		}

		// Convert to kebab-case for CLI flag name
		flagName := camelToKebab(fieldName)
		if prefix != "" {
			flagName = prefix + "-" + flagName
		}

		// Check if the flag was set
		if !cmd.Flags().Changed(flagName) {
			continue
		}

		// Get the field type
		fieldType := field.Type
		isPtr := fieldType.Kind() == reflect.Ptr
		if isPtr {
			fieldType = fieldType.Elem()
		}

		// Get value from flag based on type
		var fieldValue reflect.Value
		switch fieldType.Kind() {
		case reflect.String:
			val, _ := cmd.Flags().GetString(flagName)
			fieldValue = reflect.New(fieldType).Elem()
			fieldValue.SetString(val)
		case reflect.Int, reflect.Int32, reflect.Int64:
			val, _ := cmd.Flags().GetInt(flagName)
			fieldValue = reflect.ValueOf(val).Convert(fieldType)
		case reflect.Bool:
			val, _ := cmd.Flags().GetBool(flagName)
			fieldValue = reflect.ValueOf(val)
		case reflect.Float32, reflect.Float64:
			val, _ := cmd.Flags().GetFloat64(flagName)
			fieldValue = reflect.ValueOf(val).Convert(fieldType)
		case reflect.Struct:
			// Recursively build nested struct
			fieldValue = buildStructFromFlags(cmd, flagName, fieldType)
		case reflect.Slice:
			// Parse JSON array from string
			val, _ := cmd.Flags().GetString(flagName)
			if val != "" {
				fieldValue = parseJSONToType(val, fieldType)
			}
		default:
			// Try to parse as JSON
			val, _ := cmd.Flags().GetString(flagName)
			if val != "" {
				fieldValue = parseJSONToType(val, fieldType)
			}
		}

		if fieldValue.IsValid() {
			// Set the field
			if isPtr {
				ptrValue := reflect.New(fieldType)
				ptrValue.Elem().Set(fieldValue)
				structValue.Field(i).Set(ptrValue)
			} else {
				structValue.Field(i).Set(fieldValue)
			}
		}
	}

	return structValue
}

// parseJSONToType parses a JSON string into a reflect.Value of the target type
func parseJSONToType(jsonStr string, targetType reflect.Type) reflect.Value {
	// Create a new instance of the target type
	newValue := reflect.New(targetType)
	
	// Try to unmarshal JSON into it
	if err := json.Unmarshal([]byte(jsonStr), newValue.Interface()); err == nil {
		return newValue.Elem()
	}
	
	// If JSON parsing fails, return zero value
	return reflect.Zero(targetType)
}

// convertStringToType converts a string to the target type
func convertStringToType(value string, targetType reflect.Type) reflect.Value {
	// Handle pointer types
	if targetType.Kind() == reflect.Ptr {
		targetType = targetType.Elem()
	}

	switch targetType.Kind() {
	case reflect.String:
		// For string types or string-based types (like enums), create a value of the target type
		// This handles both plain strings and type-aliased strings
		newValue := reflect.New(targetType).Elem()
		newValue.SetString(value)
		return newValue
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

// applyFlagsToRequestBuilder applies flag values to request builder using setter methods
func applyFlagsToRequestBuilder(cmd *cobra.Command, requestBuilder reflect.Value) reflect.Value {
	builderType := requestBuilder.Type()

	// Iterate through all methods to find setters
	for i := 0; i < builderType.NumMethod(); i++ {
		method := builderType.Method(i)

		// Skip Execute and other non-setter methods
		if method.Name == "Execute" || !method.IsExported() {
			continue
		}

		// Check if this is a setter (1 param in, returns builder type)
		if method.Type.NumIn() != 2 || method.Type.NumOut() != 1 {
			continue
		}

		// Get the flag name for this setter
		flagName := camelToKebab(method.Name)
		
		// Get the parameter type for this setter
		paramType := method.Type.In(1)
		
		// Check if we have flags for this setter (either direct or expanded struct fields)
		if hasAnyFlagsWithPrefix(cmd, flagName) {
			// Build the value from flags
			var paramValue reflect.Value
			
			// Handle pointer types
			isPtr := paramType.Kind() == reflect.Ptr
			actualType := paramType
			if isPtr {
				actualType = paramType.Elem()
			}
			
			// Build value based on type
			if actualType.Kind() == reflect.Struct {
				// Build struct from expanded flags
				paramValue = buildStructFromFlags(cmd, flagName, actualType)
				if isPtr {
					ptrValue := reflect.New(actualType)
					ptrValue.Elem().Set(paramValue)
					paramValue = ptrValue
				}
			} else if cmd.Flags().Changed(flagName) {
				// Get simple value from flag
				paramValue = getValueFromFlag(cmd, flagName, paramType)
			}
			
			// Call the setter if we have a value
			if paramValue.IsValid() && !paramValue.IsZero() {
				methodFunc := requestBuilder.MethodByName(method.Name)
				if methodFunc.IsValid() {
					result := methodFunc.Call([]reflect.Value{paramValue})
					if len(result) > 0 {
						requestBuilder = result[0]
					}
				}
			}
		}
	}

	return requestBuilder
}

// hasAnyFlagsWithPrefix checks if any flags exist with the given prefix
func hasAnyFlagsWithPrefix(cmd *cobra.Command, prefix string) bool {
	hasFlags := false
	cmd.Flags().VisitAll(func(flag *pflag.Flag) {
		if flag.Changed && strings.HasPrefix(flag.Name, prefix) {
			hasFlags = true
		}
	})
	return hasFlags
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
func registerBuilderMethodFlags(cmd *cobra.Command, builderMethod reflect.Method, paramDescriptions []parser.ParameterDescription) {
	builderType := builderMethod.Type

	// Check if the return type has setter methods
	settersExpanded := false
	if builderType.NumOut() > 0 {
		returnType := builderType.Out(0)
		// Check if the return type has setter methods that suggest request body fields
		if returnType.Kind() == reflect.Struct || (returnType.Kind() == reflect.Ptr && returnType.Elem().Kind() == reflect.Struct) {
			actualType := returnType
			if returnType.Kind() == reflect.Ptr {
				actualType = returnType.Elem()
			}

			// Find setter methods on the request type and expand them as flags
			for i := 0; i < actualType.NumMethod(); i++ {
				method := actualType.Method(i)
				// Look for setter methods (single param, returns same type)
				if method.Type.NumIn() == 2 && method.Type.NumOut() == 1 && method.Name != "Execute" {
					// This looks like a setter
					paramType := method.Type.In(1)
					paramName := camelToKebab(method.Name)
					
					// Expand this setter parameter into flags
					addBuilderFlagForType(cmd, paramName, paramType, builderMethod, fmt.Sprintf("%s field", paramName))
					settersExpanded = true
				}
			}
		}
	}

	// Only add --request flag for additional fields if we expanded setters
	if settersExpanded {
		cmd.Flags().String("request", "", "JSON payload for additional optional fields not exposed as flags")
	} else {
		cmd.Flags().String("request", "", "JSON payload for request body fields")
	}

	// Iterate through builder method parameters (skip receiver at index 0, context at index 1)
	for i := 2; i < builderType.NumIn(); i++ {
		paramType := builderType.In(i)
		paramName := getParamNameFromPosition(builderMethod, i, paramDescriptions)

		// Check if flag already exists (to avoid redefinition)
		if cmd.Flags().Lookup(paramName) != nil {
			continue
		}

		// Get description from SDK if available
		var paramDesc string
		// Find matching parameter description
		paramIndex := i - 2
		actualParamIndex := 0
		for _, param := range paramDescriptions {
			if param.Name == "ctx" {
				continue
			}
			if actualParamIndex == paramIndex {
				paramDesc = param.Description
				break
			}
			actualParamIndex++
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
	case reflect.Struct:
		// For struct types, expand fields as individual flags with prefix
		expandStructFields(cmd, paramName, paramType, isOptional)
	default:
		// For complex types, accept as string (will be parsed later)
		cmd.Flags().String(paramName, "", fmt.Sprintf("%s (JSON or string)", description))
		if !isOptional {
			_ = cmd.MarkFlagRequired(paramName)
		}
	}
}

// getFieldDescription looks up field description from SDK documentation
func getFieldDescription(typeName string, fieldName string) string {
	if descriptionsData == nil {
		return ""
	}
	
	// Look in global types first
	if typeDesc, ok := descriptionsData.Types[typeName]; ok {
		if fieldDesc, ok := typeDesc.Fields[fieldName]; ok {
			return fieldDesc.Description
		}
	}
	
	// Also check in service-specific types
	for _, service := range descriptionsData.Services {
		if typeDesc, ok := service.Types[typeName]; ok {
			if fieldDesc, ok := typeDesc.Fields[fieldName]; ok {
				return fieldDesc.Description
			}
		}
	}
	
	return ""
}

// expandStructFields recursively expands struct fields into CLI flags
func expandStructFields(cmd *cobra.Command, prefix string, structType reflect.Type, parentOptional bool) {
	// Limit expansion depth to avoid excessive flags
	const maxDepth = 2
	expandStructFieldsWithDepth(cmd, prefix, structType, parentOptional, 0, maxDepth)
}

// expandStructFieldsWithDepth expands struct fields with depth control
func expandStructFieldsWithDepth(cmd *cobra.Command, prefix string, structType reflect.Type, parentOptional bool, depth int, maxDepth int) {
	if depth >= maxDepth {
		// Stop recursion, just add a JSON flag for the whole struct
		description := fmt.Sprintf("%s (JSON)", prefix)
		cmd.Flags().String(prefix, "", description)
		return
	}

	for i := 0; i < structType.NumField(); i++ {
		field := structType.Field(i)
		
		// Skip unexported fields
		if !field.IsExported() {
			continue
		}

		// Get JSON tag for field name
		fieldName := field.Name
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" && jsonTag != "-" {
			// Use JSON tag name (before the comma)
			parts := strings.Split(jsonTag, ",")
			if parts[0] != "" {
				fieldName = parts[0]
			}
		}

		// Convert to kebab-case for CLI
		flagName := camelToKebab(fieldName)
		if prefix != "" {
			flagName = prefix + "-" + flagName
		}

		// Check if flag already exists
		if cmd.Flags().Lookup(flagName) != nil {
			continue
		}

		// Determine if field is optional (pointer or has omitempty)
		fieldType := field.Type
		isOptional := parentOptional || fieldType.Kind() == reflect.Ptr || strings.Contains(jsonTag, "omitempty")
		
		if fieldType.Kind() == reflect.Ptr {
			fieldType = fieldType.Elem()
		}

		// Get field description from SDK documentation
		typeName := structType.Name()
		sdkFieldDesc := getFieldDescription(typeName, field.Name)
		
		fieldDesc := sdkFieldDesc
		if fieldDesc == "" {
			// Fallback to using the flag name
			fieldDesc = flagName
		}
		
		if !isOptional && fieldType.Kind() != reflect.Bool {
			fieldDesc = fmt.Sprintf("%s (required)", fieldDesc)
		}

		// Add flag based on field type
		switch fieldType.Kind() {
		case reflect.String:
			cmd.Flags().String(flagName, "", fieldDesc)
			if !isOptional {
				_ = cmd.MarkFlagRequired(flagName)
			}
		case reflect.Int, reflect.Int32, reflect.Int64:
			cmd.Flags().Int(flagName, 0, fieldDesc)
			if !isOptional {
				_ = cmd.MarkFlagRequired(flagName)
			}
		case reflect.Bool:
			cmd.Flags().Bool(flagName, false, fieldDesc)
		case reflect.Float32, reflect.Float64:
			cmd.Flags().Float64(flagName, 0.0, fieldDesc)
		case reflect.Slice:
			// For slices, accept JSON array as string
			cmd.Flags().String(flagName, "", fmt.Sprintf("%s (JSON array)", fieldDesc))
			if !isOptional {
				_ = cmd.MarkFlagRequired(flagName)
			}
		case reflect.Struct:
			// Recursively expand nested structs
			expandStructFieldsWithDepth(cmd, flagName, fieldType, isOptional, depth+1, maxDepth)
		default:
			// For other complex types, accept as JSON string
			cmd.Flags().String(flagName, "", fmt.Sprintf("%s (JSON)", fieldDesc))
			if !isOptional {
				_ = cmd.MarkFlagRequired(flagName)
			}
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
