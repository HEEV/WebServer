package api

import (
	"fmt"
	"net/http"

	"github.com/HEEV/WebServer/datastore"

	log "github.com/sirupsen/logrus"
)

// LatestRunHandler handles retrieval of data for /latestRun endpoint
// Returns: A string of the data to return
func LatestRunHandler(r *http.Request) string {
	// Grab the database
	db := datastore.GetDatabase("data/test.sqlite")

	//Make sure there is no error when grabbing the data
	if db == nil {
		log.Error("Unable to connect to database for CarNameHandler")
		return ""
	}

	//Do the sql query
	rows, err := db.Query("SELECT * FROM SensorData ORDER BY Id DESC LIMIT 1;")

	//TODO: Fix from here on down
	cols, err := rows.Columns()
	if err != nil {
		log.Error("Failed to get columns", err)
		return ""
	}
	//Use the data from sql query to send back carName as a string
	rawResult := make([][]byte, len(cols))
	dest := make([]interface{}, len(cols))
	var runData string

	for i := range rawResult {
		dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
	}

	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			fmt.Println("Failed to scan row", err)
			return ""
		}

		for raw := range rawResult {
			runData += string(raw) + " "
		}
		runData += "\n"
	}

	if err != nil {
		httpErr := fmt.Errorf("Failed to scan row for CSVHandler")
		log.Error(httpErr)
		log.Error(err)
		return ""
	}

	return runData
}
