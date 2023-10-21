package repository

import "time"

// Post represents a post from database in the system.
type Post struct {
	ID        int64     `db:"id"`
	Content   string    `db:"content"`
	Likes     int64     `db:"likes"`
	CreatedAt time.Time `db:"created_at"`
}

// Comment represents a comment from database associated with a post.
type Comment struct {
	ID        int64     `db:"id"`
	PostID    int64     `db:"post_id"`
	Content   string    `db:"content"`
	CreatedAt time.Time `db:"created_at"`
}
