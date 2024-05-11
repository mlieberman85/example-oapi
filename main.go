package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mlieberman85/example-oapi/api"
)

// optional code omitted

type Server struct{}

func NewServer() Server {
	return Server{}
}

func (Server) GetFoo(w http.ResponseWriter, r *http.Request) {
	resp := "baz"

	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}

func main() {
	// create a type that satisfies the `api.ServerInterface`, which contains an implementation of every operation from the generated code
	server := api.NewServer()

	r := http.NewServeMux()

	// get an `http.Handler` that we can use
	h := api.HandlerFromMux(server, r)

	s := &http.Server{
		Handler: h,
		Addr:    "0.0.0.0:8080",
	}

	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
