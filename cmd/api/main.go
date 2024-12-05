package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/maxwelbm/pod_example/internal/controller"
	"github.com/maxwelbm/pod_example/internal/repository"
	"github.com/maxwelbm/pod_example/internal/service/model"
	"golang.org/x/text/cmd/gotext/examples/extract_http/pkg"
)

func main() {
	ok, err := pkg.Connection("", "", "", "", "")
	if err != nil {
		panic(err)
	}

	db := repository.NewMeliDB(ok)
	serv := model.NewServiceProducts(db)
	ctrl := controller.NewControllerProducts(serv)

	rt := chi.NewRouter()

	rt.Route("/products", func(r chi.Router) {
		r.Post("/", ctrl.Create)
		r.Get("/", ctrl.GetAll)
		r.Get("/{id}", ctrl.GetById)
		r.Get("/search", ctrl.Search)
	})

	if err := http.ListenAndServe(":8080", rt); err != nil {
		panic(err)
	}
}
