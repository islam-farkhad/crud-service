package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"homework-3/internal/handlers"
	"homework-3/internal/pkg/repository/postgresql"
	"homework-3/internal/utils"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	database := utils.ConnectDB(ctx)
	defer database.GetConnectionsPool(ctx).Close()

	app := handlers.NewApp(mux.NewRouter(), postgresql.NewRepo(database), "")

	http.Handle("/", app.Router)

	srv := http.Server{
		Addr:              utils.GetAPIPort(),
		WriteTimeout:      1 * time.Second,
		ReadHeaderTimeout: 1 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
