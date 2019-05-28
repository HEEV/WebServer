package api

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"net/http"

	"github.com/HEEV/WebServer/datastore"
	"github.com/HEEV/WebServer/packets"

	log "github.com/sirupsen/logrus"
)

const query string = "SELECT * FROM SensorData WHERE RunNumber = ?;"

var headerSlice = []string{
	"androidId",
	"batteryVoltage",
	"secondaryBatteryVoltage",
	"carId",
	"coolantTemperature",
	"groundSpeed",
	"intakeTemperature",
	"latitude",
	"longitude",
	"lapNumber",
	"logTime",
	"lKillSwitch",
	"mKillSwitch",
	"rKillSwitch",
	"runNumber",
	"systemCurrent",
	"wheelRpm",
	"windSpeed",
}

// CSVHandler handles retrieval of data for /csv endpoint
// Returns: A string of the data to return, int of HTTP status, and an error
func CSVHandler(r *http.Request) (string, int, error) {
	// Validate the request was via POST method
	code, err := ValidateMethod(r, "POST")
	if err != nil {
		log.Error(err)
		return "", code, err
	}

	log.Infof("Handling CSV data request...")

	// Attempt to get runId query argument
	runId, err := TryGetQueryArg(r, "runId")
	if err != nil {
		log.Error(err)
		return "", http.StatusBadRequest, err
	}

	// Grab the database
	db := datastore.GetDatabase("data/test.sqlite")

	// Make sure there is no error when retrieving the database connection
	if db == nil {
		err := fmt.Errorf("Unable to connect to database for CSVHandler")
		log.Error(err)
		return "", http.StatusInternalServerError, err
	}

	// Retrieve data related to specific run number. Should only be one run
	rows, err := db.Queryx(query, runId)
	if err != nil {
		log.Error("Unable to retrieve sensor data for CSVHandler")
		log.Error(err)
		return "", http.StatusInternalServerError, fmt.Errorf(internalServerErrMsg)
	}

	var curRow packets.DBSensorData

	// Attempt to fill initial record with data
	dbRows := make([]packets.DBSensorData, 1)
	for rows.Next() {
		if err := rows.StructScan(&curRow); err != nil {
			log.Error(err)
			return "", http.StatusInternalServerError, fmt.Errorf(internalServerErrMsg)
		}

		dbRows = append(dbRows, curRow)
	}

	// Should only ever have 1 row returned, warn if otherwise
	if len(dbRows) > 1 {
		log.Warn("Multiple records retrieved from table SensorData for run number", runId)
	}

	// Convert rows into CSV format
	buf := new(bytes.Buffer)
	dbWriter := csv.NewWriter(buf)

	// Write header
	dbWriter.Write(headerSlice)

	for _, row := range dbRows {
		dbWriter.Write(row.ToCSVString())
	}
	dbWriter.Flush()

	return buf.String(), http.StatusOK, nil
}
