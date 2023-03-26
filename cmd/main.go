package main

import (
	"fmt"
	"log"
	"net/http"

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
	repo := repository.NewRepository(db)
	handlers := transport.NewHandler(repo)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/projects", handlers.Projects)
	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(cfg.HttpServer.Port, nil))
}
