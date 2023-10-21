package crud

import (
	"encoding/json"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"homework-3/internal/pkg/repository"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getReadRequestAndResponseRecorder(id int64) (*http.Request, *httptest.ResponseRecorder) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("/post/%d", id), nil)
	return req, httptest.NewRecorder()
}

func TestGetPostByID(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		// arrange
		s := setUp(t)
		defer s.tearDown()

		postID := int64(1)
		req, rr := getReadRequestAndResponseRecorder(postID)

		s.mockRepo.EXPECT().GetPostByID(gomock.Any(), postID).Return(&repository.Post{
			ID:      postID,
			Content: "Content",
			Likes:   5,
		}, nil)

		s.mockRepo.EXPECT().GetCommentsByPostID(gomock.Any(), postID).Return(
			[]repository.Comment{
				{ID: 1, PostID: postID, Content: "awesome"},
				{ID: 2, PostID: postID, Content: "hater is here!"},
			},
			nil)

		// act
		s.mockApp.Router.ServeHTTP(rr, req)

		// assert
		require.Equal(t, http.StatusOK, rr.Code)

		var response getPostByIDResponse
		err := json.Unmarshal(rr.Body.Bytes(), &response)
		require.NoError(t, err)

		assert.Equal(t, postID, response.Post.ID)
		assert.Equal(t, "Content", response.Post.Content)
		assert.Equal(t, int64(5), response.Post.Likes)
		assert.Len(t, response.Comments, 2)

		assert.Equal(t, int64(1), response.Comments[0].ID)
		assert.Equal(t, int64(2), response.Comments[1].ID)

		assert.Equal(t, postID, response.Comments[0].PostID)
		assert.Equal(t, postID, response.Comments[1].PostID)

		assert.Equal(t, "awesome", response.Comments[0].Content)
		assert.Equal(t, "hater is here!", response.Comments[1].Content)
	})

	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		t.Run("post not found", func(t *testing.T) {
			t.Parallel()

			// arrange
			s := setUp(t)
			defer s.tearDown()

			postID := int64(1)
			req, rr := getReadRequestAndResponseRecorder(postID)

			s.mockRepo.EXPECT().GetPostByID(gomock.Any(), postID).Return(nil, repository.ErrObjectNotFound)

			// act
			s.mockApp.Router.ServeHTTP(rr, req)

			// assert
			require.Equal(t, http.StatusNotFound, rr.Code)

			require.Contains(t, rr.Body.String(), fmt.Sprintf("postRepo with id=%d not found, err: %v", postID, repository.ErrObjectNotFound))
		})

		t.Run("internal error", func(t *testing.T) {
			t.Parallel()

			// arrange
			s := setUp(t)
			defer s.tearDown()

			postID := int64(1)
			req, rr := getReadRequestAndResponseRecorder(postID)

			s.mockRepo.EXPECT().GetPostByID(gomock.Any(), postID).Return(nil, assert.AnError)

			// act
			s.mockApp.Router.ServeHTTP(rr, req)

			// assert
			require.Equal(t, http.StatusInternalServerError, rr.Code)

			require.Contains(t, rr.Body.String(), fmt.Sprintf("getting post error: %v", assert.AnError))
		})

		t.Run("getting comments error", func(t *testing.T) {
			t.Parallel()

			// arrange
			s := setUp(t)
			defer s.tearDown()

			postID := int64(1)
			req, rr := getReadRequestAndResponseRecorder(postID)

			s.mockRepo.EXPECT().GetPostByID(gomock.Any(), postID).Return(&repository.Post{
				ID:      postID,
				Content: "Content",
				Likes:   5,
			}, nil)

			s.mockRepo.EXPECT().GetCommentsByPostID(gomock.Any(), postID).Return(nil, assert.AnError)

			// act
			s.mockApp.Router.ServeHTTP(rr, req)

			// assert
			require.Equal(t, http.StatusInternalServerError, rr.Code)
			require.Contains(t, rr.Body.String(), fmt.Sprintf("getting comments error: %v", assert.AnError))
		})
	})
}
