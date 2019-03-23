package api

import (
	"fmt"
	"net/http"

	"github.com/HEEV/WebServer/sql"
	log "github.com/sirupsen/logrus"
)

// GraphHandler handles retrieval of data for /graph endpoint
// Returns: A string of the data to return
func GraphHandler(r *http.Request) (string, error) {
	// Validate the request was via POST method
	_, err := ValidateMethod(r, "POST")
	if err != nil {
		log.Error(err)
		return "", err
	}

	// Attempt to get runId query argument
	runId, err := TryGetQueryArg(r, "runId")
	if err != nil {
		log.Error(err)
		return "", err
	}

	///Grab the database
	db := sql.GetDatabase("data/test.sqlite")

	//Make sure there is no error when grabbing the data
	if db == nil {
		err := fmt.Errorf("Unable to connect to database for GraphHandler")
		log.Error(err)
		return "", err
	}

	//Do the sql query
	row, err := db.Query("SELECT * FROM SensorData WHERE RunNumber = ?;", runId)

	if row == nil {
		httpErr := fmt.Errorf("Unable to connect to database for GraphHandler")
		log.Error(httpErr)
		log.Error(err)
		return "", httpErr
	}

	//TODO: change this so that it takes a whole array
	//Use the data from sql query to send back carName as a string
	var carName string
	err = row.Scan(&carName)
	if err != nil {
		httpErr := fmt.Errorf("Failed to scan row for CSVHandler")
		log.Error(httpErr)
		log.Error(err)
		return "", httpErr
	}

	return carName, nil
}
