package main

import "fmt"
//If the first letter of the variable or constant is capital then it can be accesed by any other variables .
const Nepal string = "Country"

func main() {
	// var username string = "Ankit"
	var isLoggedIn bool = false 
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n",isLoggedIn)
	var smallVal int = 256
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n",smallVal)

	var smallFLoat float64 = 255.356556
	fmt.Println(smallFLoat)
	fmt.Printf("Variable is of type: %T \n",smallFLoat)

	var anotherVariable string
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n",anotherVariable)

	// implictit type
	var website = "Ankit"
	fmt.Println(website)
	// no var style
	numberOfUser := 300 //Not allowed to become a global variable only inside a local variable only
	fmt.Println(numberOfUser)

	fmt.Println(Nepal)
}

