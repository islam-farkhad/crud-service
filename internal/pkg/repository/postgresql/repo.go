package postgresql

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"

	"crud-service/internal/pkg/db"
	"crud-service/internal/pkg/repository"
)

// Repo is a struct representing the PostgreSQL repository for handling Post and Comment entities.
type Repo struct {
	db db.DBops
}

// NewRepo creates a new instance of the PostgreSQL repository.
func NewRepo(database db.DBops) *Repo {
	return &Repo{db: database}
}

// AddPost adds a new Post to the database and returns its ID.
func (r *Repo) AddPost(ctx context.Context, post *repository.Post) (int64, error) {
	var id int64
	var createdAt time.Time

	err := r.db.ExecQueryRow(ctx, `INSERT INTO posts(content, likes) VALUES($1, $2) RETURNING id, created_at;`, post.Content, post.Likes).Scan(&id, &createdAt)

	post.CreatedAt = createdAt
	return id, err
}

// AddComment adds a new Comment to the database and returns its ID.
func (r *Repo) AddComment(ctx context.Context, comment *repository.Comment) (int64, error) {
	var id int64
	var createdAt time.Time

	err := r.db.ExecQueryRow(ctx, `INSERT INTO comments(post_id, content) VALUES($1, $2) RETURNING id, created_at;`, comment.PostID, comment.Content).Scan(&id, &createdAt)

	comment.CreatedAt = createdAt
	return id, err
}

// UpdatePost updates an existing Post in the database.
func (r *Repo) UpdatePost(ctx context.Context, post *repository.Post) (int64, error) {
	var id int64
	var createdAt time.Time

	err := r.db.ExecQueryRow(ctx, `UPDATE posts SET content=$1, likes=$2 WHERE id=$3 RETURNING id, created_at;`, post.Content, post.Likes, post.ID).Scan(&id, &createdAt)
	if errors.Is(err, pgx.ErrNoRows) {
		return post.ID, repository.ErrObjectNotFound
	}

	post.CreatedAt = createdAt
	return id, err
}

// GetPostByID retrieves a Post from the database based on its ID.
func (r *Repo) GetPostByID(ctx context.Context, id int64) (*repository.Post, error) {
	var post repository.Post
	err := r.db.Get(ctx, &post, "SELECT id, content, likes, created_at FROM posts WHERE id=$1;", id)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, repository.ErrObjectNotFound
		}
		return nil, err
	}
	return &post, nil
}

// GetCommentsByPostID retrieves Comments associated with a specific Post from the database.
func (r *Repo) GetCommentsByPostID(ctx context.Context, postID int64) ([]repository.Comment, error) {
	var comments []repository.Comment
	err := r.db.Select(ctx, &comments, "SELECT id, content, created_at FROM comments WHERE post_id=$1;", postID)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

// DeletePostByID deletes a Post from the database based on its ID.
func (r *Repo) DeletePostByID(ctx context.Context, id int64) (bool, error) {
	res, err := r.db.Exec(ctx, "DELETE FROM posts WHERE id=$1", id)
	if err != nil {
		return false, fmt.Errorf("error when executed delete: %v", err)
	}

	if res.RowsAffected() == 0 {
		return false, repository.ErrObjectNotFound
	}
	return true, nil
}
