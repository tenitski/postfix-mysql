package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/bouk/httprouter"

	log "github.com/sirupsen/logrus"

	"smtpkeeper/db"
	"smtpkeeper/smtp"
)

func NewUpdateUserHandler(repo db.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user smtp.User

		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			log.WithError(err).Warn("Error while processing request")
			respondWithError(w, r, http.StatusBadRequest, err)
			return
		}

		// Email is our primary key, make sure it is same in the passed data as in query string
		user.Email = httprouter.GetParam(r, "email")

		// todo: data validation

		log.WithField("data", user).Debugf("Processing request %s %s", r.Method, r.URL.Path)

		err = repo.Update(user)
		if err != nil {
			if err == sql.ErrNoRows {
				respondWithStatusCode(w, r, http.StatusNotFound)
				return
			} else {
				log.WithError(err).Error("Error while processing request")
				respondWithError(w, r, http.StatusInternalServerError, err)
				return
			}
		}

		respondWithStatusCode(w, r, http.StatusOK)
	}
}
