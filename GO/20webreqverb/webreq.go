package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("Welcome to this")
	PerfromGetRequest()
}

func PerfromGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)

	checkError(err)
	defer response.Body.Close()
	var responseString strings.Builder
	if response.StatusCode == 200 {
		content, _ := io.ReadAll(response.Body)
		byteCount, _ := responseString.Write(content)

		if(byteCount != 0){
			fmt.Println(responseString.String())

		}


	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
