package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"gladiatortravel/internal/handler"
	"gladiatortravel/internal/repository"
	"gladiatortravel/internal/usecase"
)

func main() {
	cfg := loadConfig()

	repo, err := connectWithRetry(cfg.DatabaseURL, 20, 2*time.Second)
	if err != nil {
		log.Fatalf("failed to connect db after retries: %v", err)
	}
	defer repo.Close()

	uc := usecase.NewService(repo)
	h := handler.NewHTTPHandler(uc)

	log.Printf("gladiatortravel api listening on %s", cfg.Port)
	if err := http.ListenAndServe(cfg.Port, h.Router()); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}

func connectWithRetry(databaseURL string, attempts int, delay time.Duration) (*repository.PostgresRepository, error) {
	var lastErr error
	for i := 1; i <= attempts; i++ {
		repo, err := repository.NewPostgresRepository(databaseURL)
		if err == nil {
			return repo, nil
		}
		lastErr = err
		log.Printf("db connection attempt %d/%d failed: %v", i, attempts, err)
		time.Sleep(delay)
	}
	return nil, lastErr
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
