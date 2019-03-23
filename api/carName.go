package api

import (
	"fmt"
	"net/http"

	"github.com/HEEV/WebServer/sql"

	log "github.com/sirupsen/logrus"
)

const carNameQuery string = "SELECT c.Name FROM Cars c WHERE c.Id = %s;"

// CarNameHandler handles retrieval of data for /carName endpoint
// Returns: A string of the data to return
func CarNameHandler(r *http.Request) (string, error) {
	if r.Method != "POST" {
		return "", fmt.Errorf("Method not allowed")
	}

	carId, exists := r.URL.Query()["CarId"]

	if !exists {
		return "", fmt.Errorf("CarId parameter not passed")
	}

	// Grab the database
	db := sql.GetDatabase("data/test.sqlite")

	// Make sure there is no error when grabbing the data
	if db == nil {
		err := fmt.Errorf("Unable to connect to database for CarNameHandler")
		log.Error(err)
		return "", err
	}
	row := db.QueryRow("Select c.Name FROM Cars c WHERE c.ID = ?", carId)

	if row == nil {
		err := fmt.Errorf("Unable to query database for CarNameHandler")
		log.Error(err)
		return "", err
	}

	// Use the data from sql query to send back carName as a string
	var carName string
	err := row.Scan(&carName)
	if err != nil {
		log.Error(err)
		return "", fmt.Errorf("Error during DB row scanning")
	}

	return carName, nil
}
