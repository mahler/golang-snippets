package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		output := ""
		if i%3 == 0 {
			output += "Fizz"
		}
		if i%5 == 0 {
			output += "Buzz"
		}

		if output == "" {
			// Convert int to string.
			output = fmt.Sprintf("%v", i)
		}

		fmt.Println(output)
	}
}
