package main

import (
	"fmt"
	"net/http"
)

func main() {
	request1()
}

func homePage(response http.ResponseWriter, request *http.Request) {
	fmt.Printf("Welcome to the Go Web API!")
	fmt.Printf("Endpoint Hit: homePage")
}

func aboutMe(response http.ResponseWriter, request *http.Request) {
	who := "Salman"

	fmt.Fprintf(response, "A little bit about Salman...")
	fmt.Println("Endpoint Hit: ", who)
}

func request1() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/aboutme", aboutMe)

	http.ListenAndServe(":8080", nil)
}
