package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spending-tracking/db"
	"github.com/spending-tracking/handlers"
)

func main() {
	r := chi.NewRouter()

	db.OpenDB()

	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World!"))
	})

	r.Get("/account", handlers.AccountHandler)
	http.ListenAndServe(":3000", r)
}
