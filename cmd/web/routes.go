package main

import (
	"github.com/fouched/go-contact-app/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Instance.Home)

	mux.Route("/contacts", func(r chi.Router) {
		r.Get("/", handlers.Instance.ContactsList)
		r.Get("/new", handlers.Instance.ContactsAddNew)
		r.Post("/new", handlers.Instance.ContactsAdd)
		r.Get("/{id}", handlers.Instance.ContactsById)
	})

	return mux
}
