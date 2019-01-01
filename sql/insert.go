package sql

import (
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
		"INSERT INTO sensorData (" +
			"runNumber, batteryVoltage, groundSpeed, intakeTemperature, lKillSwitch, rKillSwitch, mKillSwitch, " +
			"longitude, latitude, logTime, lapNumber, secondaryBatteryVoltage, wheelSpeed, windSpeed, systemCurrent, " +
			"coolantTemperature" +
			")" +
			"VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
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
