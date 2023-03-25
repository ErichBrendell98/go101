package main

import "fmt"

func main()  {
	// maps = hash table (has no ordering)
	ages := make(map[string]uint8)
	ages["Erich"] = 24
	ages["Tiago"] = 31
	ages["Dani"] = 36
	ages["Maria"] = 22

	fmt.Println(ages)

	fmt.Println(ages["Erich"])
	fmt.Println(ages["Lucas"])

	val, ok := ages["Erich"]
	fmt.Println(val, ok)
	val2, ok := ages["Lucas"]
	fmt.Println(val2, ok)

	// val := &idades["Lucas"]	=> don't do that
}