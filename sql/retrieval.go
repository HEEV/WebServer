package sql

import (
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

func GetDatabase(dbFile string) *sql.DB {
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

// NOTE: Use this function as an example, but this won't work with our
// database, the types are too disparate.

// GetValue gets a value from the database based on the query
// func GetValue(query string) ([][]interface{}, error) {
// 	// Response slice of strings
// 	var resp = make([][]interface{}, 0)

// 	// Get database to read from
// 	db := getDatabase("data/test.sqlite")
// 	if db == nil {
// 		log.Error("Unable to access database")
// 		return nil, fmt.Errorf("unable to open database to retrieve next run number")
// 	}

// 	// Query database
// 	rows, err := db.Query(query)
// 	defer rows.Close()
// 	if err != nil {
// 		log.Error("Error occurred during data retrieval")
// 		log.Errorf("Query: %s", query)
// 		log.Error(err)
// 		return nil, err
// 	}

// 	for rows.Next() {
// 		var row = make([]interface{}, 1)
// 		append(resp, rows.Scan)
// 	}

// 	if rows.Err() != nil {
// 		return nil, err
// 	}

// 	return nil, nil
// }

// GetNextRunNumber retrieves the current run number from the SQL database,
// returning -1 and an error if retrieval fails
func GetNextRunNumber(androidID string) (int, error) {
	// Retrieve database connection
	db := GetDatabase("data/test.sqlite")

	//Inserts the data in and prints out the run number
	carID, err := androidToCar(db, androidID)
	if err != nil {
		return -1, err
	}

	log.Infof("Car ID from DB: %s", carID)

	row := db.QueryRow("SELECT MAX(RunNumber)."+
		" FROM SensorData."+
		" WHERE Car id = ?", androidID)

	var carRunNum int
	if err = row.Scan(&carRunNum); err != nil {
		return -1, err
	}

	log.Infof("Current Run Number: %d", carRunNum)

	// Get the next car number
	var nextRunNumber = carRunNum + 1

	return nextRunNumber, nil
}
