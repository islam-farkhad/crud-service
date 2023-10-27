package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"homework-3/internal/pkg/repository"
	"homework-3/internal/utils"
	"log"
	"net/http"
)

// CreateComment adds the comment to the repository.
func (app *App) CreateComment(ctx context.Context, comment *repository.Comment) ([]byte, int) {

	commentID, err := app.Repo.AddComment(ctx, comment)
	if err != nil {
		return []byte(fmt.Sprintf("could not add comment. err: %v", err)), http.StatusBadRequest
	}

	comment.ID = commentID
	commentJSON, err := json.Marshal(comment)
	if err != nil {
		return []byte(fmt.Sprintf("can not marshal comment. err: %v", err)), http.StatusInternalServerError
	}

	return commentJSON, http.StatusOK
}

func parseCreateComment(body []byte, postID int64) (*repository.Comment, int) {
	var unmarshal addCommentRequest
	if err := json.Unmarshal(body, &unmarshal); err != nil {
		return nil, http.StatusBadRequest
	}

	if len(unmarshal.Content) == 0 {
		return nil, http.StatusBadRequest
	}

	comment := &repository.Comment{
		PostID:  postID,
		Content: unmarshal.Content,
	}

	return comment, http.StatusOK
}

// HandleCreateComment processes an HTTP request to create a comment for a post.
func (app *App) HandleCreateComment(w http.ResponseWriter, req *http.Request) {
	body, status := utils.RetrieveBody(req)
	if status != http.StatusOK {
		w.WriteHeader(status)
		return
	}

	postID, status := utils.RetrieveID(req)
	if status != http.StatusOK {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	commentRepo, status := parseCreateComment(body, postID)
	if status != http.StatusOK {
		w.WriteHeader(status)
		return
	}
	data, status := app.CreateComment(req.Context(), commentRepo)
	w.WriteHeader(status)
	_, err := w.Write(data)
	if err != nil {
		log.Println(err.Error())
		return
	}
}
