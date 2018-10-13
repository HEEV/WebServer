package packets

// Identification represents the identification of a message.
// All packets include these fields
type Identification struct {
	MessageType string      `json:"messageType"`
	AndroidID   string      `json:"androidId"`
	MessageID   string      `json:"id"`
	Data        interface{} `json:"data"` // Isn't currently used, but keeping for later
}

// LogData represents an information update from client
type LogData struct {
	Ident                   Identification `json:"ident"`
	RunNumber               int            `json:"runNumber"`
	BatteryVoltage          float32        `json:"batteryVoltage"`
	GroundSpeed             float32        `json:"groundSpeed"`
	IntakeTemperature       float32        `json:"intakeTemperature"`
	LKillSwitch             int            `json:"lKillSwitch"`
	Latitude                float32        `json:"latitude"`
	LogTime                 string         `json:"logTime"`
	Longitude               float32        `json:"longitude"`
	LapNumber               int            `json:"lapNumber"`
	MKillSwitch             int            `json:"mKillSwitch"`
	RKillSwitch             int            `json:"rKillSwitch"`
	SecondaryBatteryVoltage float32        `json:"secondaryBatteryVoltage"`
	WheelRpm                float32        `json:"wheelRPM"`
	WindSpeed               float32        `json:"windSpeed"`
	SystemCurrent           float32        `json:"systemCurrent"`
	CoolantTemperature      float32        `json:"coolantTemperature"`
}
