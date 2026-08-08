package api

import (
	"log"
	"net/http"
)

func (cfg *Cfg) DeleteURL(w http.ResponseWriter, r *http.Request) {
	shortCode := r.PathValue("short_code")

	if _, err := cfg.DB.DeleteUrl(r.Context(), shortCode); err != nil {
		respondWithJSON(w, http.StatusNotFound, struct{ body string } { body: "bad Request Body"})
		return 
	}

	if err := respondWithJSON(w, http.StatusNoContent, ""); err != nil {
		log.Fatal(err)
	}
}