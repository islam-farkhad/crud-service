//go:build integration

package tests

import (
	"encoding/json"
	"fmt"
	"homework-3/internal/pkg/repository"
	"homework-3/internal/utils"
	"homework-3/tests/app"
	"homework-3/tests/states"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_CreatePost(t *testing.T) {
	t.Parallel()
	var (
		method  = "POST"
		route   = "/test/post"
		testApp = app.NewTestApp(database.DB)
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		database.SetUp(t)
		defer database.TearDown()

		// arrange
		postRaw := []byte(fmt.Sprintf(`{"content":"%s","likes":%d}`, states.Post1Content, states.Post1Likes))

		req, rr := utils.GetRequestAndResponseRecorder(method, route, postRaw)

		//act
		testApp.Router.ServeHTTP(rr, req)

		//assert
		post := &repository.Post{}
		fmt.Println(string(rr.Body.Bytes()))
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
		req, rr := utils.GetRequestAndResponseRecorder(method, route, postRaw)

		//act
		testApp.Router.ServeHTTP(rr, req)

		//assert
		require.Equal(t, http.StatusBadRequest, rr.Code)
	})
}
