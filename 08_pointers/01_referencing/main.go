package main

import "fmt"

func main()  {
	a:= 43

	fmt.Println(a)
	fmt.Println(&a)

	//assign b int pointer to memory address a
	var b *int = &a

	fmt.Println(b)
}