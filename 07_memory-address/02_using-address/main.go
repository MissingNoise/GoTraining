package main

import "fmt"

const metersToYards float64 = 1.09361

func main() {
	var meters float64
	fmt.Print("Enter meters swam: ")
	// ask for meters
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println("you swam ", meters, " meters or ", yards, " yards.")
}