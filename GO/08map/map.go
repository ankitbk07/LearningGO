package main

import "fmt"

func main() {
	fmt.Println("Welcome to world of map")
	// map syntax: map[key_type]value_type
	language := make(map[string]string)
	language["JS"] = "Javascript"
	language["RB"] = "RUBY"
	
	fmt.Println("List ", language)

	delete(language,"RB")
	fmt.Println("List ", language)

	//loops
	for key,value := range language{
		fmt.Printf("For key%v, value is %v\n",key, value)
	}
}
