package api

import (
	"fmt"
	"net/http"
)

// ValidateMethod validates the HTTP method from the request passed
// Returns HTTP status code and error message
func ValidateMethod(r *http.Request, method string) (int, error) {
	if r.Method != method {
		return http.StatusMethodNotAllowed, fmt.Errorf("Method not allowed")
	}

	// No errors here!
	return http.StatusOK, nil
}

// ValidateMethods validates the HTTP method from the request passed
// Returns HTTP status code and error message
func ValidateMethods(r *http.Request, methods ...string) (int, error) {
	// Guilty until proven innocent
	err := fmt.Errorf("Method not allowed")

	for _, method := range methods {
		if r.Method == method {
			// Request method matches one of the allowed methods
			err = nil
			break
		}
	}

	if err != nil {
		return http.StatusMethodNotAllowed, err
	}

	// No errors here!
	return http.StatusOK, nil
}

// TryGetQueryArg attempts to retrieve a query value from the request matching key
func TryGetQueryArg(r *http.Request, key string) (string, error) {
	// Attempt to read first value paired with key from query
	value := r.URL.Query().Get(key)

	// If the value is empty, the parameter is missing
	if value == "" {
		return "", fmt.Errorf("Missing parameter %s", key)
	}

	// Got the value
	return value, nil
}
