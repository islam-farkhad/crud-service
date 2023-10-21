package crud

import (
	"encoding/json"
	"errors"
	"fmt"
	"homework-3/internal/pkg/repository"
	"homework-3/internal/utils"
	"net/http"
)

// GetPostByID handles the HTTP request for retrieving a post by its ID.
// It reads the post ID from query parameters, fetches the post and its comments
// from the repository, and responds with a JSON with the post and comments.
func (app *App) GetPostByID(w http.ResponseWriter, req *http.Request) {
	id, ok := utils.GetIDFromQueryParams(w, req)
	if !ok {
		return
	}

	postRepo, err := app.Repo.GetPostByID(req.Context(), id)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			utils.HandleError(w, http.StatusNotFound, fmt.Errorf("postRepo with id=%d not found, err: %w", id, err))
			return
		}
		utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("getting post error: %w", err))
		return
	}

	comments, err := app.Repo.GetCommentsByPostID(req.Context(), id)
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("getting comments error: %w", err))
		return
	}

	response := getPostByIDResponse{
		Post:     *postRepo,
		Comments: comments,
	}

	responseJSON, err := json.Marshal(response)
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("can not marshal response: %v, err: %w", response, err))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(responseJSON)
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("writing error: %w", err))
		return
	}
}
