package main

import (
	"fmt"
	"log"
	"os"
)

var (
	head, tail int
)

func tossCoin(side string)  {
	switch side {
	case "head":
		head++

	case "tail":
		tail++

	default:
		fmt.Println("Fell on his feet")
	}
}

func main()  {
	a, b := 10, 13

	// if statement
	if a > b {
		fmt.Println("a > b")
	} else if a < b {
		fmt.Println("a < b")
	} else {
		fmt.Println("a = b")
	}

	// switch statement
	switch {
	case a > b:
		fmt.Println("a > b")
		
	case a < b:
		fmt.Println("a < b")

	default:
		fmt.Println("a =")
	}

	// opening the file
	file, err := os.Open("hello.txt")
	if err != nil {
		log.Panic(err)
	}
	
	// variables within the scope of the IF
	data := make([]byte, 100)
	if _, err := file.Read(data); err != nil {
		log.Panic(err)
	}

	fmt.Println(string(data))	// convert data[] => data (string)
}