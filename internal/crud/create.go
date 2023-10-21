package crud

import (
	"encoding/json"
	"fmt"
	"homework-3/internal/pkg/repository"
	"homework-3/internal/utils"
	"io"
	"net/http"
)

// CreatePost handles the HTTP request for creating a post.
// It reads the post data from the request body, validates the input,
// adds the post to the repository, and responds with the created post.
func (app *App) CreatePost(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("could not read body from request: %w", err))
		return
	}

	var unmarshal addPostRequest
	if err = json.Unmarshal(body, &unmarshal); err != nil {
		utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("could not unmarshal from body. Body: %v, err: %w", body, err))
		return
	}

	postRepo := &repository.Post{
		Content: unmarshal.Content,
		Likes:   unmarshal.Likes,
	}
	id, err := app.Repo.AddPost(req.Context(), postRepo)
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("could not add post. Body: %v, err: %w", body, err))
		return
	}
	postRepo.ID = id
	postJSON, err := json.Marshal(postRepo)
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("can not marshal Post. postRepo.Content: %s, postRepo.Likes: %d, err: %w", postRepo.Content, postRepo.Likes, err))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(postJSON)
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("write error: %w", err))
		return
	}
}
