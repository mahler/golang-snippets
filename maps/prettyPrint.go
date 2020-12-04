package main

import (
	"fmt"
	"strings"
)

func main() {
	myMap := make(map[string]string)

	myMap["sample"] = "data"
	myMap["example"] = "more data"
	myMap["something"] = "else"

	printMap(myMap)

}

func printMap(m map[string]string) {
	fmt.Println("PrintMap")
	var maxLenKey int
	for k := range m {
		if len(k) > maxLenKey {
			maxLenKey = len(k)
		}
	}

	for k, v := range m {
		fmt.Println(k + ": " + strings.Repeat(" ", maxLenKey-len(k)) + v)
	}
}
