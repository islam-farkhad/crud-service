//go:build integration

package tests

import (
	"crud-service/internal/pkg/repository"
	"crud-service/internal/utils"
	"crud-service/tests/states"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_UpdatePost(t *testing.T) {
	t.Parallel()
	var (
		route = "/test/post"
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		database.SetUp(t)
		defer database.TearDown()

		// arrange
		post := AddPostToTestDB(&testApp)

		body := []byte(fmt.Sprintf(`{"content":"%s","likes":%d,"id":%d}`, states.Post2Content, states.Post2Likes, post.ID))
		req, rr := utils.GetRequestAndResponseRecorder(states.PutMethod, route, body)
		//act
		testApp.Router.ServeHTTP(rr, req)

		//assert
		post = &repository.Post{}
		err := json.Unmarshal(rr.Body.Bytes(), post)
		if err != nil {
			panic(err)
		}

		require.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, post.Content, states.Post2Content)
		require.Equal(t, post.Likes, states.Post2Likes)
	})

	t.Run("fail - post not found", func(t *testing.T) {
		t.Parallel()
		database.SetUp(t)
		defer database.TearDown()

		// arrange
		postRaw := []byte(fmt.Sprintf(`{"content":"%s","likes":%d,"id":1}`, states.Post1Content, states.Post1Likes))
		req, rr := utils.GetRequestAndResponseRecorder(states.PutMethod, route, postRaw)

		//act
		testApp.Router.ServeHTTP(rr, req)

		//assert
		require.Equal(t, http.StatusNotFound, rr.Code)
	})
}
