package accounts

import (
	"net/http"
	"yatter-backend-go/app"

	"github.com/go-chi/chi/v5"
)

// Implementation of handler
type handler struct {
	app *app.App
}

// Create Handler for `/v1/accounts/`
func NewRouter(app *app.App) http.Handler {
	r := chi.NewRouter()

	h := &handler{app: app}
	r.Post("/", h.Create)
	r.Get("/{username}", h.Find)

	return r
}
