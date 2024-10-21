package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Data is a map where each key is a string and its value is another map
type Data map[string]map[string]string

func main() {
	// Example JSON data as a string (you can also read from a file)
	jsonData := `{
		"group1": {
			"app1": "123",
			"app2": "456"
		},
		"group2": {
			"app3": "789",
			"app4": "101"
		}
	}`

	// Declare a variable of type Data
	var data Data

	// Unmarshal the JSON data into the Data variable
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		log.Fatal("Error unmarshalling JSON:", err)
	}

	// Print the resulting map
	for group, apps := range data {
		fmt.Printf("%s:\n", group)
		for app, version := range apps {
			fmt.Printf("  %s: %s\n", app, version)
		}
	}

	// Marshal it back to JSON to see the structure
	marshaledData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal("Error marshaling to JSON:", err)
	}

	// Print the marshaled JSON
	fmt.Println(string(marshaledData))
}
