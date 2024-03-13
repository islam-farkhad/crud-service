package utils

import (
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

// GetIDFromQueryParams extracts and validates the "id" query parameter from the request.
func GetIDFromQueryParams(req *http.Request) (int64, bool) {
	vars := mux.Vars(req)
	idStr := vars["id"]
	if idStr == "" {
		log.Println("id is not provided in query params.")
		return 0, false
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		log.Printf("Invalid id = %s provided. Could not convert to int64\n", idStr)
		return 0, false
	}
	return id, true
}
