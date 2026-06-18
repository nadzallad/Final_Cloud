package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("auth-service healthy"))
	})

	log.Println("Auth Service running on :5001")
	log.Fatal(http.ListenAndServe(":5001", nil))
}