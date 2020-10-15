package handler

import (
	"net/http"

	"github.com/bouk/httprouter"
	log "github.com/sirupsen/logrus"

	"smtpkeeper/db"
)

func NewGetSendersHandler(repo db.SenderRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Debugf("Processing request %s %s", r.Method, r.URL.Path)

		login := httprouter.GetParam(r, "login")

		// todo: 404 if no user?

		senders, err := repo.Get(login)
		if err != nil {
			log.WithError(err).Error("Error while processing request")
			respondWithError(w, r, http.StatusInternalServerError, err)
			return
		}

		if len(senders) == 0 {
			// A bit of a hack to force empty json array instead of null
			senders = make([]string, 0)
		}

		respondWithData(w, r, http.StatusOK, senders)
	}
}
