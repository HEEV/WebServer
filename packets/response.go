package packets

// NextRunNumberData represents the data passed in response to the
// GetNextRunNumber MessageType
type NextRunNumberData struct {
	RunNumber int `json:"runNumber"`
}
