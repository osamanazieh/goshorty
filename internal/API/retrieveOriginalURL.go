package api

import (
	"log"
	"net/http"

	"github.com/osamaNazieh/shorty/internal/database"
)

func (cfg *Cfg)RetrieveOriginalURL(w http.ResponseWriter, r *http.Request) {
	
	shortCode := r.PathValue("short_code")
	
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
	record, err := cfg.DB.GetUrl(r.Context(), shortCode)
	if err != nil {
		respondWithJSON(w, http.StatusNotFound, struct{ body string } { body: "there is no associated url with that code"})
		return 
	}
	
	if err := respondWithJSON(w, http.StatusOK, record); err != nil {
		log.Fatal(err)
	}

}