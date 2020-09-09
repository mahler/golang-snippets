package main

// Sourced from https://stackoverflow.com/questions/60497938/read-exif-metadata-with-go
// Make sure to:
//  go get github.com/rwcarlsen/goexif/exif
//  go get github.com/rwcarlsen/goexif/tiff

import (
	"fmt"
	"log"
	"os"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/tiff"
)

// Printer is an empty struct to allow any struct to use utility function
type Printer struct{}

// Walk is a utility function to print all exif fields found.
func (p Printer) Walk(name exif.FieldName, tag *tiff.Tag) error {
	fmt.Printf("%40s: %s\n", name, tag)
	return nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("please give filename as argument")
	}
	fname := os.Args[1]

	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}

	x, err := exif.Decode(f)
	if err != nil {
		log.Fatal(err)
	}

	var p Printer
	x.Walk(p)
}
