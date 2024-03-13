package fixtures

import (
	"crud-service/internal/pkg/repository"
	"crud-service/tests/states"
	"time"
)

// CommentBuilder is a builder pattern for constructing instances of `repository.Comment`.
type CommentBuilder struct {
	instance *repository.Comment
}

// BuildComment creates a new `CommentBuilder` instance with an empty `repository.Comment`.
func BuildComment() *CommentBuilder {
	return &CommentBuilder{instance: &repository.Comment{}}
}

// ID sets the `ID` field of the comment being built.
func (b *CommentBuilder) ID(v int64) *CommentBuilder {
	b.instance.ID = v
	return b
}

// PostID sets the `PostID` field of the comment being built.
func (b *CommentBuilder) PostID(v int64) *CommentBuilder {
	b.instance.PostID = v
	return b
}

// Content sets the `Content` field of the comment being built.
func (b *CommentBuilder) Content(v string) *CommentBuilder {
	b.instance.Content = v
	return b
}

// CreatedAt sets the `CreatedAt` field of the comment being built.
func (b *CommentBuilder) CreatedAt(v time.Time) *CommentBuilder {
	b.instance.CreatedAt = v
	return b
}

// P returns the built `repository.Comment`.
func (b *CommentBuilder) P() *repository.Comment {
	return b.instance
}

// V returns the built `repository.Comment`.
func (b *CommentBuilder) V() repository.Comment {
	return *b.instance
}

// Valid creates a new `CommentBuilder` instance with predefined values for testing.
func (b *CommentBuilder) Valid() *CommentBuilder {
	return BuildComment().ID(states.Comment1ID).PostID(states.Post1ID).Content(states.Comment1Content).CreatedAt(time.Time{})
}
