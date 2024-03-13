package app

import (
	"crud-service/internal/handlers"
	"crud-service/internal/pkg/db"
	"crud-service/internal/pkg/repository/postgresql"
	"net/http"

	"github.com/gorilla/mux"
)

// NewTestApp creates test app with prefix /test in all routes
func NewTestApp(database db.DBops) handlers.App {
	testApp := handlers.NewApp(mux.NewRouter(), postgresql.NewRepo(database), "/test")
	http.Handle("/", testApp.Router)
	return testApp
}
