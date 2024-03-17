package main

import (
	"github.com/fouched/go-contact-app/internal/apix"
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
	mux.Get("/help", handlers.Instance.Help)
	mux.Get("/settings", handlers.Instance.Settings)

	mux.Route("/contacts", func(r chi.Router) {
		r.Get("/", handlers.Instance.ContactsList)
		r.Get("/new", handlers.Instance.ContactsNewGet)
		r.Post("/new", handlers.Instance.ContactsNew)
		r.Get("/{id}", handlers.Instance.ContactsView)
		r.Get("/{id}/edit", handlers.Instance.ContactsEditGet)
		r.Post("/{id}/edit", handlers.Instance.ContactsEditPost)
		r.Delete("/{id}/delete", handlers.Instance.ContactsDelete)
	})

	mux.Get("/playground", handlers.Instance.PlaygroundGet)

	mux.Post("/apix/contacts", apix.Instance.ContactsList)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
