package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("wilcome in forth programing languege\nversion: 0.0.1")

	scanner := bufio.NewScanner(os.Stdin)
	scanr := blocCode(scanner)
	if err := scanr.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

}

func blocCode(scanner *bufio.Scanner) *bufio.Scanner {
	var text string
	for scanner.Scan() {
		line := scanner.Text()
		quit := strings.Contains(line, "quit")
		//colon := strings.HasSuffix(line, ";")

		text += "\n" + line

		if strings.HasSuffix(line, ";") {
			fmt.Println(text)
			text = ""
		}

		if quit {
			fmt.Println("bye")
			return nil
		}

		fmt.Println(text) // Println will add back the final '\n'
	}

	return scanner
}

/*
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
*/
