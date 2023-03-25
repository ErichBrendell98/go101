package main

import "fmt"

type People struct {
	// public attributes
	Name string
	Surname string
	Age uint8
	Status bool
	// private attributes
	cpf string
}

type Category struct {
	Name string
	Father *Category
}

func main()  {
	// 1º way
	var p People
	p.Name = "Erich Brendell"
	p.Surname = "Araújo Medeiros"
	p.Age = 24
	p.Status = true
	fmt.Printf("p: %v\n", p)
	fmt.Println(p)
	
	println()
	
	// 2º way
	p2 := People{
		Name: "Brendell Erich",
		Surname: "Medeiros Araújo",
		Age: 42,
		Status: false,
	}
	fmt.Printf("p2: %v\n", p2)
	fmt.Println(p2)
	
	println()
	
	// 3º way - it isn't recommended to do this because the structure may change
	p3 := People{"Tiago", "Temporin", 31, true, "000.000.000 - 00"}
	fmt.Printf("p3: %v\n", p3)
	fmt.Println(p3)
}