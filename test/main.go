package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Config struct {
	Vals  []Val `json:"vals"`
	Level int   `json:"level"`
}

type Val struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func main() {
	var cfg Config
	dec := json.NewDecoder(os.Stdin)
	if err := dec.Decode(&cfg); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Config level: %d.\n", cfg.Level)
	fmt.Printf("Val count: %d.\n", len(cfg.Vals))
	if len(cfg.Vals) > 0 {
		fmt.Printf("First val name: %q.\n", cfg.Vals[0].Name)
	}
}
