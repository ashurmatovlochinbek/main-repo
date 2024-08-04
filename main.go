package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"submodule"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", submodule.HandlerF1)
	err := http.ListenAndServe(":8080", r)

	if err != nil {
		panic(err)
	}
}
