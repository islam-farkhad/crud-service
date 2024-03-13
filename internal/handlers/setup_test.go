package handlers

import (
	mockrepository "crud-service/internal/pkg/repository/mocks"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/mux"
)

type appFixture struct {
	ctrl     *gomock.Controller
	mockApp  App
	mockRepo mockrepository.MockRepo
}

func setUp(t *testing.T) appFixture {
	ctrl := gomock.NewController(t)
	mockRepo := mockrepository.NewMockRepo(ctrl)
	mockApp := App{
		Router: mux.NewRouter(),
		Repo:   mockRepo,
	}

	return appFixture{
		ctrl:     ctrl,
		mockApp:  mockApp,
		mockRepo: *mockRepo,
	}
}

func (a *appFixture) tearDown() {
	a.ctrl.Finish()
}
