package packets

// ResponseBase represents the basic information every response packet will
// contain
type ResponseBase struct {
	AndroidID string
	Data      interface{}
}

// NextRunNumberData represents the data passed in response to the
// GetNextRunNumber MessageType
type NextRunNumberData struct {
	RunNumber int
}
