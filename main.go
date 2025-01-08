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
		fmt.Fprintf(w, "Server is running on 4000!")
	})

	// Start the server on port 4000
	fmt.Println("Server is starting on port 4000...")
	if err := http.ListenAndServe(":4000", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
