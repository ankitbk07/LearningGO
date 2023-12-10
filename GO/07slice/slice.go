package main

import (
	"fmt"
)

func main() {
	var fruitList = []string{"apple","mango","peach"}
	fmt.Println("Types of fruitlist is \n",len(fruitList))

	// fruitList = append(fruitList[1:]))
	
	fmt.Println(fruitList)
	fruitList = append(fruitList[1:2])//range are non inclusice

	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 234
	highScores[1] = 224
	highScores[2] = 214
	highScores[3] = 204

	highScores =append(highScores, 555,666,321) // This helps to make this make memory efficient to be really cool
	// fmt.Println(highScores)
	
	// sort.Ints(highScores)

	// fmt.Println(highScores)

	//how to remove a value from slices based index

	var courses = []string{"life","history","python","ruby"}

	fmt.Println(courses)
	var index int  = 2

	courses = append(courses[:index],courses[index+1:]...) //These removes the data from the array we want to
	fmt.Println(courses)

}