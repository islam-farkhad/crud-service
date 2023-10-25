package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"

	"homework-3/internal/crud"
	"homework-3/internal/pkg/db"
	"homework-3/internal/pkg/repository/postgresql"
	"homework-3/internal/utils"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	database, err := db.NewDB(ctx, db.GetDBConnectionString())
	if err != nil {

		println(db.GetDBConnectionString())
		println("AAAAA")

		log.Fatal(err)
	}
	defer database.GetConnectionsPool(ctx).Close()

	app := crud.NewApp(mux.NewRouter(), postgresql.NewRepo(database), "")

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
