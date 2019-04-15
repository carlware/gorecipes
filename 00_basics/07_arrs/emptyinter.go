package main

import "fmt"

func main() {
	anyType := make([]interface{}, 10)

	anyType = append(anyType, 1)
	anyType = append(anyType, "two")
	anyType = append(anyType, false)
	anyType = append(anyType, 15.9)

	fmt.Println(anyType)

	noInitialSize := []interface{}{}
	noInitialSize = append(noInitialSize, 1)
	noInitialSize = append(noInitialSize, "two")
	noInitialSize = append(noInitialSize, false)
	noInitialSize = append(noInitialSize, 90.9)

	fmt.Println(noInitialSize)

}
