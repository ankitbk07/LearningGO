package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")
	days := []string{"sunday","Tuesday","Wednesday","Friday","Saturday"}

	fmt.Println(days)
	//First method
	// for d:=0;d<len(days);d++{
	// 	fmt.Println("The day is ",days[d])
	// }
	for i := range days{
		fmt.Println(days[i])
	}

	// for index, value := range days{
	
	// }
	
}