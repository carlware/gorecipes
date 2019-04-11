package greetings

import "fmt"

var MagicNumber int

func GoGreetings() {

	fmt.Println("A very jolly hello my fellow gophers! I'm printing from the GoGreetings() function")
	printGreetingsUnexported()

}

func init() {

	MagicNumber = 108

}
