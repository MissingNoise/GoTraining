package main

import "fmt"

func max(x int) int  {
	return 42 + x
}

func main()  {
	max := max(7) //bad coding
	fmt.Println(max)
}

// Don't do shadow variables