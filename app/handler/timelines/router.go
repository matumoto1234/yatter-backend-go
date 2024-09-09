package timelines

import (
	"net/http"
	"yatter-backend-go/app/usecase"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	timelineUsecase usecase.Timeline
}
// Create Handler for `/v1/timelines/`
func NewRouter(tu usecase.Timeline) http.Handler {
	r := chi.NewRouter()

	h := &handler{timelineUsecase: tu}
	r.Get("/public", h.GetPublic)

	return r

}