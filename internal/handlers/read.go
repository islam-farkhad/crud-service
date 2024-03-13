package handlers

import (
	"context"
	"crud-service/internal/pkg/repository"
	"crud-service/internal/utils"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

// GetPostByID handles the HTTP request for retrieving a post by its ID.
func (app *App) GetPostByID(ctx context.Context, ID int64) ([]byte, int) {

	postRepo, err := app.Repo.GetPostByID(ctx, ID)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			return []byte((fmt.Sprintf("postRepo with ID=%d not found, err: %v", ID, err))), http.StatusNotFound
		}
		return []byte(fmt.Sprintf("getting post error: %v", err)), http.StatusInternalServerError
	}

	comments, err := app.Repo.GetCommentsByPostID(ctx, ID)
	if err != nil {
		return []byte(fmt.Sprintf("getting comments error: %v", err)), http.StatusInternalServerError
	}

	response := GetPostByIDResponse{
		Post:     *postRepo,
		Comments: comments,
	}

	responseJSON, _ := json.Marshal(response)
	return responseJSON, http.StatusOK
}

// HandleGetPostByID processes an HTTP request to retrieve a post by its ID.
func (app *App) HandleGetPostByID(w http.ResponseWriter, req *http.Request) {
	id, status := utils.RetrieveID(req)
	if status != http.StatusOK {
		w.WriteHeader(status)
		return
	}
	data, status := app.GetPostByID(req.Context(), id)
	w.WriteHeader(status)
	_, err := w.Write(data)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
