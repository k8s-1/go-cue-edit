package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
  "strings"
)

func main() {
	// Read JSON data from the file
	data, err := os.ReadFile("version-test.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert the data to a string
	content := string(data)

	// Find the index of the first occurrence of '{'
	index := strings.Index(content, "{")

	// Check if '{' was found
	if index == -1 {
		log.Fatal("No opening brace found in the content.")
		return
	}

	// Store the characters before '{' in a variable
	beforeBrace := content[:index]
	
	// Extract the JSON string starting from '{'
	jsonString := content[index:]

	// Print the results
	fmt.Println("Characters before '{':", beforeBrace)
	fmt.Println("JSON String:", jsonString)
}



func updateVersion() {
	jsonData, err := os.ReadFile("version.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	var appVersions map[string]string

	err = json.Unmarshal(jsonData, &appVersions)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	fmt.Printf("%+v\n", appVersions)

	// Update the value of app1
	appVersions["app1"] = "999"

	// Print the updated map
	fmt.Printf("Updated: %+v\n", appVersions)

	// Optionally, you can write the updated map back to the file (if desired)
	// To do this, you would marshal the map back to JSON and save it
	updatedJsonData, err := json.MarshalIndent(appVersions, "", "  ")
	if err != nil {
		log.Fatal("Error during MarshalIndent(): ", err)
	}

	err = os.WriteFile("version-test.json", append(updatedJsonData, '\n'), 0644)
	if err != nil {
		log.Fatal("Error writing updated file:", err)
	}
}
