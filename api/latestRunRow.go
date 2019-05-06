package api

import (
	"fmt"
	"net/http"

	"github.com/HEEV/WebServer/datastore"

	log "github.com/sirupsen/logrus"
)

type RunRow struct {
	AndroidID               string `db:"androidId"`
	RunNumber               int    `db:"runNumber"`
	BatteryVoltage          float32
	GroundSpeed             float32
	IntakeTemperature       float32
	LKillSwitch             int
	MKillSwitch             int
	RKillSwitch             int
	Longitude               float32
	Latitude                float32
	LogTime                 string
	LapNumber               int
	SecondaryBatteryVoltage float32
	WheelRpm                float32
	WindSpeed               float32
	SystemCurrent           float32
	CoolantTemperature      float32
	CarId                   string
}

const latestRunNumQuery string = "SELECT * FROM SensorData ORDER BY id DESC LIMIT 1"

// LatestRunHandler handles retrieval of data for /latestRun endpoint
// Returns: A string of the data to return
func LatestRunHandler(r *http.Request) (string, error) {
	// Validate the request was via GET method
	_, err := ValidateMethod(r, "GET")
	if err != nil {
		log.Error(err)
		return "", err
	}

	// Grab the database
	db := datastore.GetDatabase("data/test.sqlite")

	//Make sure there is no error when grabbing the data
	if db == nil {
		log.Error(fmt.Errorf("Unable to connect to database for CarNameHandler"))
		return "", fmt.Errorf(internalServerErrMsg)
	}

	// Do the sql query
	row := db.QueryRowx(latestRunNumQuery)

	if row == nil {
		log.Error("Nil next row when querying LatestRunHandler")
		return "", fmt.Errorf(internalServerErrMsg)
	}

	var runData RunRow
	row.StructScan(&runData)
	log.Infof("%+v", runData)

	return "", nil
}
