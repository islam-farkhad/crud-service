package handlers

import (
	"fmt"
	"homework-3/internal/pkg/repository"
	"homework-3/internal/pkg/repository/postgresql"

	"github.com/gorilla/mux"
)

// App represents the main application structure.
type App struct {
	Router *mux.Router
	Repo   repository.Repo
}

// NewApp creates a new instance of the App type with the provided router and repository.
// It initializes the routes for the application.
func NewApp(router *mux.Router, repo *postgresql.Repo, prefix string) App {
	app := App{
		Router: router,
		Repo:   repo,
	}
	app.initializeRoutes(prefix)
	return app
}

func (app *App) initializeRoutes(prefix string) {
	app.Router.HandleFunc(fmt.Sprintf("%s/post", prefix), app.HandleCreatePost).Methods("POST")
	app.Router.HandleFunc(fmt.Sprintf("%s/post/{id:[\\S]*}", prefix), app.HandleGetPostByID).Methods("GET")
	app.Router.HandleFunc(fmt.Sprintf("%s/post/{id:[0-9]+}/comment", prefix), app.HandleCreateComment).Methods("POST")
	app.Router.HandleFunc(fmt.Sprintf("%s/post", prefix), app.HandleUpdatePost).Methods("PUT")
	app.Router.HandleFunc(fmt.Sprintf("%s/post/{id:[\\S]*}", prefix), app.HandleDeletePostByID).Methods("DELETE")
}
