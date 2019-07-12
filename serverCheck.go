package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// For this case, we will always pipe "Hello World" into the response writer
	fmt.Fprintf(w, "Listening, Say whatever you want to say budd!")
}

func main() {

	http.HandleFunc("/", handler)
	// http.HandleFunc("/user", userShow)
	// http.HandleFunc("/admin", adminShow)
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
