package app

import (
	"fmt"
	"github.com/paw1a/transaction-service/internal/db"
	"github.com/paw1a/transaction-service/internal/handler"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"time"
)

func Run() {
	log.Info("application startup...")

	fmt.Printf("%s\n", os.Getenv("PORT"))

	conn, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	log.Info("database connection initialized")

	server := &http.Server{
		Handler:      handler.NewHandler(conn).Init(),
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Infof("server started on port %s", os.Getenv("PORT"))

	log.Fatal(server.ListenAndServe())
}
