package handler

import (
	"encoding/json"
	"net/http"

	log "github.com/sirupsen/logrus"

	"smtpkeeper/db"
	"smtpkeeper/smtp"
)

func NewCreateUserHandler(repo db.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user smtp.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.WithError(err).Warn("Error while processing request")
			respondWithError(w, r, http.StatusBadRequest, err)
			return
		}

		// todo: data validation

		log.WithField("data", user).Debugf("Processing request %s %s", r.Method, r.URL.Path)

		err = repo.Create(user)
		if err != nil {
			log.WithError(err).Error("Error while processing request")
			respondWithError(w, r, http.StatusInternalServerError, err)
			return
		}

		respondWithStatusCode(w, r, http.StatusOK)
	}
}
