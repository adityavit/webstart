package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

// Define all Handlers
func addHandler(r *mux.Router) {
	r.HandleFunc("/", mainHandler)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello World")
}
