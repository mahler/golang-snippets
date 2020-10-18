package main

// Sourced from https://stackoverflow.com/questions/60497938/read-exif-metadata-with-go
// Make sure to:
//  go get github.com/rwcarlsen/goexif/exif
//  go get github.com/rwcarlsen/goexif/tiff

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
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

			var dateFileName, err = getDateFromPhotofilename(currentName)
			if err != nil {
				log.Fatal(err)
			}

			fixedDate := fixName(dateFileName)
			fileName := fixedDate
			var extension = filepath.Ext(currentName)
			fullFileName := fileName + extension

			//	fmt.Println(fullFileName)
			os.Rename(currentName, fullFileName)

			log.Println(currentName, "->", fullFileName)

			err := os.Rename(renameDirectory+currentName, renameDirectory+fullFileName)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func getDateFromPhotofilename(fname string) (string, error) {
	// Validate that file exits
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Try to read photoData
	photoData, err := exif.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	photoDate, err := photoData.Get(exif.DateTimeOriginal)
	if err == nil {
		// fmt.Println("hit 1")
	} else if photoDate, err = photoData.Get(exif.DateTime); err != nil {
		// fmt.Println("hit 2")
	} else if photoDate, err = photoData.Get(exif.DateTimeDigitized); err != nil {
		// fmt.Println("hit 3")
	}

	returnDate := photoDate.String()

	if len(returnDate) == 0 {
		theFile, err := os.Stat(fname)
		returnDate = theFile.ModTime().String()
	}

	return returnDate, err
}

func fixName(dirtyName string) string {
	newName := strings.ToLower(dirtyName)
	newName = strings.ReplaceAll(newName, " ", "-")
	newName = strings.ReplaceAll(newName, "(", "-")
	newName = strings.ReplaceAll(newName, ")", "-")
	newName = strings.ReplaceAll(newName, ":", "-")
	newName = strings.ReplaceAll(newName, "\"", "")
	return newName
}
