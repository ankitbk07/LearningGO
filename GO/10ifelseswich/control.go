package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// fmt.Println("If else ")

	// loginCount := 10
	// var result string
	// if loginCount < 10{
	// 	result = "Regular user"
	// }else if loginCount > 10{
	// 	result = "Watch Out"
	// }else{
	// 	result = "Exacatky 10 login count"
	// }

	// fmt.Println(result)

	// if 9%2 ==0{
	// 	fmt.Println("Number is even")
	// }else{
	// 	fmt.Println("Number is odd")
	// }

	fmt.Println("Switch and case in golang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6)+1

	fmt.Println("Value of dice is ", diceNumber)

	switch diceNumber{

	case 1:
		fmt.Println("Dice value is 1 and you can open")
	default:
		fmt.Println("you roll again")
	
	}

}