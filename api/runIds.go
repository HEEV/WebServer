package api

import (
	"fmt"
	"net/http"

	"github.com/HEEV/WebServer/datastore"

	log "github.com/sirupsen/logrus"
)

// NOTE: This endpoint is disabled until more clarification is gained on
// how it is supposed to operate

const runIdsQuery string = `
SELECT Cars.Name, a.LogTime, a.RunNumber 
FROM Cars
JOIN (
    SELECT MIN(SensorData.CarId) as CarId,
           MIN(SensorData.LogTime) as LogTime,
           SensorData.RunNumber 
    FROM SensorData 
    GROUP BY SensorData.RunNumber
) a 
ON a.CarId = Cars.Id
ORDER BY a.LogTime
DESC
`

// RunIdsHandler handles retrieval of data for /runIds endpoint
// Returns: A string of the data to return
func RunIdsHandler(r *http.Request) (string, error) {
	// Validate the request was via GET method
	_, err := ValidateMethod(r, "GET")
	if err != nil {
		log.Error(err)
		return "", err
	}

	log.Infof("Handling Run IDs request...")

	db := datastore.GetDatabase("data/test.sqlite")

	//Make sure there is no error when grabbing the data
	if db == nil {
		err := fmt.Errorf("Unable to connect to database for Run IDs")
		log.Error(err)
		return "", err
	}

	rows, err := db.Query(runIdsQuery)

	if rows == nil {
		httpErr := fmt.Errorf("Unable to connect to database for GraphHandler")
		log.Error(httpErr)
		log.Error(err)
		return "", httpErr
	}

	//Use the data from sql query to send back carName as a string
	var runId string
	var time string
	var car string
	var combined string

	for rows.Next() {
		rows.Scan(&runId)
		rows.Scan(&time)
		rows.Scan(&car)
		combined += runId + " " + time + " " + " " + car + " \n"
	}

	if err != nil {
		httpErr := fmt.Errorf("Failed to scan row for CSVHandler")
		log.Error(httpErr)
		log.Error(err)
		return "", httpErr
	}
	return combined, nil
}
