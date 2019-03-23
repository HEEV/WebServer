package api

import (
	"net/http"

	"github.com/HEEV/WebServer/sql"

	log "github.com/sirupsen/logrus"
)

// CSVHandler handles retrieval of data for /csv endpoint
// Returns: A string of the data to return
func CSVHandler(r *http.Request) string {
	if r.Method == "POST"{
		var runId = r.URL.Query()["runId"]

		///Grab the database
		db := sql.GetDatabase("data/test.sqlite")

		//Make sure there is no error when grabbing the data
		if db == nil {
			log.Error("Unable to connect to database for CSVHandler")
			return ""
		}

		//Do the sql query
		rows, err :=  db.Query("SELECT * FROM SensorData WHERE RunNumber = ?;", runId)

		if(err != nil){
			log.Error("Unable to connect to database for CSVHandler")
			return ""
		}

		//This is gotten from stack overflow
		cols, err := rows.Columns()
		if err != nil {
			log.Error("Failed to get columns", err)
			return ""
		}

		rawResult := make([][]byte, len(cols))
		runData := make([]interface{}, len(cols))
		for i, _ := range rawResult {
			runData[i] = &rawResult[i] // Put pointers to each string in the interface slice
		}
		//TODO: grab data from the row of run id and change it to cvs
		var csv string = ""
		for rows.Next(){
			err = rows.Scan(&runData)
			if err != nil{
				log.Error("Failed to Scan")
				return ""
			}
			for i := 0; i < (len(cols)); i++{
				var temp string = string(rawResult[i])
				csv += temp + ","
			}
			csv = "\n"
		}

		//Use the data from sql query to send back carName as a string
		err = rows.Scan(&runData)
		if err != nil {
			return ""
		}

		//Create our csv formatted string runData is formatted like [col][row]
		return csv
	}
	return ""
}
