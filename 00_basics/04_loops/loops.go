package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		if i == 0 {
			continue
		}
		fmt.Println("i: ", i)
	}
	fmt.Println("")

	j := -20
	for j != 0 {
		fmt.Println("loop cond: ", j)
		j++
	}
	fmt.Println("")

	loopTimer := time.NewTimer(time.Second * 9)

	for {
		fmt.Println("Inside loop")
		<-loopTimer.C
		break
	}
}
