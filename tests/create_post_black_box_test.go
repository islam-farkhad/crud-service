//go:build integration

package tests

import (
	"bytes"
	"context"
	db2 "crud-service/internal/pkg/db"
	"crud-service/internal/pkg/repository"
	"crud-service/internal/utils"
	"crud-service/tests/states"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/jackc/pgconn"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

const (
	createURL = "http://localhost:8080/post"
)

type db interface {
	Exec(ctx context.Context, query string, args ...interface{}) (pgconn.CommandTag, error)
	Get(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

// CreatePostTestSuite is a create post test suite
type CreatePostTestSuite struct {
	suite.Suite
	db db
}

func newCreatePostTestSuite(db db) *CreatePostTestSuite {
	return &CreatePostTestSuite{
		db: db,
	}
}

// SetupTest performs table cleaning
func (suite *CreatePostTestSuite) SetupTest() {
	_, err := suite.db.Exec(context.Background(), "delete from posts")
	if err != nil {
		log.Println(err)
	}
}

// TestCreateValid performs valid case
func (suite *CreatePostTestSuite) TestCreateValid() {

	//arrange
	var jsonStr = []byte(fmt.Sprintf(`{"content":"%s","likes":%d}`, states.Post1Content, states.Post1Likes))
	req, err := http.NewRequest("POST", createURL, bytes.NewBuffer(jsonStr))
	assert.NoError(suite.T(), err)
	req.Header.Set("Content-Type", "application/json")

	newPost := &repository.Post{}
	assert.Equal(suite.T(), int64(0), newPost.ID)

	client := &http.Client{}

	// act
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Error closing response body: %v", err)
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading bytes: %v", err)
	}
	if err = json.Unmarshal(body, newPost); err != nil {
		log.Printf("Could not unmarshal body: %v", err)
	}

	//assert
	assert.EqualValues(suite.T(), http.StatusOK, resp.StatusCode)

	assert.NotEqual(suite.T(), int64(0), newPost.ID)
	assert.Equal(suite.T(), states.Post1Content, newPost.Content)
	assert.Equal(suite.T(), states.Post1Likes, newPost.Likes)

	// check in db
	var retrievedPost repository.Post
	err = suite.db.Get(context.Background(), &retrievedPost, "SELECT * FROM posts WHERE id = $1", newPost.ID)
	if err != nil {
		log.Printf("Error querying the database: %v", err)
	}

	assert.Equal(suite.T(), states.Post1Content, retrievedPost.Content)
	assert.Equal(suite.T(), states.Post1Likes, retrievedPost.Likes)
}

func TestCreateSuite(t *testing.T) {
	ctx := context.Background()

	config := utils.GetEnvDBConnectionConfig()
	config.Host = "localhost"

	dbNew, err := db2.NewDB(ctx, utils.MakeDBConnStr(config))
	if err != nil {
		log.Fatal(err)
	}

	suite.Run(t, newCreatePostTestSuite(dbNew))
}
