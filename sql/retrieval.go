package sql

import (
	"sync"

	"database/sql"
	_"github.com/mattn/go-sqlite3"
)

// Keeps track of the single-instance connection to the database
var once sync.Once
var dbConn *sql.DB

func GetDatabase() *sql.DB {
	// Attempts to connect to SQLite database and set the single instance dbConn
	// Only is run once, after that it is ignored and the DB is already connected
	once.Do(func() {
		// Attempt to connect to the local SQLite database
		db, err := ConnectToLocalDB("../data/local_dump.sqlite")
		if (err != nil) {
			// Log stuff
			return
		}

		// Update the single instance
		dbConn = db
	})

	return dbConn
}

// ConnectToLocalDB connects to the local sqlite3 database used for testing
func ConnectToLocalDB(dbFile string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}

	return db, nil
}

// GetNextRunNumber retrieves the current run number from the SQL database,
// returning -1 and an error if retrieval fails
func GetNextRunNumber(androidID string) (int, error) {
	return 0, nil
}
