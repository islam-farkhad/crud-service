package crud

import "homework-3/internal/pkg/repository"

type addPostRequest struct {
	Content string `json:"content"`
	Likes   int64  `json:"likes"`
}

type updatePostRequest struct {
	addPostRequest
	ID int64 `json:"id"`
}

type addCommentRequest struct {
	Content string `json:"content"`
}

// GetPostByIDResponse struct for HandleGetPostByID
type GetPostByIDResponse struct {
	Post     repository.Post      `json:"post"`
	Comments []repository.Comment `json:"comments"`
}
