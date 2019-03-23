package api

import (
	"net/http"

	"github.com/HEEV/WebServer/sql"

	log "github.com/sirupsen/logrus"
)

const carNameQuery string = "SELECT c.Name FROM Cars c WHERE c.Id = %s;"

// CarNameHandler handles retrieval of data for /carName endpoint
// Returns: A string of the data to return
func CarNameHandler(r *http.Request) string {
	if r.Method != "POST" || r.URL.Query()["CarId"] == nil {
		return ""
	}

	carId := r.URL.Query()["CarId"]

	// Grab the database
	db := sql.GetDatabase("data/test.sqlite")

	// Make sure there is no error when grabbing the data
	if db == nil {
		log.Error("Unable to connect to database for CarNameHandler")
		return ""
	}
	row := db.QueryRow("Select c.Name FROM Cars c WHERE c.ID = ?", carId)

	if row == nil {
		log.Error("Unable to connect to database for CarNameHandler")
		return ""
	}

	// Use the data from sql query to send back carName as a string
	var carName string
	err := row.Scan(&carName)
	if err != nil {
		return ""
	}

	return carName

	return ""
}
