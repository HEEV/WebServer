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
	db := GetDatabase("data/test.sqlite")

	// Create the update SQL statement
	statement, err := db.Prepare(
		"INSERT INTO SensorData (" +
			"runNumber, batteryVoltage, groundSpeed, intakeTemperature, " +
			"lKillSwitch, rKillSwitch, mKillSwitch, longitude, latitude, " +
			"logTime, lapNumber, secondaryBatteryVoltage, wheelSpeed, " +
			"windSpeed, systemCurrent, coolantTemperature, carId, androidId" +
			")" +
			"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
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
		data.CarId,
	)

}

// androidToCar uses the androidId to grab the carId
func androidToCar(db *sql.DB, androidID string) (string, error) {
	// Gets the database
	row := db.QueryRow(
		"SELECT CarId FROM CarTablet WHERE AndroidId = ?", androidID)

	var carID string

	if err := row.Scan(&carID); err != nil {
		return "", err
	}
	return carID, nil
}
