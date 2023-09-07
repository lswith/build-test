package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

var who string

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello %s! %s", who, time.Now())
}

func main() {
	http.HandleFunc("/", greet)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	http.ListenAndServe(":"+port, nil)
}
