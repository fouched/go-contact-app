package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(SessionLoad)

	mux.Get("/", Home)

	mux.Route("/contacts", func(r chi.Router) {
		r.Get("/", ContactsList)
		r.Get("/new", ContactsView)
		r.Post("/new", ContactsAdd)
		r.Get("/{id}", ContactsById)
	})

	return mux
}
