package main

import "fmt"

func main()  {
	for i := 0; i <=10; i++ {
		for p := 0; p < 10; p++ {
			fmt.Println(i, " - ", p)
		}
	}
}