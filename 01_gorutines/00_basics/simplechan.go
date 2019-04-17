package main

import "fmt"

func main() {
	message := make(chan string, 3)
	message <- "one"
	message <- "two"
	message <- "three"

	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
}
