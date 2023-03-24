package main

import (
	"fmt"
	"strconv"
)

func Hello(nome string)  {
	fmt.Println("Hello", nome, "!")
}

func sum(a, b int) int {
	return a + b
}

func convertAndSum(a int, b string) (total int, err error) {
	i, err := strconv.Atoi(b)
	if err != nil {
		return
	}

	total = a + i

	return
}

func main()  {
	Hello("Erich Brendell")
	fmt.Println("total:", sum(10, 5))
	total, err := convertAndSum(10, "23")
	fmt.Println("total:", total, err)
}