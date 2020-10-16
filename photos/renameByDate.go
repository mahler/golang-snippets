package main

// Sourced from https://stackoverflow.com/questions/60497938/read-exif-metadata-with-go
// Make sure to:
//  go get github.com/rwcarlsen/goexif/exif
//  go get github.com/rwcarlsen/goexif/tiff

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/rwcarlsen/goexif/exif"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("please give filename as argument")
	}

	fname := os.Args[1]

	var dateFileName, err = getDateFromPhotofilename(fname)
	if err != nil {
		log.Fatal(err)
	}

	fixedDate := fixName(dateFileName)
	fileName := fixedDate
	var extension = filepath.Ext(fname)
	fullFileName := fileName + extension

	//	fmt.Println(fullFileName)
	os.Rename(fname, fullFileName)
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
