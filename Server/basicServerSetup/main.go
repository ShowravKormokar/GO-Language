package main

// Every Go program starts with a package declaration.
// "main" is special: it tells Go this is an executable program (not a library).
// The compiler looks for a main() function inside package main as the entry point.

import (
	"fmt"      // Provides formatted I/O functions (Println, Fprintln, etc.)
	"net/http" // Standard library for building HTTP servers and clients
)

// helloHandler is a handler function for the "/hello" route.
// Parameters:
// - rw http.ResponseWriter: used to send data back to the client (response).
// - req *http.Request: represents the incoming HTTP request (method, headers, body, etc.).
func helloHandler(rw http.ResponseWriter, req *http.Request) {
	// Fprintln writes a string plus newline to the response writer.
	fmt.Fprintln(rw, "Hello world! This is basic go server.")
	fmt.Printf("SUCCESS: %s %s -> Status %d\n", req.Method, req.URL.Path, http.StatusOK)
}

// aboutHandler is another handler function for the "/about" route.
func aboutHandler(rw http.ResponseWriter, req *http.Request) {
	// Fprint writes a string (without newline) to the response writer.
	fmt.Fprint(rw, "I am Showrav Kormokar.")
	fmt.Printf("SUCCESS: %s %s -> Status %d\n", req.Method, req.URL.Path, http.StatusOK)
}

func main() {
	// Create a new ServeMux (multiplexer).
	// A ServeMux maps URL paths (like "/hello") to handler functions.
	mux := http.NewServeMux()

	// Register route handlers with the mux.
	// HandleFunc takes a path and a function with signature (ResponseWriter, *Request).
	mux.HandleFunc("/hello", helloHandler)
	mux.HandleFunc("/about", aboutHandler)

	// Print a message to the console so you know the server started.
	fmt.Println("Server running successfully on port 3000.")

	// Start the HTTP server on port 3000, using mux to handle requests.
	// ListenAndServe blocks forever until the server stops or errors.
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		// If the server fails to start (e.g., port in use), print the error.
		fmt.Println("Couldn't connect to server:", err)
		return
	}
}
