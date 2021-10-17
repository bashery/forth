package main

import (
	"fmt"
)

func main() {
	fmt.Println("wilcome in forth programing languete")

	var input string

	for {

		fmt.Print("> ")
		_, err := fmt.Scanf("%s", &input)

		if err != nil {
			fmt.Println("my error is : ", err)
		}

		fmt.Println(input)
		if input == "quit" {
			fmt.Println("bye")
			return
		}

	}

}
