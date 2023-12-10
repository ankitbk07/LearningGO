package main

import "fmt"
//Pro is really good
func main(){
	fmt.Println("Welcome to functions")
	result,myMessage := proAdder(3,5,10,20)

	fmt.Println("The sum is", result)
	fmt.Println(myMessage)

}

func greeter(){
	fmt.Println("Hello from Nepal")
}

func proAdder(values ...int)(int,string){ 
	totals := 0

	for _,i := range(values){
		totals += i
	}
	return totals, "Hi pro Result function"
} //proadder takes as many values as it wants and really thats cool

func adder(one int, two int)int{
	return one+two
} // Function signature must be given