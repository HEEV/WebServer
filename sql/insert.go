package sql

import (
	"../packets"

	log "github.com/sirupsen/logrus"
)

// LogToDatabase takes the information from the tablet and updates the database
func LogToDatabase(data packets.LogData) {
	log.Info("Updating database...")

	// Get DB connection
	db := GetDatabase()

	// Create the update SQL statement
	statement, err := db.Prepare("")
	if err != nil {
		log.Error("Error during database update with log data...")
		log.Error(err)
		return
	}
	statement.Exec()
}
