// Package main provides a CLI tool to extract SDK descriptions
package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/equinix/cli/internal/parser"
)

func main() {
	sdkPath := flag.String("sdk-path", "", "Path to the SDK source directory (required)")
	outputFile := flag.String("output", "descriptions.json", "Output JSON file path")
	flag.Parse()

	if *sdkPath == "" {
		fmt.Fprintln(os.Stderr, "Error: --sdk-path is required")
		flag.Usage()
		os.Exit(1)
	}

	fmt.Printf("Extracting descriptions from SDK at: %s\n", *sdkPath)
	descriptions, err := parser.ExtractDescriptions(*sdkPath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error extracting descriptions: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Found %d services\n", len(descriptions.Services))
	for name, service := range descriptions.Services {
		fmt.Printf("  - %s: %d methods\n", name, len(service.Methods))
	}

	fmt.Printf("Saving descriptions to: %s\n", *outputFile)
	if err := descriptions.SaveToFile(*outputFile); err != nil {
		fmt.Fprintf(os.Stderr, "Error saving descriptions: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("Successfully extracted and saved SDK descriptions")
}
