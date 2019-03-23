package api

import (
	"fmt"
	"github.com/HEEV/WebServer/sql"
	log "github.com/sirupsen/logrus"
	"net/http"
)
// RunIdsHandler handles retrieval of data for /runIds endpoint
// Returns: A string of the data to return
func RunIdsHandler(r *http.Request) (string, error)  {
	// Validate the request was via POST method
	_, err := ValidateMethod(r, "GET")
	if err != nil {
		log.Error(err)
		return "", err
	}

	db := sql.GetDatabase("data/test.sqlite")

	//Make sure there is no error when grabbing the data
	if db == nil {
		err := fmt.Errorf("Unable to connect to database for GraphHandler")
		log.Error(err)
		return "", err
	}

	rows, err := db.Query("SELECT Cars.Name, a.LogTime, a.RunNumber " +
		"FROM CarsJOIN (SELECT MIN(SensorData.CarId) as CarId, " +
		"MIN(SensorData.LogTime) LogTime, SensorData.RunNumber" +
		"FROM SensorData GROUP BY SensorData.RunNumber) a " +
		"ON a.CarId = Cars.Id" +
		"ORDER BY a.LogTime DESC;")

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
	return "", nil
}
