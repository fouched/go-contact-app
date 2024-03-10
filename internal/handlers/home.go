package handlers

import "net/http"

// Home is the home page handler
func (m *HandlerConfig) Home(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/contacts", http.StatusSeeOther)
}
