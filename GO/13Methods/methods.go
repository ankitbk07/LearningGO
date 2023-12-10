package main

import "fmt"
//For more contains in https://go.dev/tour/methods/1
type User struct {
	Name   string
	Email  string
	Age    int
	status bool // Capital Letter means its global
}

func main() {
	ankit := User{"ANkit", "ankit@gmail.com", 16, true}
	fmt.Printf("The details are %+v\n",ankit)
	ankit.GetStatus()
	ankit.SetEmail()
	fmt.Printf("The details are %+v\n", ankit)
}
func (u User)GetStatus(){
	fmt.Println("Is user active: ",u.status)
}

func (u *User)SetEmail(){
	u.Email = "ankit1@gmail.com"
	fmt.Println("The email is: ",u.Email)
}
//Methods with pointer receivers can modify the value with ease. When we pass value only, it does not have any affect on the original struct. Value receiver gets of copy of the original User.
//More info https://go.dev/tour/methods/4