package main

import "fmt"

type Document interface {
	Doc() string
}

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

func (pp PhysicalPerson) Doc() string {
	return fmt.Sprintf("CPF: %s", pp.cpf)
}

type LegalPerson struct {
	Person
	CorporateName string
	cnpj string
}

func (lp LegalPerson) Doc() string {
	return fmt.Sprintf("CNPJ: %s", lp.cnpj)
}

func show(d Document) {
	fmt.Println(d)
	fmt.Println(d.Doc())
}
func main()  {
	pp := PhysicalPerson{
		Person{Name: "Erich", Age: 24, Status: true},
		"Brendell",
		"000.000.000 - 00",
	}

	show(pp)

	fmt.Println()

	lp := LegalPerson{
		Person{Name: "Aprenda Golang", Age: 0, Status: true},
		"Temporin LTDA",
		"00.000.000/000 - 00",
	}

	show(lp)
}