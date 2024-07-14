package main

import (
	"fmt"
	"net/http"

	sava "ascii-art-web/handlers"
)

func main() {
	http.HandleFunc("/", sava.Handler)
	http.HandleFunc("/ascii-art", sava.HandleasciiArt)
	fmt.Println("Server is listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}
