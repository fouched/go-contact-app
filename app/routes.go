package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Get("/", Home)
	mux.Get("/contacts", ContactsList)
	mux.Get("/contacts/new", ContactsAdd)
	mux.Post("/contacts/new", ContactsNew)

	return mux
}
