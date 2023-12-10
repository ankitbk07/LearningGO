package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Golang")
	content := "Hello from the file"

	file,err :=os.Create("./logfile.txt") //creates the file

	checkNil(err)//This is try and catch

	length, err := io.WriteString(file,content) //gives a length back

	checkNil(err)

	fmt.Println("The length is",length)
	defer file.Close()

	readFile("./logfile.txt")
}

func readFile(filename string){
	content, err:=os.ReadFile(filename)
	checkNil(err)
	fmt.Println("Text data inside the file is \n", string(content))
}

func checkNil(err error){
	if err != nil{
		panic(err)//Stops the normal execution of go compiler
	}
}