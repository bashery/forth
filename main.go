package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("wilcome in forth programing languete\n")

	scanner := bufio.NewScanner(os.Stdin)
	var text string

	for scanner.Scan() {
		line := scanner.Text()
		quit := strings.Contains(line, "quit")
		colon := strings.Contains(line, ";")

		text += "\n" + line
		if colon {
			fmt.Println(text, "\n")
			text = ""
		}

		if quit {
			fmt.Println("bye")
			return
		}

		//fmt.Println(text) // Println will add back the final '\n'
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}
