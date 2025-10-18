// Package parser provides utilities for extracting SDK documentation
// from Go source files using go/ast, go/parser, and go/token.
package parser

import (
	"encoding/json"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// MethodDescription holds documentation for a method
type MethodDescription struct {
	ShortDescription string            `json:"short_description"`
	LongDescription  string            `json:"long_description"`
	Parameters       map[string]string `json:"parameters,omitempty"`
}

// ServiceDescriptions holds all descriptions for a service
type ServiceDescriptions struct {
	ServiceName        string                       `json:"service_name"`
	ServiceDescription string                       `json:"service_description"`
	Methods            map[string]MethodDescription `json:"methods"`
}

// SDKDescriptions holds all service descriptions
type SDKDescriptions struct {
	Services map[string]ServiceDescriptions `json:"services"`
}

// ExtractDescriptions extracts method and parameter descriptions from SDK source files
func ExtractDescriptions(sdkPath string) (*SDKDescriptions, error) {
	descriptions := &SDKDescriptions{
		Services: make(map[string]ServiceDescriptions),
	}

	// Walk through the SDK directory
	err := filepath.Walk(sdkPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Only process .go files (not test files)
		if !info.IsDir() && strings.HasSuffix(path, ".go") && !strings.HasSuffix(path, "_test.go") {
			if err := parseFile(path, descriptions); err != nil {
				// Log error but continue processing other files
				fmt.Fprintf(os.Stderr, "Warning: failed to parse %s: %v\n", path, err)
			}
		}

		return nil
	})

	return descriptions, err
}

// parseFile parses a single Go source file and extracts documentation
func parseFile(filePath string, descriptions *SDKDescriptions) error {
	fset := token.NewFileSet()
	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return err
	}

	// Extract service name from file (e.g., api_connections.go -> connections)
	baseName := filepath.Base(filePath)
	var serviceName string
	if strings.HasPrefix(baseName, "api_") && strings.HasSuffix(baseName, ".go") {
		serviceName = strings.TrimSuffix(strings.TrimPrefix(baseName, "api_"), ".go")
		serviceName = strings.ReplaceAll(serviceName, "_", "-")
	}

	// Walk through all declarations in the file
	for _, decl := range node.Decls {
		funcDecl, ok := decl.(*ast.FuncDecl)
		if !ok || funcDecl.Doc == nil {
			continue
		}

		// Check if this is a method on a service
		if funcDecl.Recv == nil || len(funcDecl.Recv.List) == 0 {
			continue
		}

		// Extract receiver type name
		var receiverType string
		switch recv := funcDecl.Recv.List[0].Type.(type) {
		case *ast.StarExpr:
			if ident, ok := recv.X.(*ast.Ident); ok {
				receiverType = ident.Name
			}
		case *ast.Ident:
			receiverType = recv.Name
		}

		// Check if this is an API service (ends with ApiService)
		if !strings.HasSuffix(receiverType, "ApiService") {
			continue
		}

		// Extract service name from receiver type if not already set
		if serviceName == "" && strings.HasSuffix(receiverType, "ApiService") {
			serviceName = strings.TrimSuffix(receiverType, "ApiService")
			serviceName = camelToKebab(serviceName)
		}

		// Ensure service exists in map
		if _, exists := descriptions.Services[serviceName]; !exists {
			descriptions.Services[serviceName] = ServiceDescriptions{
				ServiceName: serviceName,
				Methods:     make(map[string]MethodDescription),
			}
		}

		service := descriptions.Services[serviceName]

		// Extract method documentation
		methodName := funcDecl.Name.Name
		doc := funcDecl.Doc.Text()

		// Parse the documentation
		lines := strings.Split(strings.TrimSpace(doc), "\n")
		var shortDesc, longDesc string
		var params map[string]string

		if len(lines) > 0 {
			// First line is typically "MethodName Short Description"
			firstLine := strings.TrimSpace(lines[0])
			// Remove method name from start if present
			if strings.HasPrefix(firstLine, methodName) {
				shortDesc = strings.TrimSpace(strings.TrimPrefix(firstLine, methodName))
			} else {
				shortDesc = firstLine
			}

			// Rest is long description
			if len(lines) > 1 {
				longLines := []string{}
				params = make(map[string]string)

				for i := 1; i < len(lines); i++ {
					line := strings.TrimSpace(lines[i])
					// Check for @param annotations
					if strings.HasPrefix(line, "@param") {
						parts := strings.SplitN(strings.TrimPrefix(line, "@param"), " ", 3)
						if len(parts) >= 3 {
							paramName := strings.TrimSpace(parts[0])
							paramDesc := strings.TrimSpace(parts[2])
							params[paramName] = paramDesc
						}
					} else if !strings.HasPrefix(line, "@") && line != "" {
						longLines = append(longLines, line)
					}
				}
				longDesc = strings.Join(longLines, " ")
			}
		}

		// Store method description
		service.Methods[methodName] = MethodDescription{
			ShortDescription: shortDesc,
			LongDescription:  longDesc,
			Parameters:       params,
		}

		descriptions.Services[serviceName] = service
	}

	return nil
}

// camelToKebab converts CamelCase to kebab-case
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

// SaveToFile saves descriptions to a JSON file
func (d *SDKDescriptions) SaveToFile(filePath string) error {
	data, err := json.MarshalIndent(d, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal descriptions: %w", err)
	}

	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// LoadFromFile loads descriptions from a JSON file
func LoadFromFile(filePath string) (*SDKDescriptions, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var descriptions SDKDescriptions
	if err := json.Unmarshal(data, &descriptions); err != nil {
		return nil, fmt.Errorf("failed to unmarshal descriptions: %w", err)
	}

	return &descriptions, nil
}
