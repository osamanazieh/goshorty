package api

import (
	"log"
	"net/http"
	"time"

	"github.com/osamaNazieh/shorty/internal/database"
)

func (cfg *Cfg) UpdateURL(w http.ResponseWriter, r *http.Request) {
	shortCode := r.PathValue("short_code")
	
	type Body struct {
		Url string `json:"url"`
	}

	defer r.Body.Close()
	body, err := getJSON[Body](r.Body)

	if err != nil || body.Url == "" {
		respondWithJSON(w, http.StatusBadRequest, struct{ body string } { body: "bad request" })
		return 
	}

	record, err := cfg.DB.UpdateUrl(r.Context(), database.UpdateUrlParams{
		ShortCode: shortCode,
		Url: body.Url,
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		respondWithError(w, err, "something wronge when updateing url")
		return
	}

	hits, err := cfg.DB.GetHits(r.Context(), shortCode)
	if err != nil {
		respondWithJSON(w, http.StatusNotFound, struct{ body string } { body: "there is no associated url with that code"})
		return 
	}
	
	if err := cfg.DB.SetHits(r.Context(), database.SetHitsParams{
		Hits: hits + 1, 
		ShortCode: shortCode,
	}); err != nil {
		respondWithJSON(w, http.StatusBadRequest, struct{}{})
		return 
	}

	if err := respondWithJSON(w, http.StatusOK, record); err != nil {
		log.Fatal(err)
		return
	}
}