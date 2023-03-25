package main

import "fmt"

type Person struct {
	Name string
	Age uint8
	Status bool
}

func (p Person) String() string {
	return fmt.Sprintf("Hello, my name is %s and I'm %d years old", p.Name, p.Age)
}

type PhysicalPerson struct {
	Person
	Surname string
	cpf string
}

type LegalPerson struct {
	CorporateName string
	cnpj string
}

func main()  {
	p := PhysicalPerson{
		Person{Name: "Erich", Age: 24, Status: true},
		"Brendell",
		"000.000.000 - 00",
	}

	fmt.Println(p)
}