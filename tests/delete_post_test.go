//go:build integration

package tests

import (
	"encoding/json"
	"fmt"
	"homework-3/internal/pkg/repository"
	"homework-3/internal/utils"
	"homework-3/tests/states"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func getDeletePostRoute(postID int64) string {
	return fmt.Sprintf("/test/post/%d", postID)
}

func Test_DeletePost(t *testing.T) {
	t.Parallel()
	var (
		createPostRoute = "/test/post"
		deleteRoute     string
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		database.SetUp(t)
		defer database.TearDown()

		// arrange
		postRaw := []byte(fmt.Sprintf(`{"content":"%s","likes":%d}`, states.Post1Content, states.Post1Likes))
		req, rr := utils.GetRequestAndResponseRecorder(states.PostMethod, createPostRoute, postRaw)

		testApp.Router.ServeHTTP(rr, req)

		post := &repository.Post{}
		err := json.Unmarshal(rr.Body.Bytes(), post)
		if err != nil {
			panic(err)
		}

		deleteRoute = getDeletePostRoute(post.ID)
		req, rr = utils.GetRequestAndResponseRecorder(states.DeleteMethod, deleteRoute, nil)

		//act
		testApp.Router.ServeHTTP(rr, req)

		//assert
		require.Equal(t, http.StatusOK, rr.Code)
	})

	t.Run("fail - no post with such id", func(t *testing.T) {
		t.Parallel()
		database.SetUp(t)
		defer database.TearDown()

		// arrange
		deleteRoute = getDeletePostRoute(states.Post1ID)
		req, rr := utils.GetRequestAndResponseRecorder(states.DeleteMethod, deleteRoute, nil)

		//act
		testApp.Router.ServeHTTP(rr, req)

		//assert
		require.Equal(t, http.StatusNotFound, rr.Code)
	})
}
