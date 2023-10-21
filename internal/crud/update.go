package crud

import (
	"encoding/json"
	"errors"
	"fmt"
	"homework-3/internal/pkg/repository"
	"homework-3/internal/utils"
	"io"
	"net/http"
)

// UpdatePost handles the HTTP request for updating a post. It reads the request body,
// unmarshals the updatePostRequest, and updates the corresponding post in the repository.
// It returns the updated post in the response.
func (app *App) UpdatePost(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("can not read from request. req.Body: %v, err: %w", req.Body, err))
		return
	}

	var unmarshal updatePostRequest
	if err = json.Unmarshal(body, &unmarshal); err != nil {
		utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("can not unmarshal from body. body: %v, err: %w", body, err))
		return
	}

	if unmarshal.ID == 0 {
		utils.HandleError(w, http.StatusBadRequest, fmt.Errorf("provide a post id"))
		return
	}

	postRepo := &repository.Post{
		ID:      unmarshal.ID,
		Content: unmarshal.Content,
		Likes:   unmarshal.Likes,
	}
	id, err := app.Repo.UpdatePost(req.Context(), postRepo)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			utils.HandleError(w, http.StatusNotFound, fmt.Errorf("postRepo with id=%d not found, err: %w", id, err))
			return
		}
		utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("updating post error: %w", err))
		return
	}
	postJSON, err := json.Marshal(postRepo)
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("can not marshal Post. postRepo.Content: %s, postRepo.Likes: %d, err: %w", postRepo.Content, postRepo.Likes, err))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(postJSON)
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("writing err: %w", err))
		return
	}
}
