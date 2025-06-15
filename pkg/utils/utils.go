package utils

import (
	"encoding/json"
	"net/http"
)

// parseRequestBody reads the JSON body from the given HTTP request
// and decodes it into the provided bodyContent, which can be any struct or interface.
// bodyContent must be a pointer to the variable where the decoded data will be stored.
// any will allow to store *anyObject passed to func
func ParseBody(r *http.Request, bodyContent any) {
	err := json.NewDecoder(r.Body).Decode(&bodyContent)
	if err != nil {
		return
	}
}
