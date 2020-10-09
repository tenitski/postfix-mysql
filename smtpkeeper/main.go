package main

import (
	"net/http"
	"os"
	"smtpkeeper/db"
	"smtpkeeper/handler"

	"github.com/bouk/httprouter"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Get database server details
	var dsn string
	if len(os.Args) > 1 {
		dsn = os.Args[1]
	}
	if dsn == "" {
		log.Fatalln("Database server not provided")
	}

	// Allow to specify log level via ENV var
	if ll := os.Getenv("LOG_LEVEL"); ll != "" {
		logLevel, err := log.ParseLevel(ll)
		if err != nil {
			log.WithError(err).Fatal("Unable to parse log level")
		}
		log.SetLevel(logLevel)
	}

	conn, err := db.New(dsn)
	if err != nil {
		log.WithError(err).Fatal("Unable to init store")
	}

	userRepo := db.NewUserRepository(conn)

	// Configure routes
	router := httprouter.New()
	router.GET("/user/:email", handler.NewGetUserHandler(userRepo))
	router.GET("/users", handler.NewGetUsersHandler(userRepo))
	router.POST("/users", handler.NewCreateUserHandler(userRepo))
	router.PUT("/user/:email", handler.NewUpdateUserHandler(userRepo))
	router.DELETE("/user/:email", handler.NewDeleteUserHandler(userRepo))

	// Start server
	addr := ":18080"
	log.WithField("addr", addr).Info("Starting server")
	log.Fatal(http.ListenAndServe(addr, router))
}
