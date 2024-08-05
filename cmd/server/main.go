package main

import (
	"log"
	"net/http"

	"github.com/Sabir222/backend-go/internal/server"
)

func main() {
	srv := server.NewServer()
	log.Printf("Server running on port 8080")

	if err := http.ListenAndServe(":8080", srv); err != nil {
		log.Fatalf("Server failed to run : %v", err)
	}
	//log.Fatal(http.ListenAndServe(":8080", srv))
}
