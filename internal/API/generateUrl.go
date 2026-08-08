package api

import (
	"crypto/rand"
	"encoding/base64"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/osamaNazieh/shorty/internal/database"
)

func generateRand() (string, error) {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.RawStdEncoding.EncodeToString(b), nil
}



func (cfg *Cfg)GenerateUrl(w http.ResponseWriter, r *http.Request) {	
	
	
	newUrl, err := generateRand()
	if err != nil {
		respondWithError(w, err, "at generateRand")
		return 
	}
	
	defer r.Body.Close()

	type Body struct{
		Url string `json:"url"` 
	}
	body, err := getJSON[Body](r.Body)
	
	if err != nil || body.Url == "" {
		respondWithJSON(w, http.StatusBadRequest, struct{ body string } { body: "Invlaid Request" })
	}

	record, err := cfg.DB.CreateUrl(r.Context(), database.CreateUrlParams{
		ID: uuid.New(), 
		Url: body.Url,
		ShortCode: newUrl,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithJSON(w, http.StatusBadRequest, struct {
			body string
		} {
			body: "the request was invalid",
		})
		return 
	}

	if err := respondWithJSON(w, http.StatusCreated, record); err != nil {
		log.Fatal(err)
	}
}