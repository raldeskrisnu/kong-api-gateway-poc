package main

import (
	"fmt"
	"log"
	"net/http"
)

func protectedHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Access granted: Protected resource")
}

func apiHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Access Api Successfully")
}

func main() {
	http.HandleFunc("/protected", protectedHandler)
	http.HandleFunc("/api", apiHandler)

	fmt.Println("Server running on port 8090...")
	log.Fatal(http.ListenAndServe(":8090", nil))
}
