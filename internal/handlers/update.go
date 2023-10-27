package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"homework-3/internal/pkg/repository"
	"homework-3/internal/utils"
	"log"
	"net/http"
)

// UpdatePost handles the HTTP request for updating a post. It reads the request body,
// unmarshals the updatePostRequest, and updates the corresponding post in the repository.
// It returns the updated post in the response.
func (app *App) UpdatePost(ctx context.Context, postRepo *repository.Post) ([]byte, int) {

	_, err := app.Repo.UpdatePost(ctx, postRepo)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			return []byte(fmt.Sprintf("postRepo with ID=%d not found, err: %v", postRepo.ID, err)), http.StatusNotFound
		}
		return []byte(fmt.Sprintf("updating post error: %v", err)), http.StatusInternalServerError
	}

	postJSON, err := json.Marshal(postRepo)
	if err != nil {
		return []byte(fmt.Sprintf("can not marshal postRepo.  err: %v", err)), http.StatusInternalServerError

	}

	return postJSON, http.StatusOK
}

func parseUpdatePost(body []byte) (*repository.Post, int) {
	var unmarshal updatePostRequest
	if err := json.Unmarshal(body, &unmarshal); err != nil {
		return nil, http.StatusInternalServerError
	}

	if unmarshal.ID == 0 {
		return nil, http.StatusBadRequest
	}

	postRepo := &repository.Post{
		ID:      unmarshal.ID,
		Content: unmarshal.Content,
		Likes:   unmarshal.Likes,
	}

	return postRepo, http.StatusOK
}

// HandleUpdatePost function handles the HTTP request for updating a post.
func (app *App) HandleUpdatePost(w http.ResponseWriter, req *http.Request) {
	body, status := utils.RetrieveBody(req)
	if status != http.StatusOK {
		w.WriteHeader(status)
		return
	}

	postRepo, status := parseUpdatePost(body)
	if status != http.StatusOK {
		w.WriteHeader(status)
		return
	}
	data, status := app.UpdatePost(req.Context(), postRepo)
	w.WriteHeader(status)
	_, err := w.Write(data)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
