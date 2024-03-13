package postgresql

import (
	"context"
	"crud-service/internal/pkg/repository"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRepo_GetPostByID(t *testing.T) {
	t.Parallel()
	var (
		ctx            = context.Background()
		id       int64 = 1
		queryRow       = "SELECT id, content, likes, created_at FROM posts WHERE id=$1;"
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		// arrange
		s := setUp(t)
		defer s.tearDown()

		s.mockDb.EXPECT().Get(gomock.Any(), gomock.Any(), queryRow, gomock.Any()).Return(nil)

		// act
		post, err := s.repo.GetPostByID(ctx, id)

		// assert
		require.NoError(t, err)
		assert.Equal(t, int64(0), post.ID)

	})

	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		t.Run("not found", func(t *testing.T) {
			t.Parallel()
			// arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockDb.EXPECT().Get(gomock.Any(), gomock.Any(), queryRow, gomock.Any()).Return(pgx.ErrNoRows)

			// act
			post, err := s.repo.GetPostByID(ctx, id)

			// assert
			require.EqualError(t, err, repository.ErrObjectNotFound.Error())
			require.Nil(t, post)
		})

		t.Run("internal error", func(t *testing.T) {
			t.Parallel()
			// arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockDb.EXPECT().Get(gomock.Any(), gomock.Any(), queryRow, gomock.Any()).Return(assert.AnError)

			// act
			post, err := s.repo.GetPostByID(ctx, id)

			// assert
			require.EqualError(t, err, assert.AnError.Error())
			require.Nil(t, post)
		})
	})
}
