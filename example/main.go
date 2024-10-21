package main

import (
	// "fmt"
  _ "embed"
	// "os"
	//
	// "cuelang.org/go/cue"
	// "cuelang.org/go/cue/ast"
	// "cuelang.org/go/cue/cuecontext"
	// "cuelang.org/go/cue/format"
	// "cuelang.org/go/cue/load"
)

//go:embed example.cue
var s string

func main() {
  print(s)
}
