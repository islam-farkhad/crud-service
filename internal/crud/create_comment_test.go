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
	"strings"
	"testing"
)

func getCreateCommentRequestAndResponseRecorder(body []byte, postID int64) (*http.Request, *httptest.ResponseRecorder) {
	req, _ := http.NewRequest("POST", fmt.Sprintf("/post/%d/comment", postID), strings.NewReader(string(body)))
	return req, httptest.NewRecorder()
}

func TestCreateComment(t *testing.T) {
	t.Parallel()
	var (
		commentRequest = addCommentRequest{
			Content: "Test Comment",
		}
		postID = int64(1)
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		// arrange
		s := setUp(t)
		defer s.tearDown()

		s.mockRepo.EXPECT().AddComment(gomock.Any(), gomock.Any()).Return(int64(1), nil)

		reqBody, _ := json.Marshal(commentRequest)
		req, rr := getCreateCommentRequestAndResponseRecorder(reqBody, postID)

		// act
		s.mockApp.Router.ServeHTTP(rr, req)

		// assert
		require.Equal(t, http.StatusOK, rr.Code)

		var commentResponse repository.Comment
		err := json.Unmarshal(rr.Body.Bytes(), &commentResponse)
		require.NoError(t, err)

		require.Equal(t, http.StatusOK, rr.Code)

		assert.Equal(t, int64(1), commentResponse.ID)
		assert.Equal(t, postID, commentResponse.PostID)
		assert.Equal(t, commentRequest.Content, commentResponse.Content)
	})

	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		t.Run("bad input", func(t *testing.T) {
			t.Parallel()

			// arrange
			s := setUp(t)
			defer s.tearDown()

			badRequest := []byte(`{"bad_field": "bad"}`)
			req, rr := getCreateCommentRequestAndResponseRecorder(badRequest, postID)

			// act
			s.mockApp.Router.ServeHTTP(rr, req)

			// assert
			require.Equal(t, http.StatusBadRequest, rr.Code)
			require.Contains(t, rr.Body.String(), fmt.Sprintf(`"content" field is missing`))
		})

		t.Run("repository error", func(t *testing.T) {
			t.Parallel()

			// arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockRepo.EXPECT().AddComment(gomock.Any(), gomock.Any()).Return(int64(0), assert.AnError)

			reqBody, _ := json.Marshal(commentRequest)
			req, _ := http.NewRequest("POST", fmt.Sprintf("/post/%d/comment", postID), strings.NewReader(string(reqBody)))
			rr := httptest.NewRecorder()

			// act
			s.mockApp.Router.ServeHTTP(rr, req)

			// assert
			require.Equal(t, http.StatusInternalServerError, rr.Code)
			require.Contains(t, rr.Body.String(), fmt.Sprintf("could not add comment. Body: %v, err: %v", reqBody, assert.AnError))
		})
	})
}
