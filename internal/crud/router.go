package crud

import (
	"homework-3/internal/pkg/repository/postgresql"

	"github.com/gorilla/mux"
)

// App represents the main application structure.
type App struct {
	Router *mux.Router
	Repo   *postgresql.Repo
}

// NewApp creates a new instance of the App type with the provided router and repository.
// It initializes the routes for the application.
func NewApp(router *mux.Router, repo *postgresql.Repo) App {
	app := App{
		Router: router,
		Repo:   repo,
	}
	app.initializeRoutes()
	return app
}

func (app *App) initializeRoutes() {
	app.Router.HandleFunc("/post", app.CreatePost).Methods("POST")
	app.Router.HandleFunc("/post/{id:[\\S]*}", app.GetPostByID).Methods("GET")
	app.Router.HandleFunc("/post/{id:[0-9]+}/comment", app.CreateComment).Methods("POST")
	app.Router.HandleFunc("/post", app.UpdatePost).Methods("PUT")
	app.Router.HandleFunc("/post/{id:[0-9]+}", app.DeletePostByID).Methods("DELETE")
}
