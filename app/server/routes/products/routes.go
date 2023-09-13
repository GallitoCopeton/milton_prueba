package products

import (
	"challenge/handlers"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const RouteName = "/products"

func Routes() http.Handler {
	r := chi.NewRouter()
	h := handlers.NewProductHandler()

	r.Get(RouteName, h.GetProducts())
	// Other handlers...

	return r
}
