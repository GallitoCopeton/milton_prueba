package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ReturnJson(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	// Encode the product struct into JSON
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		// Return an error if the JSON encoding fails
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the JSON response
	_, err = w.Write(jsonBytes)
	if err != nil {
		// Return if returning response gives an error
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func WriteErrorResponse(w http.ResponseWriter, status int, msg string, error error) {
	resp := make(map[string]string)
	resp["message"] = msg

	// Log error
	fmt.Println(error)

	// Return error JSON
	ReturnJson(w, resp)
	w.WriteHeader(status)
}
