package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start gorutines")

	go func() {
		defer wg.Done()
		fmt.Println("gorutine 1 prints")
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
		fmt.Println("")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("gorutine 2 prints")
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
		fmt.Println("")
	}()

	fmt.Println("wait to finish")
	wg.Wait()
	fmt.Println("\nTerminating Program")
}
