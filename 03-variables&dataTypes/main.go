package main

import "fmt"

var name string

var (
	n1 int
	n2 int32
)

func main()  {
	message := "Class 03 - Go 101"
	fmt.Println(message)

	var (
		b1 float32
		b2 float64
	)

	name = "Erich Brendell"
	fmt.Println("Hello", name, "!")

	b1 = 6.5
	b2 = 8.9
	println("b1:", b1, "b2:", b2)

	var total int	// total = 0
	fmt.Println("total:", total)
	total++	// total = total + 1 => total =	1
	fmt.Println("total:", total)

	var b, f, s = true, 2.3, "Hello"
	fmt.Println("b:", b, "- f:", f, "- s:", s)

	var x, y = 10, 20
	fmt.Println("x =", x,"y =", y)
	x, y = y, x
	fmt.Println("x =", x,"y =", y)
}