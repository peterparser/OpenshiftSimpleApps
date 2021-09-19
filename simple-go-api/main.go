package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


func main() {

	r := mux.NewRouter()
	r.HandleFunc("/health", HealthHandler)
	http.Handle("/", r)
	log.Printf("Starting Application\nServices:\n/health\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Handlers

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Status OK\n")
}
