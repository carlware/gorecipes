package main

import (
	"fmt"
	"sync"
)

var waitWrite chan bool
var message string
var mutex = &sync.Mutex{}

func greet() {
	mutex.Lock()
	message = "Hello buddy"
	mutex.Unlock()
	waitWrite <- true
}

func main() {

	waitWrite = make(chan bool, 1)
	go greet()
	mutex.Lock()
	message = "Hello friend"
	fmt.Println(message)
	mutex.Unlock()
	<-waitWrite
}
