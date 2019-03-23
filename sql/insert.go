package sql

import (
	"database/sql"

	"github.com/HEEV/WebServer/packets"
	log "github.com/sirupsen/logrus"
)

// LogToDatabase takes the information from the tablet and updates the database
func LogToDatabase(data packets.LogData) {
	log.Info("Updating database...")

	// Get DB connection
	db := getDatabase("data/test.sqlite")

	// Create the update SQL statement
	statement, err := db.Prepare(
		"INSERT INTO sensorData (" +
			"runNumber, batteryVoltage, groundSpeed, intakeTemperature, lKillSwitch, rKillSwitch, mKillSwitch, " +
			"longitude, latitude, logTime, lapNumber, secondaryBatteryVoltage, wheelSpeed, windSpeed, systemCurrent, " +
			"coolantTemperature, carId" +
			")" +
			"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? ,?, ?)",
	)

	if err != nil {
		log.Error("Error during database update with log data...")
		log.Error(err)
		return
	}

	statement.Exec(
		data.RunNumber,
		data.BatteryVoltage,
		data.GroundSpeed,
		data.IntakeTemperature,
		data.LKillSwitch,
		data.RKillSwitch,
		data.MKillSwitch,
		data.Longitude,
		data.Latitude,
		data.LogTime,
		data.LapNumber,
		data.SecondaryBatteryVoltage,
		data.WheelRpm,
		data.WindSpeed,
		data.SystemCurrent,
		data.CoolantTemperature,
	)

}

// androidToCar uses the androidId to grab the carId
func androidToCar(db *sql.DB, androidID string) (string, error) {
	// Gets the database
	rows, err := db.Query("SELECT CarId."+
		" FROM CarTablet. "+
		" WHERE AndroidId = ?", androidID)

	if err != nil {
		return "", err
	}

	var carID string

	if err = rows.Scan(&carID); err != nil {
		return "", err
	}

	return carID, nil
}

// This gets the current run number
func getNextRunNumber(androidID string) (int, error) {
	// Retrieve database connection
	db := getDatabase("data/test.sqlite")

	//Inserts the data in and prints out the run number
	carID, err := androidToCar(db, androidID)
	if err != nil {
		return -1, err
	}

	log.Infof("Car ID from DB: %s", carID)

	rows, err := db.Query("SELECT MAX(RunNumber)."+
		" FROM SensorData."+
		" WHERE Car id = ?", androidID)

	if err != nil {
		return -1, err
	}

	var carRunNum int

	if err = rows.Scan(&carRunNum); err != nil {
		return -1, err
	}

	log.Infof("Current Run Number: %d", carRunNum)

	// Get the next car number
	var nextRunNumber = carRunNum + 1
	db.Close()

	return nextRunNumber, nil
}
