package packets

import "fmt"

// TabletInitial represents the initial communication with the server
type TabletInitial struct {
	Identification
}

// Identification represents the message identification information
// The type of the message and the device ID
type Identification struct {
	MessageType string
	AndroidID   string
}

// LogData represents an information update from client
type LogData struct {
	Identification
	SensorData
}

// SensorData represents raw sensor data
type SensorData struct {
	BatteryVoltage          float32 `json:"batteryVoltage" db:"batteryVoltage"`
	SecondaryBatteryVoltage float32 `json:"secondaryBatteryVoltage" db:"secondaryBatteryVoltage"`
	CarId                   string  `json:"carId" db:"carId"`
	CoolantTemperature      float32 `json:"coolantTemperature" db:"coolantTemperature"`
	GroundSpeed             float32 `json:"groundSpeed" db:"groundSpeed"`
	IntakeTemperature       float32 `json:"intakeTemperature" db:"intakeTemperature"`
	Latitude                float32 `json:"latitude" db:"latitude"`
	Longitude               float32 `json:"longitude" db:"longitude"`
	LapNumber               int     `json:"lapNumber" db:"lapNumber"`
	LogTime                 string  `json:"logTime" db:"logTime"`
	LKillSwitch             int     `json:"lKillSwitch" db:"lKillSwitch"`
	MKillSwitch             int     `json:"mKillSwitch" db:"mKillSwitch"`
	RKillSwitch             int     `json:"rKillSwitch" db:"rKillSwitch"`
	RunNumber               int     `json:"runNumber" db:"runNumber"`
	SystemCurrent           float32 `json:"systemCurrent" db:"systemCurrent"`
	WheelRpm                float32 `json:"wheelRPM" db:"wheelRpm"`
	WindSpeed               float32 `json:"windSpeed" db:"windSpeed"`
}

// DBSensorData represents the data from the SensorData table in the DB
type DBSensorData struct {
	SensorData
	AndroidID string `json:"androidId" db:"androidId"`
	ID        int    `json:"-" db:"id"`
}

func (s DBSensorData) ToCSVString() []string {
	return []string{
		s.AndroidID,
		fmt.Sprintf("%f", s.BatteryVoltage),
		fmt.Sprintf("%f", s.SecondaryBatteryVoltage),
		s.CarId,
		fmt.Sprintf("%f", s.CoolantTemperature),
		fmt.Sprintf("%f", s.GroundSpeed),
		fmt.Sprintf("%f", s.IntakeTemperature),
		fmt.Sprintf("%f", s.Latitude),
		fmt.Sprintf("%f", s.Longitude),
		fmt.Sprintf("%d", s.LapNumber),
		s.LogTime,
		fmt.Sprintf("%d", s.LKillSwitch),
		fmt.Sprintf("%d", s.RKillSwitch),
		fmt.Sprintf("%d", s.MKillSwitch),
		fmt.Sprintf("%d", s.RunNumber),
		fmt.Sprintf("%f", s.SystemCurrent),
		fmt.Sprintf("%f", s.WheelRpm),
		fmt.Sprintf("%f", s.WindSpeed),
	}
}
