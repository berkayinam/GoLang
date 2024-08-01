package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hii guyys")	
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("localhost devrede ? ? omgg baya iyiii")
	http.ListenAndServe(":8080", nil)
}