package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
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
