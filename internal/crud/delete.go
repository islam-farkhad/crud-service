package crud

import (
	"errors"
	"fmt"
	"homework-3/internal/pkg/repository"
	"homework-3/internal/utils"
	"net/http"
)

// DeletePostByID handles the HTTP request for deleting a post by its ID.
// It reads the post ID from the request parameters, attempts to delete the post,
// and responds with an appropriate status code.
func (app *App) DeletePostByID(w http.ResponseWriter, req *http.Request) {
	id, ok := utils.GetIDFromQueryParams(w, req)
	if !ok {
		return
	}

	_, err := app.Repo.DeletePostByID(req.Context(), id)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			utils.HandleError(w, http.StatusNotFound, fmt.Errorf("postRepo with id=%d not found, err: %w", id, err))
			return
		}
		utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("deleting post by id error: %w", err))
		return
	}
	w.WriteHeader(http.StatusOK)
}
