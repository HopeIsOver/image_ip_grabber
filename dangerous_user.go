package main

import (
	"dangerous_user/handler"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Handle requests to "/image"
	http.HandleFunc("/", handler.ImageHandler)

	// Start the server on port 8080
	fmt.Println("Server listening on :8080...")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
