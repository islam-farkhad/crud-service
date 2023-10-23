package crud

import (
	"context"
	"encoding/json"
	"fmt"
	"homework-3/internal/pkg/repository"
	"homework-3/tests/fixtures"
	"homework-3/tests/states"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPostByID(t *testing.T) {
	t.Parallel()

	var (
		ctx = context.Background()
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		// arrange
		s := setUp(t)
		defer s.tearDown()

		s.mockRepo.EXPECT().GetPostByID(gomock.Any(), states.Post1ID).Return(fixtures.BuildPost().Valid().P(), nil)
		s.mockRepo.EXPECT().GetCommentsByPostID(gomock.Any(), states.Post1ID).Return(
			[]repository.Comment{
				fixtures.BuildComment().Valid().V(),
				fixtures.BuildComment().Valid().ID(states.Comment2ID).Content(states.Comment2Content).V(),
			},
			nil)

		// act
		result, statusCode := s.mockApp.GetPostByID(ctx, states.Post1ID)

		// assert
		require.Equal(t, http.StatusOK, statusCode)

		var response getPostByIDResponse
		err := json.Unmarshal(result, &response)
		require.NoError(t, err)

		assert.Equal(t, states.Post1ID, response.Post.ID)
		assert.Equal(t, states.Post1Content, response.Post.Content)
		assert.Equal(t, states.Post1Likes, response.Post.Likes)
		assert.Equal(t, states.Post1CreatedAt, response.Post.CreatedAt)
		assert.Len(t, response.Comments, 2)

		assert.Equal(t, states.Comment1ID, response.Comments[0].ID)
		assert.Equal(t, states.Comment2ID, response.Comments[1].ID)

		assert.Equal(t, states.Post1ID, response.Comments[0].PostID)
		assert.Equal(t, states.Post1ID, response.Comments[1].PostID)

		assert.Equal(t, states.Comment1Content, response.Comments[0].Content)
		assert.Equal(t, states.Comment2Content, response.Comments[1].Content)
	})

	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		t.Run("post not found", func(t *testing.T) {
			t.Parallel()

			// arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockRepo.EXPECT().GetPostByID(gomock.Any(), states.Post2ID).Return(nil, repository.ErrObjectNotFound)

			// act
			result, statusCode := s.mockApp.GetPostByID(ctx, states.Post2ID)

			// assert
			require.Equal(t, http.StatusNotFound, statusCode)
			require.Contains(t, string(result), fmt.Sprintf("postRepo with ID=%d not found, err: %v", states.Post2ID, repository.ErrObjectNotFound))
		})

		t.Run("internal error", func(t *testing.T) {
			t.Parallel()

			// arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockRepo.EXPECT().GetPostByID(gomock.Any(), states.Post2ID).Return(nil, assert.AnError)

			// act
			result, statusCode := s.mockApp.GetPostByID(ctx, states.Post2ID)

			// assert
			require.Equal(t, http.StatusInternalServerError, statusCode)

			require.Contains(t, string(result), fmt.Sprintf("getting post error: %v", assert.AnError))
		})

		t.Run("getting comments error", func(t *testing.T) {
			t.Parallel()

			// arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockRepo.EXPECT().GetPostByID(gomock.Any(), states.Post2ID).Return(fixtures.BuildPost().Valid().P(), nil)
			s.mockRepo.EXPECT().GetCommentsByPostID(gomock.Any(), states.Post2ID).Return(nil, assert.AnError)

			// act
			result, statusCode := s.mockApp.GetPostByID(ctx, states.Post2ID)

			// assert
			require.Equal(t, http.StatusInternalServerError, statusCode)
			require.Contains(t, string(result), fmt.Sprintf("getting comments error: %v", assert.AnError))
		})
	})
}
