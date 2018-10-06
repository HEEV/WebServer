package sql

import (
	"../packets"

	log "github.com/sirupsen/logrus"
)

// LogToDatabase takes the information from the tablet and updates the database
func LogToDatabase(data packets.LogData) {
	log.Info("Updating database")
}
