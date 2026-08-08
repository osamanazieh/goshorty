package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/osamaNazieh/shorty/internal/API"
	"github.com/osamaNazieh/shorty/internal/database"
	_ "github.com/lib/pq" 
)




func main() {
	godotenv.Load()
	dbUrl := os.Getenv("DB_URL")
	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatal(err)
	}


	dbQueries := database.New(db)
	cfg := &api.Cfg{
		DB: dbQueries,
	}

	mux := http.NewServeMux()

	mux.HandleFunc("POST /shorten", cfg.GenerateUrl)
	mux.HandleFunc("GET /shorten/{short_code}", cfg.RetrieveOriginalURL)
	mux.HandleFunc("PUT /shorten/{short_code}", cfg.UpdateURL)
	mux.HandleFunc("DELETE /shorten/{short_code}", cfg.DeleteURL)
	mux.HandleFunc("DELETE /shorten/{short_code}/stats", cfg.GetStats)
	server := &http.Server{
		Addr: ":8080", 
		Handler: mux,
	}

	log.Fatal(server.ListenAndServe())
}