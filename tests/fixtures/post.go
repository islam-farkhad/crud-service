package fixtures

import (
	"homework-3/internal/pkg/repository"
	"homework-3/tests/states"
	"time"
)

// PostBuilder is a builder pattern for creating instances of `repository.Post`.
type PostBuilder struct {
	instance *repository.Post
}

// BuildPost creates a new `PostBuilder` instance with an empty `repository.Post`.
func BuildPost() *PostBuilder {
	return &PostBuilder{instance: &repository.Post{}}
}

// ID sets the `ID` field of the post being built.
func (b *PostBuilder) ID(v int64) *PostBuilder {
	b.instance.ID = v
	return b
}

// Content sets the `Content` field of the post being built.
func (b *PostBuilder) Content(v string) *PostBuilder {
	b.instance.Content = v
	return b
}

// Likes sets the `Likes` field of the post being built.
func (b *PostBuilder) Likes(v int64) *PostBuilder {
	b.instance.Likes = v
	return b
}

// CreatedAt sets the `CreatedAt` field of the post being built.
func (b *PostBuilder) CreatedAt(v time.Time) *PostBuilder {
	b.instance.CreatedAt = v
	return b
}

// P returns the built `repository.Post`.
func (b *PostBuilder) P() *repository.Post {
	return b.instance
}

// V returns the built `repository.Post`.
func (b *PostBuilder) V() repository.Post {
	return *b.instance
}

// Valid creates a new `PostBuilder` instance with predefined values for testing.
func (b *PostBuilder) Valid() *PostBuilder {
	return BuildPost().ID(states.Post1ID).Content(states.Post1Content).Likes(states.Post1Likes).CreatedAt(states.Post1CreatedAt)
}
