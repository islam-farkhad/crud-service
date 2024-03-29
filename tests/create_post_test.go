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

func Test_CreatePost(t *testing.T) {
	t.Parallel()
	var (
		route = "/test/post"
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		database.SetUp(t)
		defer database.TearDown()

		// arrange
		postRaw := []byte(fmt.Sprintf(`{"content":"%s","likes":%d}`, states.Post1Content, states.Post1Likes))

		req, rr := utils.GetRequestAndResponseRecorder(states.PostMethod, route, postRaw)

		//act
		testApp.Router.ServeHTTP(rr, req)

		//assert
		post := &repository.Post{}
		err := json.Unmarshal(rr.Body.Bytes(), post)
		if err != nil {
			panic(err)
		}

		require.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, post.Content, states.Post1Content)
		require.Equal(t, post.Likes, states.Post1Likes)
	})

	t.Run("fail - empty content", func(t *testing.T) {
		t.Parallel()
		database.SetUp(t)
		defer database.TearDown()

		// arrange
		postRaw := []byte(fmt.Sprintf(`{"no_content":"%s","likes":%d}`, states.Post1Content, states.Post1Likes))
		req, rr := utils.GetRequestAndResponseRecorder(states.PostMethod, route, postRaw)

		//act
		testApp.Router.ServeHTTP(rr, req)

		//assert
		require.Equal(t, http.StatusBadRequest, rr.Code)
	})
}
