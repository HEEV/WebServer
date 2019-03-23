package api

import "net/http"

const carNameQuery string = "SELECT c.Name FROM Cars c WHERE c.Id = %s;"

// CarNameHandler handles retrieval of data for /carName endpoint
// Returns: A string of the data to return
func CarNameHandler(r *http.Request) string {
	return ""
}
