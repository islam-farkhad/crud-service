package crud

import (
	"errors"
	"fmt"
	"homework-3/internal/pkg/repository"
	"homework-3/internal/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// DeletePostByID handles the HTTP request for deleting a post by its ID.
// It reads the post ID from the request parameters, attempts to delete the post,
// and responds with an appropriate status code.
func (app *App) DeletePostByID(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	idStr := vars["id"]
	if idStr == "" {
		utils.HandleError(w, http.StatusBadRequest, fmt.Errorf("provide id in query parameters"))
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		utils.HandleError(w, http.StatusBadRequest, fmt.Errorf("id should be a number"))
		return
	}

	_, err = app.Repo.DeletePostByID(req.Context(), id)
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
