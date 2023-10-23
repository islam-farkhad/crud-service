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

func Test_CreatePost(t *testing.T) {
	t.Parallel()
	var (
		ctx         = context.Background()
		postRequest = addPostRequest{
			Content: states.Post1Content,
			Likes:   states.Post1Likes,
		}
	)

	t.Run("success", func(t *testing.T) {
		t.Parallel()

		// arrange
		s := setUp(t)
		defer s.tearDown()

		s.mockRepo.EXPECT().AddPost(gomock.Any(), gomock.Any()).Return(states.Post1ID, nil)

		post := fixtures.BuildPost().Valid().P()

		// act
		result, status := s.mockApp.CreatePost(ctx, post)

		// assert
		require.Equal(t, http.StatusOK, status)

		var postResponse repository.Post
		err := json.Unmarshal(result, &postResponse)
		require.NoError(t, err)

		assert.Equal(t, states.Post1ID, postResponse.ID)
		assert.Equal(t, postRequest.Content, postResponse.Content)
		assert.Equal(t, postRequest.Likes, postResponse.Likes)
	})

	t.Run("fail", func(t *testing.T) {
		t.Parallel()
		t.Run("repository error", func(t *testing.T) {
			t.Parallel()

			// arrange
			s := setUp(t)
			defer s.tearDown()

			s.mockRepo.EXPECT().AddPost(gomock.Any(), gomock.Any()).Return(int64(0), assert.AnError)

			post := fixtures.BuildPost().Valid().P()

			// act
			result, status := s.mockApp.CreatePost(ctx, post)

			// assert
			require.Equal(t, http.StatusInternalServerError, status)
			require.Contains(t, string(result), fmt.Sprintf("could not add post. err: %v", assert.AnError))
		})
	})
}

func Test_parseCreatePost(t *testing.T) {
	type args struct {
		body []byte
	}
	tests := []struct {
		name  string
		args  args
		want  *repository.Post
		want1 int
	}{
		{
			name:  "success",
			args:  args{[]byte(fmt.Sprintf("{\"content\":\"%s\",\"likes\":%d}", states.Post1Content, states.Post1Likes))},
			want:  fixtures.BuildPost().Content(states.Post1Content).Likes(states.Post1Likes).P(),
			want1: http.StatusOK,
		},
		{
			name:  "fail - invalid quotation marks",
			args:  args{[]byte(fmt.Sprintf("{\"content\":%s,\"likes\":\"%d\"}", states.Post1Content, states.Post1Likes))},
			want:  nil,
			want1: http.StatusBadRequest,
		},
		{
			name:  "fail",
			args:  args{[]byte(fmt.Sprintf("{\"not_a_content\":\"%s\",\"likes\":%d}", states.Post1Content, states.Post1Likes))},
			want:  nil,
			want1: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseCreatePost(tt.args.body)
			assert.Equalf(t, tt.want, got, "parseCreatePost(%v)", tt.args.body)
			assert.Equalf(t, tt.want1, got1, "parseCreatePost(%v)", tt.args.body)
		})
	}
}
