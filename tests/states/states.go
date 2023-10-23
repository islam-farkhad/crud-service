package states

import "time"

// General posts fields for testing
const (
	Post1ID = int64(500001)
	Post2ID = int64(500002)

	Post1Content = "post1"

	Post1Likes = int64(10)
)

// Post1CreatedAt General CreatedAt of post for testing
var Post1CreatedAt = time.Date(1999, 8, 25, 14, 10, 7, 0, time.UTC)

// General comment fields for testing
const (
	Comment1ID = int64(7001)
	Comment2ID = int64(7002)

	Comment1Content = "awesome post!"
	Comment2Content = "hater is here!"
)
