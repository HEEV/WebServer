package api

import (
	"fmt"
	"net/http"

	"github.com/HEEV/WebServer/sql"

	log "github.com/sirupsen/logrus"
)

// CSVHandler handles retrieval of data for /csv endpoint
// Returns: A string of the data to return, int of HTTP status, and an error
func CSVHandler(r *http.Request) (string, int, error) {
	// Validate the request was via POST method
	code, err := ValidateMethod(r, "POST")
	if err != nil {
		log.Error(err)
		return "", code, err
	}

	// Attempt to get runId query argument
	runId, err := TryGetQueryArg(r, "runId")
	if err != nil {
		log.Error(err)
		return "", http.StatusBadRequest, err
	}

	///Grab the database
	db := sql.GetDatabase("data/test.sqlite")

	//Make sure there is no error when grabbing the data
	if db == nil {
		err := fmt.Errorf("Unable to connect to database for CSVHandler")
		log.Error(err)
		return "", http.StatusInternalServerError, err
	}

	//Do the sql query
	rows, err := db.Query("SELECT * FROM SensorData WHERE RunNumber = ?;", runId)

	if err != nil {
		err := fmt.Errorf("Unable to retrieve row for CSVHandler")
		log.Error(err)
		return "", http.StatusInternalServerError, err
	}

	//This is gotten from stack overflow
	cols, err := rows.Columns()
	if err != nil {
		httpErr := fmt.Errorf("Failed to get columns for CSVHandler")
		log.Error(httpErr)
		log.Error(err)
		return "", http.StatusInternalServerError, httpErr
	}

	rawResult := make([][]byte, len(cols))
	runData := make([]interface{}, len(cols))
	for i, raw := range rawResult {
		runData[i] = raw // Put pointers to each string in the interface slice
	}

	//TODO: grab data from the row of run id and change it to cvs
	csv := ""
	for rows.Next() {
		err = rows.Scan(&runData)
		if err != nil {
			httpErr := fmt.Errorf("Failed to scan row for CSVHandler")
			log.Error(httpErr)
			log.Error(err)
			return "", http.StatusInternalServerError, httpErr
		}
		for i := 0; i < (len(cols)); i++ {
			temp := string(rawResult[i])
			csv += temp + ","
		}
		csv = "\n"
	}

	//Use the data from sql query to send back carName as a string
	err = rows.Scan(&runData)
	if err != nil {
		httpErr := fmt.Errorf("Failed to scan row for CSVHandler")
		log.Error(httpErr)
		log.Error(err)
		return "", http.StatusInternalServerError, httpErr
	}

	//Create our csv formatted string runData is formatted like [col][row]
	return csv, http.StatusOK, nil
}
