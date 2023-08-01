package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

func replaceColorVariables(inputFile string, outputFile string) error {
	// Read the CSS file
	cssBytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return fmt.Errorf("failed to read CSS file: %v", err)
	}

	// Convert the CSS content to string
	cssContent := string(cssBytes)

	// Define the regular expression pattern to match color variables
	colorVariablePattern := `--([a-zA-Z0-9_-]+):\s*([^;]+);`

	// Compile the regular expression pattern
	colorVariableRegex := regexp.MustCompile(colorVariablePattern)

	// Find all color variable matches in the CSS content
	colorVariableMatches := colorVariableRegex.FindAllStringSubmatch(cssContent, -1)

	// Create a map to store the color variable replacements
	colorReplacements := make(map[string]string)

	// Iterate over the color variable matches
	for _, match := range colorVariableMatches {
		// Extract the color variable name and value
		varName := match[1]
		varValue := match[2]

		// Replace the color variable instance with the hardcoded #hex value
		cssContent = strings.ReplaceAll(cssContent, "var(--"+varName+")", varValue)

		// Store the color variable replacement
		colorReplacements[varName] = varValue
	}

	// Write the updated CSS content to the output file
	err = ioutil.WriteFile(outputFile, []byte(cssContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write output file: %v", err)
	}

	// Print the color variable replacements
	fmt.Println("Color Variable Replacements:")
	for varName, varValue := range colorReplacements {
		fmt.Printf("%s: %s\n", varName, varValue)
	}

	return nil
}

func main() {
	// Provide the input and output file paths
	//inputFile := "input.css"
	//outputFile := "output.css"

	// Get the input and output file names from command line arguments
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <input_file> <output_file>")
		return
	}
	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// Call the function to replace color variables and generate the output file
	err := replaceColorVariables(inputFile, outputFile)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("CSS file processed successfully!")
}
