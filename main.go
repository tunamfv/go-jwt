package main

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	log.Fatal(http.ListenAndServe("1010", r))
}