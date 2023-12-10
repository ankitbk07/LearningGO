package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class of pointers")
	var ptr *int
	i := 10

	ptr = &i

	fmt.Println("The value of the pointers is ",ptr)

}