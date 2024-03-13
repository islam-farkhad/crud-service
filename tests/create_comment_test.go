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

func getCommentRoute(postID int64) string {
	return fmt.Sprintf("/test/post/%d/comment", postID)
}

func Test_CreateComment(t *testing.T) {
	t.Parallel()
	var (
		commentRoute = ""
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()
		database.SetUp(t)
		defer database.TearDown()

		// arrange
		post := AddPostToTestDB(&testApp)

		commentRaw := []byte(fmt.Sprintf(`{"content":"%s"}`, states.Comment1Content))
		commentRoute = getCommentRoute(post.ID)
		req, rr := utils.GetRequestAndResponseRecorder(states.PostMethod, commentRoute, commentRaw)

		//act
		testApp.Router.ServeHTTP(rr, req)

		//assert
		comment := &repository.Comment{}
		err := json.Unmarshal(rr.Body.Bytes(), comment)
		if err != nil {
			panic(err)
		}

		require.Equal(t, http.StatusOK, rr.Code)
		require.Equal(t, comment.Content, states.Comment1Content)
		require.Equal(t, comment.PostID, post.ID)
	})

	t.Run("fail - no post with such id", func(t *testing.T) {
		t.Parallel()
		database.SetUp(t)
		defer database.TearDown()

		// arrange
		commentRaw := []byte(fmt.Sprintf(`{"content":"%s"}`, states.Comment1Content))
		commentRoute = getCommentRoute(states.Post1ID) // DB has no posts
		req, rr := utils.GetRequestAndResponseRecorder(states.PostMethod, commentRoute, commentRaw)

		//act
		testApp.Router.ServeHTTP(rr, req)

		//assert
		require.Equal(t, http.StatusBadRequest, rr.Code)
	})
}
