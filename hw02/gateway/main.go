package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /ping", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprint(w, "pong")
	})

	log.Println("Gateway service started on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
