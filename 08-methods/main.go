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

func (p People) String() string {
	return fmt.Sprintf("Hello, my name is %s and I'm %d years old", p.Name, p.Age)
}

type Category struct {
	Name string
	Parent *Category
}

func (c Category) HasParent() bool {
	return c.Parent != nil
}

func (c *Category) SetParent(parent *Category) {
	c.Parent = parent
}

func main()  {
	p := People{"Erich", "Brendell", 24, true, "000.000.000 - 00"}
	p.cpf = "1"
	fmt.Println(p)

	cat := Category{Name: "Category 1"}
	cat.SetParent(&Category{Name: "Father"})
	if !cat.HasParent() {
		fmt.Println("Don't have parent")
	}
}