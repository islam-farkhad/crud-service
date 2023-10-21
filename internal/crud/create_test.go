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

func getCreatePostRequestAndResponseRecorder(body []byte) (*http.Request, *httptest.ResponseRecorder) {
	req, _ := http.NewRequest("POST", "/post", strings.NewReader(string(body)))
	return req, httptest.NewRecorder()
}

func TestCreatePost(t *testing.T) {
	t.Parallel()
	var (
		postRequest = addPostRequest{
			Content: "Test Content",
			Likes:   10,
		}
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		// arrange
		s := setUp(t)
		defer s.tearDown()

		s.mockRepo.EXPECT().AddPost(gomock.Any(), gomock.Any()).Return(int64(1), nil)

		reqBody, _ := json.Marshal(postRequest)
		req, rr := getCreatePostRequestAndResponseRecorder(reqBody)

		// act
		s.mockApp.CreatePost(rr, req)

		// assert
		require.Equal(t, http.StatusOK, rr.Code)

		var postResponse repository.Post
		err := json.Unmarshal(rr.Body.Bytes(), &postResponse)
		require.NoError(t, err)

		assert.Equal(t, int64(1), postResponse.ID)
		assert.Equal(t, postRequest.Content, postResponse.Content)
		assert.Equal(t, postRequest.Likes, postResponse.Likes)
	})

	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		t.Run("bad input", func(t *testing.T) {
			t.Parallel()

			s := setUp(t)
			defer s.tearDown()
			// arrange
			invalidRequest := []byte(`{"bad_field": "bad"}`)
			req, rr := getCreatePostRequestAndResponseRecorder(invalidRequest)

			// act
			s.mockApp.CreatePost(rr, req)

			// assert
			require.Equal(t, http.StatusBadRequest, rr.Code)
			require.Contains(t, rr.Body.String(), fmt.Sprintf(`"content" field is missing`))
		})

		t.Run("repository error", func(t *testing.T) {
			t.Parallel()

			// arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockRepo.EXPECT().AddPost(gomock.Any(), gomock.Any()).Return(int64(0), assert.AnError)

			reqBody, _ := json.Marshal(postRequest)
			req, rr := getCreatePostRequestAndResponseRecorder(reqBody)

			// act
			s.mockApp.CreatePost(rr, req)

			// assert
			require.Equal(t, http.StatusInternalServerError, rr.Code)
			require.Contains(t, rr.Body.String(), fmt.Sprintf("could not add post. Body: %v, err: %v", reqBody, assert.AnError))
		})
	})
}
