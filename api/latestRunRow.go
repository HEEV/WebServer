package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HEEV/WebServer/datastore"
	"github.com/HEEV/WebServer/packets"

	log "github.com/sirupsen/logrus"
)

const latestRunNumQuery string = "SELECT * FROM SensorData ORDER BY id DESC LIMIT 1"

// LatestRunHandler handles retrieval of data for /latestRun endpoint
// Returns: A string of the data to return
func LatestRunHandler(r *http.Request) (string, error) {
	// Validate the request was via GET method
	_, err := ValidateMethod(r, "GET")
	if err != nil {
		log.Error(err)
		return "", err
	}

	log.Infof("Handling latest run request...")

	// Grab the database
	db := datastore.GetDatabase("data/test.sqlite")

	//Make sure there is no error when grabbing the data
	if db == nil {
		log.Error(fmt.Errorf("Unable to connect to database for CarNameHandler"))
		return "", fmt.Errorf(internalServerErrMsg)
	}

	// Do the sql query
	row := db.QueryRowx(latestRunNumQuery)

	if row == nil {
		log.Error("Nil next row when querying LatestRunHandler")
		return "", fmt.Errorf(internalServerErrMsg)
	}

	// Store data retrieved from DB in struct
	var runData packets.DBSensorData
	row.StructScan(&runData)

	// Marshall response data struct into a byte array
	resultJSON, err := json.MarshalIndent(runData, "", "    ")
	if err != nil {
		log.Error("Failed to marshall latest run row response JSON")
		return "", fmt.Errorf(internalServerErrMsg)
	}

	return string(resultJSON), nil
}
