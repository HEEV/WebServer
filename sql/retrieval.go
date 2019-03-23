package sql

import (
	"fmt"
	"sync"

	"database/sql"
	//"github.com/mattn/go-sqlite3"

	log "github.com/sirupsen/logrus"
)

// Keeps track of the single-instance connection to the database
var once sync.Once
var dbConn *sql.DB

// Database SQLite file
const dbFile string = "data/test.sqlite"

func getDatabase(dbFile string) *sql.DB {
	// Attempts to connect to SQLite database and set the single instance dbConn
	// Only is run once, after that it is ignored and the DB is already connected
	once.Do(func() {
		// Attempt to connect to the local SQLite database
		db, err := connectToLocalDB(dbFile)
		if err != nil {
			// Log stuff
			log.Error(err)

			// Have to create a new sync object to allow next connection
			once = *new(sync.Once)
			return
		}

		// Update the single instance
		dbConn = db
	})

	return dbConn
}

// connectToLocalDB connects to the local sqlite3 database used for testing
func connectToLocalDB(dbFile string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// GetValue gets a value from the database based on the query
func GetValue(query string) ([]string, error) {
	// Response slice of strings

	// Get database to read from
	db := getDatabase("data/test.sqlite")
	if db == nil {
		log.Error("Unable to access database")
		return make([]string, 0), fmt.Errorf("unable to open database to retrieve next run number")
	}

	// Query database
	rows, err := db.Query(query)
	if err != nil {
		log.Error("Error occurred during data retrieval")
		log.Errorf("Query: %s", query)
		log.Error(err)
		return nil, err
	}

	return nil, nil
}

// GetNextRunNumber retrieves the current run number from the SQL database,
// returning -1 and an error if retrieval fails
func GetNextRunNumber(androidID string) (int, error) {
	db := getDatabase("data/test.sqlite")
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
