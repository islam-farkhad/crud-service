//go:build integration

package tests

import (
	"encoding/json"
	"fmt"
	"homework-3/internal/crud"
	"homework-3/internal/utils"
	"homework-3/tests/states"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func getReadPostRoute(postID int64) string {
	return fmt.Sprintf("/test/post/%d", postID)
}

func Test_ReadPost(t *testing.T) {
	t.Parallel()
	var (
		readRoute string
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		database.SetUp(t)
		defer database.TearDown()

		// arrange
		post := AddPostToTestDB(&testApp)

		readRoute = getReadPostRoute(post.ID)
		req, rr := utils.GetRequestAndResponseRecorder(states.GetMethod, readRoute, nil)

		//act
		testApp.Router.ServeHTTP(rr, req)

		response := &crud.GetPostByIDResponse{}
		err := json.Unmarshal(rr.Body.Bytes(), response)
		if err != nil {
			panic(err)
		}

		//assert
		require.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, *post, response.Post)
	})

	t.Run("fail - no post with such id", func(t *testing.T) {
		t.Parallel()
		database.SetUp(t)
		defer database.TearDown()

		// arrange
		readRoute = getReadPostRoute(states.Post1ID)
		req, rr := utils.GetRequestAndResponseRecorder(states.GetMethod, readRoute, nil)

		//act
		testApp.Router.ServeHTTP(rr, req)

		//assert
		require.Equal(t, http.StatusNotFound, rr.Code)
	})
}
