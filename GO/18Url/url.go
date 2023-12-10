package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://google.com"

func main() {
	fmt.Println("Url it is")
	res, _ := url.Parse(myurl)
	// just give the information about the which is site.
	fmt.Println(res.Scheme)
	fmt.Println(res.Host)
	fmt.Println(res.Path)
	fmt.Println(res.Port())
	fmt.Println(res.RawQuery)
}