package routes

import (
	"github.com/go-chi/chi"
	"github.com/maxwelbm/pod_example/internal/controller"
	"github.com/maxwelbm/pod_example/internal/middleware"
)

func RegisterProductRoutes(r chi.Router, productCtrl *controller.ControllerProduct) {
	r.Route("/products", func(r chi.Router) {
		r.Use(middleware.Auth)
		r.Post("/", productCtrl.Create)
		r.Get("/", productCtrl.GetAll)
		r.Get("/{id}", productCtrl.GetById)
		r.Get("/search", productCtrl.Search)
	})
}
