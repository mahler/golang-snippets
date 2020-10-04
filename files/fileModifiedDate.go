package main

/*
 * Simple demonstration of how to read the date modified timestamp of a file.
 */

import (
	"fmt"
	"os"
)

func main() {

	filename := "fileModifiedDate.go"

	// get last modified time
	file, err := os.Stat(filename)

	if err != nil {
		fmt.Println(err)
	}

	modifiedtime := file.ModTime()

	fmt.Println("Last modified time : ", modifiedtime)
}
