package api

import (
	"github.com/HEEV/WebServer/sql"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// LatestRunHandler handles retrieval of data for /latestRun endpoint
// Returns: A string of the data to return
func LatestRunHandler(r *http.Request) string {
	///Grab the database
	db := sql.GetDatabase("data/test.sqlite")

	//Make sure there is no error when grabbing the data
	if db == nil {
		log.Error("Unable to connect to database for CarNameHandler")
		return ""
	}

	//Do the sql query
	row, err :=  db.Query("SELECT * FROM SensorData ORDER BY Id DESC LIMIT 1;")


	//TODO: Fix from here on down
	if(row == nil){
		log.Error("Unable to connect to database for CarNameHandler")
		return ""
	}

	//Use the data from sql query to send back carName as a string
	var runData string
	error := row.Scan(&runData)
	if error != nil {
		return ""
	}

	return carName

}
