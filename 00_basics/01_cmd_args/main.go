package main

import (
	"flag"
	"fmt"
)

func main() {

	var name string
	flag.StringVar(&name, "name", "World", "User name")
	flag.Parse()
	fmt.Println("Hello " + name + "!")
}

