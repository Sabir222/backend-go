package server

import (
	"github.com/Sabir222/backend-go/internal/api"
	"net/http"
)

func NewServer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", api.HelloHandler)
	return mux
}
