package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/goodbye", goodbye)
	http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	w.Write([]byte("hello!"))
}

func goodbye(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r)
	w.Write([]byte("goodbye!"))
}
