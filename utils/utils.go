package utils

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
)

func Ptr[T any](v T) *T {
	return &v
}

// WriteJSON200 sets the Content-Type header and encodes the response as JSON with a 200 OK status
func WriteJSON200(w http.ResponseWriter, data interface{}) {
	writeJSONResponse(w, http.StatusOK, data)
}

// WriteJSONResponse sets the Content-Type header, status code, and encodes the response as JSON
func writeJSONResponse(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Printf("WriteJSONResponse: Error encoding response: %v", err)
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

// GetRandomInRange generates a random integer between min and max (inclusive)
func GetRandomInRange(min, max int) int {
	if min > max {
		log.Printf("Invalid range: min (%d) is greater than max (%d)\n", min, max)
		return -1 // -1 equal to error
	}
	return rand.Intn(max-min+1) + min
}
