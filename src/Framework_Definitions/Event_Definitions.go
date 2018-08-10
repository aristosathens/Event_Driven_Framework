package Framework_Definitions

import (
	"strings"
)

// ------------------------------------------- Event Definitions ------------------------------------------- //

// The use of EventType alias and the constants is like an enumerate type
type EventType int

// These are the EventTypes necessary in every application.
// To create custom event types, create a similar list in the Services.go file corresponding to your application. Use positive int values
const (
	NONE         EventType = -1
	GLOBAL_START EventType = -2
	GLOBAL_EXIT  EventType = -3
	FINISHED     EventType = -4
)

// The use of EventType alias and the constants is like an enumerate type
type ParameterType int

const (
	EMPTY ParameterType = 0
)

// Define the Event struct, which is used to pass messages between services
type Event struct {
	Type      EventType
	Parameter interface{}
	Origin    string
	Target    string
}

// Event constructor
func NewEvent(eventType EventType, param interface{}, target string) Event {
	newEvent := Event{}
	newEvent.Type = eventType
	newEvent.Parameter = param
	newEvent.Origin = myCaller()
	newEvent.Target = target
	if strings.ToLower(target) == "self" {
		newEvent.Target = myCaller()
	}
	return newEvent
}
