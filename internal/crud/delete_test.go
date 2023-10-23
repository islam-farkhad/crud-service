package crud

import (
	"context"
	"fmt"
	"homework-3/internal/pkg/repository"
	"homework-3/tests/states"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDeletePostByID(t *testing.T) {
	t.Parallel()

	var (
		ctx = context.Background()
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		// arrange
		s := setUp(t)
		defer s.tearDown()

		s.mockRepo.EXPECT().DeletePostByID(gomock.Any(), gomock.Any()).Return(true, nil)

		// act
		data, status := s.mockApp.DeletePostByID(ctx, states.Post1ID)

		// assert
		require.Equal(t, http.StatusOK, status)
		require.Nil(t, data)
	})

	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		t.Run("not existing id", func(t *testing.T) {
			t.Parallel()

			// arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockRepo.EXPECT().DeletePostByID(gomock.Any(), states.Post2ID).Return(false, repository.ErrObjectNotFound)

			// act
			data, status := s.mockApp.DeletePostByID(ctx, states.Post2ID)

			// assert
			require.Equal(t, http.StatusNotFound, status)
			require.Contains(t, string(data), repository.ErrObjectNotFound.Error())
		})

		t.Run("deleting error", func(t *testing.T) {
			t.Parallel()

			// arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockRepo.EXPECT().DeletePostByID(gomock.Any(), gomock.Any()).Return(false, assert.AnError)

			// act
			data, status := s.mockApp.DeletePostByID(ctx, states.Post2ID)

			// assert
			require.Equal(t, http.StatusInternalServerError, status)
			require.Contains(t, string(data), fmt.Sprintf("deleting post by postID error: %v", assert.AnError))
		})
	})
}
