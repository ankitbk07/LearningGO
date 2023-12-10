package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
/* Same this as above */
// func main() {

// 	welcome := "Welcome to this World"

// 	fmt.Println(welcome)
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Println("Enter the rating for our Pizza: ")
// 	//comma, ok || er ok try and catch
// 	input, _ := reader.ReadString('\n')
// 	fmt.Println("Thanks for the rating, ", input)
// 	fmt.Printf("type of rating %T, ", input)

// }

// Conversion
func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza between 1 and 5")
	reader := bufio.NewReader(os.Stdin) //os.Stdin was used to take input from keyboard
	input, _ := reader.ReadString('\n')

	fmt.Println("thanks fot rating, ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64) //convert the string into float64
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Added 1 to your rating: ",numRating+1)
		fmt.Printf("Type of rating is %T",numRating)
	}

}