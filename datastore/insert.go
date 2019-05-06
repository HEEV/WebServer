package datastore

import (

	// Import sqlite3 library for use by database/sql
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"

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

// AndroidToCar This function uses the androidId to grab the carId
func AndroidToCar(db *sqlx.DB, androidId string) string {
	// Gets the database
	rows, err := db.Query("SELECT CarId."+
		" FROM CarTablet. "+
		" WHERE AndroidId = ?", androidId)

	// Grabs the android id and puts in and aid
	var carId string
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&carId)
		checkErr(err)
		log.Println(carId)
	}
	checkErr(err)
	err = rows.Err()
	checkErr(err)

	return carId
}

//This gets the current run number
func getNextRunNumber(androidId string) int {
	//gets the database
	db := GetDatabase("data/test.sqlite")

	//The commented line below does not work still need to figure out how to do it in go
	//var mysqli = new mysqli('localhost', getDatabaseUser(), getDatabasePassword(), getDatabaseServerName());

	//Inserts the data in and prints out the run number
	var carId = AndroidToCar(db, androidId)
	log.Println("server car id ", carId)

	rows, err := db.Query("SELECT MAX(RunNumber)."+
		" FROM SensorData."+
		" WHERE Car id = ?", androidId)
	checkErr(err)

	var carRunNum int
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&carRunNum)
		checkErr(err)
		log.Println(carId)
	}
	checkErr(err)
	//TODO: grab the carRunNum not carID
	//carRunNum = stmt.Exec(carId);
	log.Println("RunNum: ", carRunNum)

	//Get the next car number
	var nextRunNumber = carRunNum + 1
	db.Close()

	return nextRunNumber
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
