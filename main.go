package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Define a map to capture the app versions
var appVersions map[string]string

func main() {
	// Open the JSON file
	jsonData, err := os.ReadFile("version.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}

	err = json.Unmarshal(jsonData, &appVersions)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v", jsonData)
}
