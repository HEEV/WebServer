package api

import (
	"net/http"

	"github.com/HEEV/WebServer/sql"
	log "github.com/sirupsen/logrus"
)

// GraphHandler handles retrieval of data for /graph endpoint
// Returns: A string of the data to return
func GraphHandler(r *http.Request) string {
	if r.Method != "POST" || r.URL.Query()["runId"] == nil {
		return ""
	}

	runId := r.URL.Query()["runId"]

	///Grab the database
	db := sql.GetDatabase("data/test.sqlite")

	//Make sure there is no error when grabbing the data
	if db == nil {
		log.Error("Unable to connect to database for GraphHandler")
		return ""
	}

	//Do the sql query
	row, err := db.Query("SELECT * FROM SensorData WHERE RunNumber = ?;", runId)

	if row == nil {
		log.Error("Unable to connect to database for GraphHandler")
		return ""
	}

	//TODO: change this so that it takes a whole array
	//Use the data from sql query to send back carName as a string
	var carName string
	error := row.Scan(&carName)
	if error != nil {
		return ""
	}

	return carName
}
