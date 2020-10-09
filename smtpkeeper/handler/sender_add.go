package handler

import (
	"encoding/json"
	"net/http"

	"github.com/bouk/httprouter"

	log "github.com/sirupsen/logrus"

	"smtpkeeper/db"
)

func NewAddSenderHandler(repo db.SenderRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var sender string

		err := json.NewDecoder(r.Body).Decode(&sender)
		if err != nil {
			log.WithError(err).Warn("Error while processing request")
			respondWithError(w, r, http.StatusBadRequest, err)
			return
		}

		login := httprouter.GetParam(r, "login")

		// todo: data validation
		// todo: 404 if no user?

		log.WithField("data", sender).Debugf("Processing request %s %s", r.Method, r.URL.Path)

		err = repo.Add(login, sender)
		if err != nil {
			log.WithError(err).Error("Error while processing request")
			respondWithError(w, r, http.StatusInternalServerError, err)
			return
		}

		respondWithStatusCode(w, r, http.StatusOK)
	}
}
