package utils

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// GetIDFromQueryParams extracts and validates the "id" query parameter from the request.
// It returns the parsed ID and a boolean indicating whether the ID extraction was successful.
// If extraction fails, it also writes an error response to the provided http.ResponseWriter.
func GetIDFromQueryParams(w http.ResponseWriter, req *http.Request) (int64, bool) {
	vars := mux.Vars(req)
	idStr := vars["id"]
	if idStr == "" {
		HandleError(w, http.StatusBadRequest, fmt.Errorf("provide id in query parameters"))
		return 0, false
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		HandleError(w, http.StatusBadRequest, fmt.Errorf("id should be a number"))
		return 0, false
	}
	return id, true
}
