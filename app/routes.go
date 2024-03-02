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
	mux.Get("/contacts", ContactsList)
	mux.Get("/contacts/new", ContactsView)
	mux.Post("/contacts/new", ContactsAdd)

	return mux
}
