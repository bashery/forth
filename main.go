package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Forth programing language, writen in go")
	//num := os.Args

	t := 20
	for i := 0; i < 10; i++ {
		t += ((t / 100) * 20) + 20
	}

	fmt.Println(t)

}
