package handler

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type errorResponse struct {
	Error string `json:"error"`
}

func respondWithData(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	body, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	_, err = w.Write(body)
	if err != nil {
		log.WithError(err).Error("Error while sending response")
	}
}

func respondWithStatusCode(w http.ResponseWriter, r *http.Request, code int) {
	w.WriteHeader(code)
}

func respondWithError(w http.ResponseWriter, r *http.Request, code int, responseErr error) {
	response := errorResponse{
		Error: responseErr.Error(),
	}

	respondWithData(w, r, code, response)
}
