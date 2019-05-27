package datastore

import (
	"fmt"
	"sync"

	// sqlite3 library for use by database/sql
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

	log "github.com/sirupsen/logrus"
)

// Keeps track of the single-instance connection to the database
var once sync.Once
var dbConn *sqlx.DB

func GetDatabase(dbFile string) *sqlx.DB {
	// Attempts to connect to SQLite database and set the single instance dbConn
	// Only is run once, after that it is ignored and the DB is already connected
	once.Do(func() {
		// Attempt to connect to the local SQLite database
		db, err := ConnectToLocalDB(dbFile)
		if err != nil {
			// Log stuff
			log.Error("Error occurred during initial DB connection...")
			log.Error(err)
			return
		}

		// Update the single instance
		dbConn = db
	})

	return dbConn
}

// ConnectToLocalDB connects to the local sqlite3 database used for testing
func ConnectToLocalDB(dbFile string) (*sqlx.DB, error) {
	db, err := sqlx.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// GetNextRunNumber retrieves the current run number from the SQL database,
// returning -1 and an error if retrieval fails
func GetNextRunNumber(androidID string) (int, error) {
	db := GetDatabase("data/test.sqlite")
	if db == nil {
		log.Error("Unable to retrieve next run number")
		return -1, fmt.Errorf("unable to open database to retrieve next run number")
	}

	// Generate the SQL statement to retrieve the current max run number
	row, err := db.Query("SELECT MAX(runNumber) FROM sensorData WHERE androidId = ?", androidID)
	if err != nil {
		log.Error("Error occurred during run number retrieval")
		return -1, err
	}

	// Retrieve information based on query
	var runNumber int
	row.Scan(&runNumber)

	// Increment the previous highest run number by 1 to get the next run number to use
	return runNumber + 1, nil
}
