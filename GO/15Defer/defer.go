package main

import "fmt"

func main() {
	//A defer statement defers the execution of a function until the surrounding function returns.
	//JUST SEE LIFO
	defer fmt.Println("world")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("Hello")

	myDefer() 

}
//myDefer is a example of stack
func myDefer(){
	for i:=0;i<5;i++{
		defer fmt.Print(i)
	}
}