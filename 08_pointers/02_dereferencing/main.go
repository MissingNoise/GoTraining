package main

import "fmt"

func main()  {
	a:= 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a

	fmt.Println(b)
	// get the value of a memory address
	fmt.Println(*b)
}