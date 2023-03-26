package main

import (
	"fmt"
	"time"
)

func numbers(n chan<- int)  {
	for i := 0; i < 10; i++ {
		n <- i
		fmt.Printf("Written on the channel: %d\n", i)
		time.Sleep(time.Millisecond * 90)
	}
	fmt.Println("End of writing")
	close(n)
}

func main()  {
	cn := make(chan int, 3)
	go numbers(cn)

	for v := range cn {
		fmt.Printf("Read from the channel: %d\n", v)
		time.Sleep(time.Millisecond * 180)
	}

	fmt.Println("End of execution")
}