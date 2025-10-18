// Package register provides utilities for dynamically registering
// Cobra commands based on API client structs using reflection.
package register

import (
	"context"
	"fmt"
	"reflect"
	"strings"

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

		// Create a command for this method
		methodName := extractMethodName(method.Name)
		methodCmd := createMethodCommand(methodName, method.Name, service, method)

		parentCmd.AddCommand(methodCmd)
	}
}

// extractMethodName converts a method name to a CLI-friendly name
// e.g., "GetConnectionByUuid" -> "get-connection-by-uuid"
func extractMethodName(methodName string) string {
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
func createMethodCommand(methodName, originalMethodName string, _ interface{}, method reflect.Method) *cobra.Command {
	cmd := &cobra.Command{
		Use:   methodName,
		Short: fmt.Sprintf("Execute %s operation", methodName),
		Long:  fmt.Sprintf("Execute the %s operation on this service", originalMethodName),
		Run: func(_ *cobra.Command, _ []string) {
			// This is a placeholder implementation
			// In a real implementation, this would:
			// 1. Parse flags to get method parameters
			// 2. Call the method using reflection
			// 3. Format and display the result
			fmt.Printf("Executing %s (placeholder - to be implemented)\n", originalMethodName)
			fmt.Println("This command will be fully implemented in future iterations")
		},
	}

	// Generate flags based on method parameters
	registerMethodFlags(cmd, method)

	return cmd
}

// registerMethodFlags inspects a method's signature and adds corresponding flags to the command
func registerMethodFlags(cmd *cobra.Command, method reflect.Method) {
	methodType := method.Type

	// Iterate through method parameters (skip receiver at index 0)
	for i := 1; i < methodType.NumIn(); i++ {
		paramType := methodType.In(i)
		paramName := generateParamName(i, paramType)

		// Add flags based on parameter type
		addFlagForType(cmd, paramName, paramType, i)
	}
}

// generateParamName creates a CLI-friendly parameter name
func generateParamName(index int, paramType reflect.Type) string {
	// Try to extract meaningful name from the type
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

	flagName := strings.ToLower(result.String())

	// Handle common types with better names
	switch {
	case strings.Contains(typeName, "Context"):
		return "context"
	case strings.Contains(typeName, "Request"):
		return "request"
	case strings.Contains(typeName, "Options"):
		return "options"
	default:
		// If the name is generic, add param index
		if flagName == "string" || flagName == "int" || flagName == "bool" {
			return fmt.Sprintf("param-%d", index)
		}
		return flagName
	}
}

// addFlagForType adds an appropriate flag to the command based on the parameter type
func addFlagForType(cmd *cobra.Command, paramName string, paramType reflect.Type, _ int) {
	// Skip context.Context parameters as they're handled internally
	if strings.Contains(paramType.String(), "context.Context") {
		return
	}

	// Handle pointer types
	if paramType.Kind() == reflect.Ptr {
		paramType = paramType.Elem()
	}

	// Add flags based on type kind
	switch paramType.Kind() {
	case reflect.String:
		cmd.Flags().String(paramName, "", fmt.Sprintf("Parameter: %s (string)", paramName))
	case reflect.Int, reflect.Int32, reflect.Int64:
		cmd.Flags().Int(paramName, 0, fmt.Sprintf("Parameter: %s (int)", paramName))
	case reflect.Bool:
		cmd.Flags().Bool(paramName, false, fmt.Sprintf("Parameter: %s (bool)", paramName))
	case reflect.Float32, reflect.Float64:
		cmd.Flags().Float64(paramName, 0.0, fmt.Sprintf("Parameter: %s (float)", paramName))
	case reflect.Struct:
		// For struct types, we could potentially generate nested flags
		// For now, we'll just indicate it needs JSON input
		cmd.Flags().String(paramName, "", fmt.Sprintf("Parameter: %s (JSON object of type %s)", paramName, paramType.Name()))
	case reflect.Slice:
		// Handle slices
		elemType := paramType.Elem()
		if elemType.Kind() == reflect.String {
			cmd.Flags().StringSlice(paramName, []string{}, fmt.Sprintf("Parameter: %s (string array)", paramName))
		} else {
			cmd.Flags().String(paramName, "", fmt.Sprintf("Parameter: %s (JSON array of %s)", paramName, elemType.Name()))
		}
	default:
		// For other complex types, accept JSON string
		cmd.Flags().String(paramName, "", fmt.Sprintf("Parameter: %s (JSON of type %s)", paramName, paramType.String()))
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
