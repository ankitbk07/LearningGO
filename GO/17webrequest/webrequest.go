package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://google.com"

func main() {
	fmt.Println("Connection")
	response,err := http.Get(url) //Sees if the conn
	if err != nil{
		panic(err)
	}

	fmt.Printf("Response is of: %T\n",response)
	defer response.Body.Close() //caller's Responsibility

	dataByte ,err := io.ReadAll(response.Body)
	if err != nil{
		panic(err)
	}
	status:= response.Status
	println(status)
	content := string(dataByte)
	fmt.Println(content)

}