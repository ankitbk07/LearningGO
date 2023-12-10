package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(("Welcome to time study of golang"))
	presentTime := time.Now()
	fmt.Println(presentTime) //this gives the present time in a very big detaiil
	fmt.Println(presentTime.Format("2006-01-02 Monday"))

	createdDate := time.Date(2020,time.August,10,23,23,0,0,time.UTC)

	fmt.Println(createdDate.Format("01-02-2006 Monday"))
	// in order to build in golang just put go command
	
}