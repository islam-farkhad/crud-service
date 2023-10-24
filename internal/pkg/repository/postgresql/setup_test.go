package postgresql

import (
	"testing"

	mockdatabase "homework-3/internal/pkg/db/mocks"
	"homework-3/internal/pkg/repository"

	"github.com/golang/mock/gomock"
)

type articlesRepoFixture struct {
	ctrl   *gomock.Controller
	repo   repository.Repo
	mockDb *mockdatabase.MockDBops
}

func setUp(t *testing.T) articlesRepoFixture {
	ctrl := gomock.NewController(t)
	mockDb := mockdatabase.NewMockDBops(ctrl)
	repo := NewRepo(mockDb)
	return articlesRepoFixture{
		ctrl:   ctrl,
		repo:   repo,
		mockDb: mockDb,
	}
}

func (a *articlesRepoFixture) tearDown() {
	a.ctrl.Finish()
}
