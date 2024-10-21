package main

import (
    "fmt"
    // "os"
    "cuelang.org/go/cue"
    "cuelang.org/go/cue/cuecontext"
    "cuelang.org/go/cue/load"
    // "cuelang.org/go/cue/build"
)

func main() {
    // Specify the CUE entrypoint (directory or file) to load, including imports
    cuePath := "./" // Assuming your .cue file and imports are in the current directory
    cueFile := "example.cue"

    // Create a CUE context
    ctx := cuecontext.New()

    // Load the CUE instance using the cue/load package
    config := &load.Config{
        Dir: cuePath, // Directory where the CUE file and imports are located
    }

    // Load the CUE package including the file and its imports
    instances := load.Instances([]string{cueFile}, config)
    if len(instances) == 0 || instances[0].Err != nil {
        fmt.Println("Error loading CUE instance:", instances[0].Err)
        return
    }

    // Build the CUE instance to get the value
    instance := ctx.BuildInstance(instances[0])
    if instance.Err() != nil {
        fmt.Println("Error building CUE instance:", instance.Err())
        return
    }

    // Modify a specific value in the CUE file (for example, "myfield")
    path := cue.ParsePath("myfield")
    updatedInstance := instance.FillPath(path, "newValue")

    fmt.Println(updatedInstance)
} 
