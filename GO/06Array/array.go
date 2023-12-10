package main

import "fmt"

func main() {
	var fruits [4]string
	var veg = [4] string{"potato","Biter gourd","Tomato","Collard Green"} //another way to do it.
	fruits[0] = "Ankit"

	fmt.Println("THe fruits are ", len(fruits)) //Array not that used slices are used.
	fmt.Println(veg)
}