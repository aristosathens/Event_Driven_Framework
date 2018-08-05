package ServiceInterface

import (
	. "container/list"
)

// ------------------------------------------- Type Definitions ------------------------------------------- //

// The use of EventType alias and the constants is like an enumerate type
type EventType int

// These are the EventTypes necessary in every application.
// Every EventType listed here must also be put in the array in the GetGlobalEventTypes() function
// To create custom event types, create a similar list in the Services.go file corresponding to your application. Use positive int values
const (
	NONE         EventType = -1
	GLOBAL_START EventType = -2
	GLOBAL_EXIT  EventType = -3
	KEY_DOWN     EventType = -4
	KEY_UP       EventType = -5
)

// Define the Event struct, which is used to pass messages between services
type Event struct {
	Type      EventType
	Parameter string
	Seen      bool
}

// Every service must implement the methods defined in the following interface
type Service interface {
	Init()
	Post(Event)
	Run()
}

// Every service must embed a ServiceFields struct
type ServiceFields struct {
	EventQueue List
	// EventTypesServiced []EventType
}

// ------------------------------------------- Utility Functions ------------------------------------------- //
