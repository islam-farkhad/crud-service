package tests

import (
	"crud-service/internal/handlers"
	"crud-service/internal/pkg/repository"
	"crud-service/internal/utils"
	"crud-service/tests/states"
	"encoding/json"
	"fmt"
)

// AddPostToTestDB adds a valid post to test db
func AddPostToTestDB(testApp *handlers.App) *repository.Post {
	postRaw := []byte(fmt.Sprintf(`{"content":"%s","likes":%d}`, states.Post1Content, states.Post1Likes))
	req, rr := utils.GetRequestAndResponseRecorder(states.PostMethod, "/test/post", postRaw)

	testApp.Router.ServeHTTP(rr, req)

	post := &repository.Post{}
	err := json.Unmarshal(rr.Body.Bytes(), post)
	if err != nil {
		panic(err)
	}
	return post
}
