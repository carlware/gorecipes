package main

import "fmt"

const (
	_        = iota
	redLight // Test
	greenLight
	yellowLight
)

func main() {
	const (
		LosAngeles = 1984 + (iota * 4)
		Seoul
		Barcelona
		Atlanta
		Sydney
		Athens
		Beijing
		London
		Rio
		Tokyo
	)

	fmt.Println("These cities hosted or will host the Summer Olympics in the provided year...")
	fmt.Printf("%-18s %-18s \n", "City", "Year")
	fmt.Printf("%-18s %-18v \n", "Los Angeles", LosAngeles)
	fmt.Printf("%-18s %-18v \n", "Atlanta", Atlanta)
	fmt.Printf("%-18s %-18v \n", "Tokyo", Tokyo)
	fmt.Println("Red Light State Code: ", redLight)
	fmt.Println("Green Light State Code: ", greenLight)
	fmt.Println("Yellow Light State Code: ", yellowLight)
}