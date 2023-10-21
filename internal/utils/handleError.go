package utils

import (
	"log"
	"net/http"
)

// HandleError writes an HTTP error response with the provided status code and error message.
// It also logs the error and status code for debugging.
func HandleError(w http.ResponseWriter, status int, err error) {
	w.WriteHeader(status)
	_, errWrite := w.Write([]byte(err.Error()))
	if errWrite != nil {
		log.Printf("Error writing response: %v", errWrite)
	}

	log.Printf("HTTP Status: %d, Error: %v", status, err)
}
