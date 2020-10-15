package handler

import (
	"database/sql"
	"net/http"

	"github.com/bouk/httprouter"
	log "github.com/sirupsen/logrus"

	"smtpkeeper/db"
)

func NewDeleteUserHandler(repo db.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("Processing request %s %s", r.Method, r.URL.Path)

		login := httprouter.GetParam(r, "login")

		err := repo.Delete(login)
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
