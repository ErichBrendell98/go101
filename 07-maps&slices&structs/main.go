package main

import "fmt"

func main()  {
	// slices
	names := []string{"Thiago", "Dani", "Marcos", "Maria"}		// dynamic length
	fmt.Println(names, len(names), cap(names))
	
	names = append(names, "Rafael")								// appending in slices
	fmt.Println(names, len(names), cap(names))
	names = append(names, "Rubia")
	fmt.Println(names, len(names), cap(names))

	names2 := make([]string, 10, 20)		// string array[]: len = 10, cap = 20
	fmt.Println(names2, len(names2), cap(names2))
}