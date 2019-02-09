package packets

// Identification represents the identification of a message.
// All packets include these fields
type Identification struct {
	MessageType string //`json:"messageType"`
	AndroidID   string //`json:"androidId"`
}

// ClientMessage represents the data expected from a client, all packets from a
// client must contain these fields
type ClientMessage struct {
	Data interface{} `json:"data,omitempty"` // Isn't currently used, but keeping for later
	Identification
}

// ServerMessage represents the packet that the server will be sending out
type ServerMessage struct {
	MessageID     string `json:"id"`
	ClientMessage        // All server messages will be reflecting a client message
}

// LogData represents an information update from client
type LogData struct {
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
	carId					int		//'json:"carId"`
}
