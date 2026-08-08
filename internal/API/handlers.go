package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)


func respondWithError(w http.ResponseWriter, err error, msg string) error {
	w.WriteHeader(http.StatusInternalServerError)
	w.Header().Set("Content-Type", "application/json")
	fmt.Println(msg)	
	if _, err := w.Write([]byte(err.Error())); err != nil {
		return err 
	} 
	return nil 
}


func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) error {
	w.WriteHeader(code)
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		return err 
	}
	if code == http.StatusNoContent {
		return nil 
	}
	if _, err := w.Write(dat); err != nil {
		return err 
	} 
	return nil 
}


func getJSON[T any](body io.ReadCloser) (T, error) {
	var info T 
	decoder := json.NewDecoder(body)
	if err := decoder.Decode(&info); err != nil {
		fmt.Println(err)
		var zero T
		return zero, err
	}
	return info, nil 
}