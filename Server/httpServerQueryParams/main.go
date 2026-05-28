package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func infoHandler(rw http.ResponseWriter, req *http.Request) {
	name := req.URL.Query().Get("name")
	salary := req.URL.Query().Get("salary")

	msg := fmt.Sprintf("Employee Info -> Name: %s, Salary: %s", name, salary)
	rw.WriteHeader(http.StatusOK)
	fmt.Fprintln(rw, msg)

	fmt.Printf("SUCCESS: %s %s -> Status: %d\n", req.Method, req.URL.Path, http.StatusOK)
}

func htmlHandler(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/html")
	html := `
	<!DOCTYPE html>
	<html>
	<head><title>Go Server</title></head>
	<body>
		<h1>Hello from GO Server</h1>
		<p>This is a sample HTML response</p>
	</body>
	</html>
	`
	rw.WriteHeader(http.StatusOK)
	fmt.Fprint(rw, html)

	fmt.Printf("SUCCESS: %s %s -> Status %d\n", req.Method, req.URL.Path, http.StatusOK)
}

func pdfHandler(rw http.ResponseWriter, req *http.Request) {
	file, err := os.Open("server.pdf")
	if err != nil {
		http.Error(rw, "File not found", http.StatusNotFound)
		fmt.Printf("ERROR: %s %s -> Status %d\n", req.Method, req.URL.Path, http.StatusNotFound)
		return
	}

	defer file.Close()

	rw.Header().Set("Content-Type", "application/pdf")
	rw.WriteHeader(http.StatusOK)

	// Copy file contents into response
	_, err = io.Copy(rw, file)
	if err != nil {
		fmt.Printf("ERROR streaming file: %v\n", err)
	}

	fmt.Printf("SUCCESS: %s %s -> Status %d (PDF served)\n", req.Method, req.URL.Path, http.StatusOK)
}

func main() {
	mux := http.NewServeMux()

	// Register Handlers
	mux.HandleFunc("/info", infoHandler)
	mux.HandleFunc("/html", htmlHandler)
	mux.HandleFunc("/pdf", pdfHandler)

	fmt.Println("Server running successfully on port 3000.")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Println("Couldn't connect to server:", err)
	}
}
