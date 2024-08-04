package server

import (
	"github.com/Sabir222/backend-go/internal/api"
	"net/http"
	"path/filepath"
)

func NewServer() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", api.HelloHandler)
	staticDir := http.Dir(filepath.Join("templates"))
	mux.Handle("/", http.FileServer(staticDir))
	return mux
}
