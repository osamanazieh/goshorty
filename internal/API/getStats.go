package api

import (
	"fmt"
	"log"
	"net/http"
)

func (cfg *Cfg) GetStats(w http.ResponseWriter, r *http.Request) {
	shortCode := r.PathValue("short_code")

	record, err := cfg.DB.GetUrl(r.Context(), shortCode)
	if err != nil {
		respondWithJSON(w, http.StatusNotFound, struct{}{})
		fmt.Println(err)
		return
	}
	
	if err := respondWithJSON(w, http.StatusOK, record); err != nil {
		log.Fatal(err)
	}
}