package main

import "fmt"

var waitWrite chan bool
var message string

func greet() {
	message = "Hello buddy"
	waitWrite <- true
}

func main() {

	waitWrite = make(chan bool, 1)
	go greet()
	message = "Hello friend"
	fmt.Println(message)
	<-waitWrite
}
