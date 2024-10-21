package main

import (
  "encoding/json"
  "os"
  "log"
)

// The data struct for the decoded data
// Notice that all fields must be exportable!
type Data struct {
    Origin string
    User   string
    Active bool
}

func main() {
}
