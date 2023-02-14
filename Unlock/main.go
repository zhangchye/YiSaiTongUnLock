package main

import (
	"flag"
	"fmt"
	"os"
)

var sourceFilePath = flag.String("sourcePath", "", "source file path")
var destFilePath = flag.String("destPath", "", "destination file path")

func main() {
	flag.Parse()
	sourceFilePath := *sourceFilePath
	destFilePath := *destFilePath
	err := os.Rename(sourceFilePath, destFilePath)
	if err != nil {
		fmt.Printf("Error:%s---->%s", sourceFilePath, destFilePath)
	}
}
