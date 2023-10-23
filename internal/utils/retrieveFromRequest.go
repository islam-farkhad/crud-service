package utils

import (
	"io"
	"net/http"
)

// RetrieveBody reads body from request and returns it.
func RetrieveBody(req *http.Request) ([]byte, int) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, http.StatusInternalServerError
	}
	return body, http.StatusOK
}

// RetrieveID gets id from query params of request and returns it.
func RetrieveID(req *http.Request) (int64, int) {
	ID, ok := GetIDFromQueryParams(req)
	if !ok {
		return 0, http.StatusBadRequest
	}
	return ID, http.StatusOK
}
