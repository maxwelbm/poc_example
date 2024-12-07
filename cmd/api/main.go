package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"github.com/maxwelbm/pod_example/internal/controller"
	"github.com/maxwelbm/pod_example/internal/repository"
	"github.com/maxwelbm/pod_example/internal/routes"
	"github.com/maxwelbm/pod_example/internal/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, loading system environment variables")
		return
	}

	db := repository.NewMeliDB()
	serv := service.NewServiceProducts(&db)
	ctrl := controller.NewControllerProducts(&serv)

	rt := chi.NewRouter()
	routes.RegisterProductRoutes(rt, ctrl)

	if err := http.ListenAndServe(":8080", rt); err != nil {
		log.Println(err)
		return
	}
}
