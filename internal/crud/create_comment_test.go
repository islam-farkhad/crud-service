package crud

import (
	"context"
	"encoding/json"
	"fmt"
	"homework-3/internal/pkg/repository"
	"homework-3/tests/fixtures"
	"homework-3/tests/states"
	"net/http"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCreateComment(t *testing.T) {
	t.Parallel()
	var (
		ctx = context.Background()
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		// arrange
		s := setUp(t)
		defer s.tearDown()

		s.mockRepo.EXPECT().AddComment(gomock.Any(), gomock.Any()).Return(states.Comment1ID, nil)
		comment := fixtures.BuildComment().Valid().P()

		// act
		data, status := s.mockApp.CreateComment(ctx, comment)

		// assert
		require.Equal(t, http.StatusOK, status)

		var commentResponse repository.Comment
		err := json.Unmarshal(data, &commentResponse)
		require.NoError(t, err)

		assert.Equal(t, states.Comment1ID, commentResponse.ID)
		assert.Equal(t, comment.PostID, commentResponse.PostID)
		assert.Equal(t, comment.Content, commentResponse.Content)
	})

	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		t.Run("violating fk constraint", func(t *testing.T) {
			t.Parallel()

			// arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockRepo.EXPECT().AddComment(gomock.Any(), gomock.Any()).Return(states.Comment2ID, assert.AnError)
			comment := fixtures.BuildComment().Valid().P()

			// act
			data, status := s.mockApp.CreateComment(ctx, comment)

			// assert
			require.Equal(t, http.StatusBadRequest, status)
			require.Contains(t, string(data), fmt.Sprintf("could not add comment. err: %v", assert.AnError))
		})
	})
}

func Test_parseCreateComment(t *testing.T) {
	type args struct {
		body   []byte
		postID int64
	}
	tests := []struct {
		name  string
		args  args
		want  *repository.Comment
		want1 int
	}{
		{
			name: "success",
			args: args{
				body:   []byte(fmt.Sprintf(`{"content":"%s"}`, states.Comment1Content)),
				postID: states.Post1ID,
			},
			want:  fixtures.BuildComment().Content(states.Comment1Content).PostID(states.Post1ID).P(),
			want1: http.StatusOK,
		},
		{
			name: "fail - empty content",
			args: args{
				body:   []byte(fmt.Sprintf(`{"content":"%s"}`, "")),
				postID: states.Post1ID,
			},
			want:  nil,
			want1: http.StatusBadRequest,
		},
		{
			name: "fail - invalid body",
			args: args{
				body:   []byte(fmt.Sprintf(`{content:"%s"}`, "")), //no quotation marks
				postID: states.Post1ID,
			},
			want:  nil,
			want1: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseCreateComment(tt.args.body, tt.args.postID)
			assert.Equalf(t, tt.want, got, "parseCreateComment(%v, %v)", tt.args.body, tt.args.postID)
			assert.Equalf(t, tt.want1, got1, "parseCreateComment(%v, %v)", tt.args.body, tt.args.postID)
		})
	}
}
