package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	fmt.Printf("This is a server")
	http.HandleFunc("/",handler)
	log.Fatal(http.ListenAndServe(":8080",nil)) //This defines a server
}

func handler(w http.ResponseWriter, r *http.Request){
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w,nil)
}