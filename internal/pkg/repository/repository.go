//go:generate mockgen -source ./repository.go -destination=./mocks/repository.go -package=mock_repository
package repository

import "context"

type Repo interface {
	AddPost(ctx context.Context, post *Post) (int64, error)
	AddComment(ctx context.Context, comment *Comment) (int64, error)
	UpdatePost(ctx context.Context, post *Post) (int64, error)
	GetPostByID(ctx context.Context, id int64) (*Post, error)
	GetCommentsByPostID(ctx context.Context, postID int64) ([]Comment, error)
	DeletePostByID(ctx context.Context, id int64) (bool, error)
}
