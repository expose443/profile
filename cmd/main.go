package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/with-insomnia/profile/internal/config"
	"github.com/with-insomnia/profile/internal/repository"
	"github.com/with-insomnia/profile/internal/transport"
)

func main() {
	cfg, err := config.Init("config.json")
	if err != nil {
		fmt.Println(err)
	}
	db, err := repository.InstancePostgres(&cfg.PostgresInfo)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	repo := repository.NewRepository(db)
	handlers := transport.NewHandler(repo)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Post("/login", handlers.Login)
	r.Post("/project", handlers.CreateProject)
	r.Get("/project", handlers.GetProjects)

	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(cfg.HttpServer.Port, r))
}
