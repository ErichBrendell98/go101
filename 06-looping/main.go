package main

import "fmt"

func main()  {
	fmt.Println("##### [ Class 06 - Loops ] #####")

	names := []string{"Tiago", "Daniel", "Maria", "Marta"}
	
	// 1º way to make a for loop
	fmt.Println(">>>> First way to make a for loop:")
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}
	
	fmt.Println()
	
	//2º way to make a for loop
	fmt.Println(">>>> Second way to make a for loop:")
	for _, name := range names {
		fmt.Println(name)
	}

	fmt.Println()
	
	//3º way to make a for loop
	fmt.Println(">>>> Third way to make a for loop:")
	var i int
	for i < len(names) {
		fmt.Println(names[i])
		i++
	}
}