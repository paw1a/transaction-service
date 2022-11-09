package app

import (
	"fmt"
	"github.com/paw1a/transaction-service/internal/db"
	"github.com/paw1a/transaction-service/internal/handler"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"time"
)

func Run() {
	log.Info("application startup...")
	log.Info("logger initialized")

	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
	log.Info("config created")

	log.Info("services, repositories and handlers initialized")

	conn, err := db.NewDB()
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Handler:      handler.NewHandler(conn).Init(),
		Addr:         fmt.Sprintf(":%s", viper.GetString("port")),
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Infof("server started on port %s", os.Getenv("PORT"))

	log.Fatal(server.ListenAndServe())
}
