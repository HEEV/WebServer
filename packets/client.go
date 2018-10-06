package packets

// Identification represents the identification of a client. All packets include
// these fields
type Identification struct {
	MessageType string `json:"MessageType"`
	AndroidID   string `json:"AndroidId"`
}

// LogData represents an information update from client
type LogData struct {
	Ident                   Identification
	RunNumber               int
	BatteryVoltage          float32
	GroundSpeed             float32
	IntakeTemperature       float32
	LKillSwitch             int
	Latitude                float32
	LogTime                 string
	Longitude               float32
	LapNumber               int
	MKillSwitch             int
	RKillSwitch             int
	SecondaryBatteryVoltage float32
	WheelRpm                float32
	WindSpeed               float32
	SystemCurrent           float32
	CoolantTemperature      float32
}
