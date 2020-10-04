package main

/*
 * Simple demonstration of how to read all files in a directory.
 */

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {

	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	target := []os.FileInfo{}

	for _, f := range files {
		if !f.IsDir() && strings.HasSuffix(f.Name(), ".go") {
			target = append(target, f)
			log.Println(f.Name())
		}
	}
}
