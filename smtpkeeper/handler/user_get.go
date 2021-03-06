package handler

import (
	"database/sql"
	"net/http"

	"github.com/bouk/httprouter"
	log "github.com/sirupsen/logrus"

	"smtpkeeper/db"
)

func NewGetUserHandler(repo db.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("Processing request %s %s", r.Method, r.URL.Path)

		login := httprouter.GetParam(r, "login")

		user, err := repo.Get(login)
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

		respondWithData(w, r, http.StatusOK, user)
	}
}

func NewGetUsersHandler(repo db.UserRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("Processing request %s %s", r.Method, r.URL.Path)

		users, err := repo.GetAll()
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

		respondWithData(w, r, http.StatusOK, users)
	}
}
