package main

import (
	"fmt"
)

func main() {
	fmt.Println("wilcome in forth programing languete\n")

	var input string

	for {

		fmt.Print("> ")
		fmt.Scanf("%s", &input)

		if input == "duble" {
			input = input + input
		}

		fmt.Println(input)
		if input == "quit" {
			fmt.Println("bye")
			return
		}

	}
}
