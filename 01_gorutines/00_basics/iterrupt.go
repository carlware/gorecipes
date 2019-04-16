package main

import (
	"fmt"
	"os"
	"os/signal"
)

func main() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	fmt.Println("waiting interrupt ...")
	for {
		select {
		case <-interrupt:
			fmt.Println("exiting ...")
			os.Exit(0)
		}
	}
}
