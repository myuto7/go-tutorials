package main

import (
	"fmt"
	"math"
	"net/http"
	"time"
)

func hello() {
	fmt.Println("Hello, World")
}

func nowTime() {
	fmt.Println("Welcome to the playground")
	fmt.Println("The time is", time.Now())
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "Hello World, %s!", request.URL.Path[1:])
}

func rand() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}

func main() {
	// hello()
	// nowTime()
	rand()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
