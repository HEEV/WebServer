package api

import (
	"fmt"
	"net/http"

	"github.com/HEEV/WebServer/datastore"

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
	rows, err := db.Query("SELECT * FROM SensorData WHERE RunNumber = ?;", runId)
	if rows == nil {
		log.Error("Unable to retrieve sensor data for GraphHandler")
		log.Error(err)
		return "", fmt.Errorf(internalServerErrMsg)
	}

	// // TODO: Make this actually return something...

	// cols, err := rows.Columns()
	// if err != nil {
	// 	log.Error("Failed to get columns", err)
	// 	return "", err
	// }

	// // Use the data from sql query to send back carName as a string
	// rawResult := make([][]byte, len(cols))
	// dest := make([]interface{}, len(cols))
	// var runData string

	// for i := range rawResult {
	// 	dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
	// }

	// for rows.Next() {
	// 	err = rows.Scan(dest...)
	// 	if err != nil {
	// 		fmt.Println("Failed to scan row", err)
	// 		return "", err
	// 	}

	// 	for raw := range rawResult {
	// 		runData += string(raw) + " "
	// 	}
	// 	runData += "\n"
	// }

	// if err != nil {
	// 	httpErr := fmt.Errorf("Failed to scan row for CSVHandler")
	// 	log.Error(httpErr)
	// 	log.Error(err)
	// 	return "", httpErr
	// }

	// return runData, nil
	return "", nil
}
