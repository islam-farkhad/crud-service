package crud

import (
	"encoding/json"
	"fmt"
	"homework-3/internal/pkg/repository"
	"homework-3/internal/utils"
	"io"
	"net/http"
)

// CreateComment handles the HTTP request for creating a comment.
// It reads the comment data from the request body, validates the input,
// adds the comment to the repository, and responds with the created comment.
func (app *App) CreateComment(w http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("could not read body from request: %w", err))
		return
	}

	var unmarshal addCommentRequest
	if err = json.Unmarshal(body, &unmarshal); err != nil {
		utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("could not unmarshal from body. Body: %v, err: %w", body, err))
		return
	}

	postID, ok := utils.GetIDFromQueryParams(w, req)
	if !ok {
		return
	}

	comment := &repository.Comment{
		PostID:  postID,
		Content: unmarshal.Content,
	}

	commentID, err := app.Repo.AddComment(req.Context(), comment)
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("could not add comment. Body: %v, err: %w", body, err))
		return
	}

	comment.ID = commentID
	commentJSON, err := json.Marshal(comment)
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("can not marshal comment. comment.PostID: %d, comment.Content: %s, err: %w", comment.PostID, comment.Content, err))
		return
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(commentJSON)
	if err != nil {
		utils.HandleError(w, http.StatusInternalServerError, fmt.Errorf("write error: %w", err))
		return
	}
}
