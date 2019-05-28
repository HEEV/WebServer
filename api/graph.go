package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/HEEV/WebServer/datastore"
	"github.com/HEEV/WebServer/packets"

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

	log.Infof("Handling Graph request...")

	// Attempt to get runId query argument
	runId, err := TryGetQueryArg(r, "runId")
	if err != nil {
		log.Error(err)
		return "", err
	}

	// Grab the database
	db := datastore.GetDatabase("data/test.sqlite")

	// Make sure there is no error when grabbing the data
	if db == nil {
		log.Error("Unable to connect to database for GraphHandler")
		return "", fmt.Errorf(internalServerErrMsg)
	}

	// Execute SQL query to retrieve sensor data for provided run
	rows, err := db.Queryx("SELECT * FROM SensorData WHERE RunNumber = ?;", runId)
	if rows == nil {
		log.Error("Unable to retrieve sensor data for GraphHandler")
		log.Error(err)
		return "", fmt.Errorf(internalServerErrMsg)
	}

	var curRow packets.DBSensorData

	dbRows := make([]packets.DBSensorData, 0)
	for rows.Next() {
		if err := rows.StructScan(&curRow); err != nil {
			log.Error(err)
			return "", fmt.Errorf(internalServerErrMsg)
		}

		dbRows = append(dbRows, curRow)
	}

	// Should only ever have 1 row returned, warn if otherwise
	if len(dbRows) > 1 {
		log.Warn("Multiple records retrieved from table SensorData for run number", runId)
	}

	// Marshall structs to JSON
	resp, err := json.MarshalIndent(dbRows, "", "    ")

	if err != nil {
		log.Error("Failed to marshal JSON result")
		log.Error(err)
		return "", fmt.Errorf(internalServerErrMsg)
	}

	return string(resp), nil
}
