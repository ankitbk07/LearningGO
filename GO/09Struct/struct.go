package main

import "fmt"

type User struct{
	Name string
	Email string
	Status bool
	Age int
	//capital letter means they can be used by anyone its global
}
func main() {
	fmt.Println("Structs in GOlang")
	ankit:= User{"ANkit","ankit@gmail.com", true, 16} //Its curly braces not brackets

	fmt.Println(ankit)

	fmt.Printf("The details are: %+v",ankit) // +v helps to print hte struct
	fmt.Printf("The name is %v and email is %v ",ankit.Name,ankit.Email) /// %v is value and we can use struct like a object
	//There is no concept of inheritance in this
}
