package main

import "fmt"

func main() {
	message := make(chan string, 3)
	message <- "one"
	message <- "two"
	message <- "three"

	close(message)

	for m := range message {
		fmt.Println(m)
	}
}
