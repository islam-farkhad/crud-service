package tests

import (
	"encoding/json"
	"fmt"
	"homework-3/internal/crud"
	"homework-3/internal/pkg/repository"
	"homework-3/internal/utils"
	"homework-3/tests/states"
)

// AddPostToTestDB adds a valid post to test db
func AddPostToTestDB(testApp *crud.App) *repository.Post {
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
