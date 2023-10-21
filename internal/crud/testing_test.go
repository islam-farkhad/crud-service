package crud

import (
	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
	mock_repository "homework-3/internal/pkg/repository/mocks"
	"testing"
)

type appFixture struct {
	ctrl     *gomock.Controller
	mockApp  App
	mockRepo mock_repository.MockRepo
}

func setUp(t *testing.T) appFixture {
	ctrl := gomock.NewController(t)
	mockRepo := mock_repository.NewMockRepo(ctrl)
	mockApp := App{
		Router: mux.NewRouter(),
		Repo:   mockRepo,
	}

	mockApp.Router.HandleFunc("/post/{id:[\\S]*}", mockApp.DeletePostByID).Methods("DELETE")
	mockApp.Router.HandleFunc("/post/{id:[\\S]*}", mockApp.GetPostByID).Methods("GET")
	mockApp.Router.HandleFunc("/post/{id:[0-9]+}/comment", mockApp.CreateComment).Methods("POST")

	return appFixture{
		ctrl:     ctrl,
		mockApp:  mockApp,
		mockRepo: *mockRepo,
	}
}

func (a *appFixture) tearDown() {
	a.ctrl.Finish()
}
