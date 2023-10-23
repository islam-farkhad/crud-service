package utils

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetIDFromQueryParams extracts and validates the "id" query parameter from the request.
func GetIDFromQueryParams(req *http.Request) (int64, bool) {
	vars := mux.Vars(req)
	idStr := vars["id"]
	if idStr == "" {
		return 0, false
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, false
	}
	return id, true
}
