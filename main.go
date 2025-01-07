// main.go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Handle requests to the root URL ('/')
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Respond with a simple message
		fmt.Fprintf(w, "Server is running!")
	})

	// Start the server on port 8080
	fmt.Println("Server is starting on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
