// In this example we use the greetingspackage
package main

import (
	"fmt"

	"carlware/greet/greetings"
)

func main() {

	greetings.PrintGreetings()
	greetings.GoGreetings()

	fmt.Println("The value of the Magic Number is:", greetings.MagicNumber)

}
