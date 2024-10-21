package main

import (
	"encoding/json"
	"fmt"
	"os"
  "log"
)

func main() {
	// Open the JSON file
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
}
