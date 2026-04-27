package main

import (
	"log"
	"net/http"
	"os"

	"gladiatortravel/internal/handler"
	"gladiatortravel/internal/repository"
	"gladiatortravel/internal/usecase"
)

func main() {
	cfg := loadConfig()

	repo, err := repository.NewPostgresRepository(cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}
	defer repo.Close()

	uc := usecase.NewService(repo)
	h := handler.NewHTTPHandler(uc)

	log.Printf("gladiatortravel api listening on %s", cfg.Port)
	if err := http.ListenAndServe(cfg.Port, h.Router()); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

type config struct {
	Port        string
	DatabaseURL string
}

func loadConfig() config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = "postgres://postgres:postgres@localhost:5432/gladiatortravel?sslmode=disable"
	}
	return config{
		Port:        ":" + port,
		DatabaseURL: dbURL,
	}
}
