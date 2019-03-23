package packets

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
	RunNumber               int     //`json:"runNumber"`
	BatteryVoltage          float32 //`json:"batteryVoltage"`
	GroundSpeed             float32 //`json:"groundSpeed"`
	IntakeTemperature       float32 //`json:"intakeTemperature"`
	LKillSwitch             int     //`json:"lKillSwitch"`
	MKillSwitch             int     //`json:"mKillSwitch"`
	RKillSwitch             int     //`json:"rKillSwitch"`
	Longitude               float32 //`json:"longitude"`
	Latitude                float32 //`json:"latitude"`
	LogTime                 string  //`json:"logTime"`
	LapNumber               int     //`json:"lapNumber"`
	SecondaryBatteryVoltage float32 //`json:"secondaryBatteryVoltage"`
	WheelRpm                float32 //`json:"wheelRPM"`
	WindSpeed               float32 //`json:"windSpeed"`
	SystemCurrent           float32 //`json:"systemCurrent"`
	CoolantTemperature      float32 //`json:"coolantTemperature"`
}
