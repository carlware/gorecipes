package main

import "fmt"

func main() {
	sum, diff := sumAndDiff(100, 8)
	fmt.Printf("100 - 8: The Sum: %d Difference: %+v\n", sum, diff)

	fmt.Println("Sum all result: ", sumAll(0, 5, 76, 2, 3, 4, 8, 1, 9))

	func() {
		fmt.Println("Hi I'm an anonymous function!")
	}()
}

func sumAndDiff(x, y int) (sum int, diff int) {
	return x + y, x - y
}

func sumAll(args ...int) int {
	sum := 0
	for i := 0; i < len(args); i++ {
		sum += args[i]
	}
	return sum
}
