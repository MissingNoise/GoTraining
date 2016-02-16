package main

import "fmt"

func zero(z int)  {
	fmt.Printf("%p\n", &z)
	fmt.Println(&z)
	z = 0
}

func main()  {
	x := 5
	fmt.Printf("%p\n", &x)
	fmt.Println(&x)
	// passes the value of 5. does not give x the memory address
	zero(x)
	fmt.Println(x)
}