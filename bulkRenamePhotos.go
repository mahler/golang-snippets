package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {

	renameDirectory := ""
	if len(os.Args) > 1 {
		renameDirectory = os.Args[1]
	} else {
		fmt.Println("Missing directory as parameter")
		os.Exit(10)
	}

	files, err := ioutil.ReadDir(renameDirectory)
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		if !f.IsDir() && (strings.HasSuffix(f.Name(), ".jpg") || strings.HasSuffix(f.Name(), ".png") || strings.HasSuffix(f.Name(), ".webp")) {
			currentName := f.Name()
			newName := fixName(currentName)

			log.Println(currentName, "->", newName)
			err := os.Rename(renameDirectory+currentName, renameDirectory+newName)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func fixName(dirtyName string) string {
	newName := strings.ToLower(dirtyName)
	newName = strings.ReplaceAll(newName, " ", "-")
	newName = strings.ReplaceAll(newName, "(", "-")
	newName = strings.ReplaceAll(newName, ")", "-")
	return newName
}
