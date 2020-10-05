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

	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

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

	fixedDate := fixName(photoDate.String())
	fileName := fixedDate
	var extension = filepath.Ext(fname)
	fullFileName := fileName + extension

	//	fmt.Println(fullFileName)
	os.Rename(fname, fullFileName)
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
