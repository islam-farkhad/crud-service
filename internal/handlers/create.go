package handlers

import (
	"context"
	"crud-service/internal/pkg/repository"
	"crud-service/internal/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// CreatePost creates a new post in database.
func (app *App) CreatePost(ctx context.Context, postRepo *repository.Post) ([]byte, int) {

	id, err := app.Repo.AddPost(ctx, postRepo)
	if err != nil {
		return []byte(fmt.Sprintf("could not add post. err: %v", err)), http.StatusInternalServerError
	}
	postRepo.ID = id

	postJSON, err := json.Marshal(postRepo)
	if err != nil {
		return []byte(fmt.Sprintf("could not marshal postRepo. err: %v", err)), http.StatusInternalServerError
	}

	return postJSON, http.StatusOK
}

func parseCreatePost(body []byte) (*repository.Post, int) {
	var unmarshal addPostRequest
	if err := json.Unmarshal(body, &unmarshal); err != nil {
		return nil, http.StatusBadRequest
	}

	if len(unmarshal.Content) == 0 {
		return nil, http.StatusBadRequest
	}

	postRepo := &repository.Post{
		Content: unmarshal.Content,
		Likes:   unmarshal.Likes,
	}

	return postRepo, http.StatusOK
}

// HandleCreatePost processes an HTTP request for creating a new post.
func (app *App) HandleCreatePost(w http.ResponseWriter, req *http.Request) {
	body, status := utils.RetrieveBody(req)
	if status != http.StatusOK {
		w.WriteHeader(status)
		return
	}

	postRepo, status := parseCreatePost(body)
	if status != http.StatusOK {
		w.WriteHeader(status)
		return
	}
	data, status := app.CreatePost(req.Context(), postRepo)
	w.WriteHeader(status)
	_, err := w.Write(data)
	if err != nil {
		log.Printf("%+v", err)
		return
	}
}
