package main

import "fmt"
//prints types of a-b
func main() {

	a := 10
	b := "golang string"
	c := 4.17
	d := true

	fmt.Printf("%v is a %T \n", a, a)
	fmt.Printf("%v is a%T \n", b, b)
	fmt.Printf("%v is a%T \n", c, c)
	fmt.Printf("%v is a%T \n", d, d)
}