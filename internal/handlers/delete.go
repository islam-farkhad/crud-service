package handlers

import (
	"context"
	"errors"
	"fmt"
	"homework-3/internal/pkg/repository"
	"homework-3/internal/utils"
	"log"
	"net/http"
)

// DeletePostByID deletes post and its comments with provided postID
func (app *App) DeletePostByID(ctx context.Context, postID int64) ([]byte, int) {

	_, err := app.Repo.DeletePostByID(ctx, postID)
	if err != nil {
		if errors.Is(err, repository.ErrObjectNotFound) {
			return []byte(fmt.Sprintf("postRepo with postID=%d not found, err: %v", postID, err)), http.StatusNotFound
		}
		return []byte(fmt.Sprintf("deleting post by postID error: %v", err)), http.StatusInternalServerError
	}
	return nil, http.StatusOK
}

// HandleDeletePostByID processes an HTTP request to retrieve a post by its ID.
func (app *App) HandleDeletePostByID(w http.ResponseWriter, req *http.Request) {

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
