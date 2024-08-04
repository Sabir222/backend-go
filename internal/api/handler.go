package api

import (
	"log"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "text/html")
	_, err := w.Write([]byte("<p>Hello from the server!</p>"))
	if err != nil {
		log.Printf("Error writing response : %v", err)
	}
}
