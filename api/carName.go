package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/HEEV/WebServer/datastore"

	log "github.com/sirupsen/logrus"
)

const carNameQuery string = "SELECT c.Name FROM Cars c WHERE c.Id = %s;"

// CarNameHandler handles retrieval of data for /carName endpoint
// Returns: A string of the data to return
func CarNameHandler(r *http.Request) (string, error) {
	if r.Method != "POST" {
		return "", fmt.Errorf("Method not allowed")
	}

	carId := r.URL.Query().Get("CarId")

	if carId == "" {
		return "", fmt.Errorf("CarId parameter not passed")
	}

	// Grab the database
	db := datastore.GetDatabase("data/test.sqlite")

	// Make sure there is no error when grabbing the data
	if db == nil {
		err := fmt.Errorf("Unable to connect to database for CarNameHandler")
		log.Error(err)
		return "", err
	}

	// Query the Cars table for the matching car ID
	row := db.QueryRow("SELECT c.Name FROM Cars c WHERE c.ID = ?", carId)

	// Use the data from sql query to send back carName as a string
	var carName string
	err := row.Scan(&carName)
	if err == sql.ErrNoRows {
		err := fmt.Errorf("Not found")
		return "", err
	}
	if err != nil {
		log.Error(err)
		return "", fmt.Errorf("Error during DB row scanning")
	}

	return carName, nil
}
