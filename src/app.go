package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	address = ":8080"
	message = os.Getenv("MSG")
)

func handler(w http.ResponseWriter, r *http.Request) {
	log.Print("🎯 Request received")
	fmt.Fprintf(w, "<h1>Hello %s!</h1>", message)
}

func main() {
	log.Printf("🧩 Message is \"%s\"", message)
	log.Printf("👋 Listening at %s", address)

	http.HandleFunc("/", handler)

	if err := http.ListenAndServe(address, nil); err != http.ErrServerClosed {
		log.Fatalf("Server error: %v", err)
	}
}
