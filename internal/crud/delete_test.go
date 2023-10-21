package crud

import (
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"homework-3/internal/pkg/repository"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getDeleteRequestAndResponseRecorder(postID any) (*http.Request, *httptest.ResponseRecorder) {
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/post/%v", postID), nil)
	return req, httptest.NewRecorder()
}

func TestDeletePostByID(t *testing.T) {
	t.Parallel()

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		// arrange
		s := setUp(t)
		defer s.tearDown()

		postID := int64(1)
		req, rr := getDeleteRequestAndResponseRecorder(postID)

		s.mockRepo.EXPECT().DeletePostByID(gomock.Any(), gomock.Any()).Return(true, nil)

		// act
		s.mockApp.Router.ServeHTTP(rr, req)

		// assert
		require.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		t.Run("invalid id", func(t *testing.T) {
			t.Parallel()

			// arrange
			s := setUp(t)
			defer s.tearDown()

			badId := "bad"
			req, rr := getDeleteRequestAndResponseRecorder(badId)

			// act
			s.mockApp.Router.ServeHTTP(rr, req)

			// assert
			require.Equal(t, http.StatusBadRequest, rr.Code)
			require.Contains(t, rr.Body.String(), "id should be a number")
		})

		t.Run("not existing id", func(t *testing.T) {
			t.Parallel()

			// arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockRepo.EXPECT().DeletePostByID(gomock.Any(), gomock.Any()).Return(false, repository.ErrObjectNotFound)

			notExistingID := 0
			req, rr := getDeleteRequestAndResponseRecorder(notExistingID)

			// act
			s.mockApp.Router.ServeHTTP(rr, req)

			// assert
			require.Equal(t, http.StatusNotFound, rr.Code)
			require.Contains(t, rr.Body.String(), repository.ErrObjectNotFound.Error())
		})

		t.Run("deleting error", func(t *testing.T) {
			t.Parallel()

			// arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockRepo.EXPECT().DeletePostByID(gomock.Any(), gomock.Any()).Return(false, assert.AnError)

			postID := int64(1)
			req, rr := getDeleteRequestAndResponseRecorder(postID)

			// act
			s.mockApp.Router.ServeHTTP(rr, req)

			// assert
			require.Equal(t, http.StatusInternalServerError, rr.Code)
			require.Contains(t, rr.Body.String(), fmt.Sprintf("deleting post by id error: %v", assert.AnError))
		})

	})

}
