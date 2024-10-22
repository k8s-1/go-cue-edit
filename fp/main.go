package main

import (
	"fmt"
  "path/filepath"
)

func main() {
  p := filepath.Join("dir1", "dir2", "filename")
  fmt.Println("p:", p)
}
