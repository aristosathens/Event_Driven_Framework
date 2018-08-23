package Services_GIS

import (
	. "Framework_Definitions"
)

// ------------------------------------------- Definitions ------------------------------------------- //

// To create new services, in a new file create a struct and implement the methods found in ServiceInterface in ServiceInterface.go
// All service structs must be in the AllServiceInterfaces array at the top of Services.go

// Put custom event types here. Assign them to positive integers
const (
	PING                  EventType = 1
	PONG                  EventType = 2
	PRINT_LINE            EventType = 5
	REQUEST_USER_INPUT    EventType = 10
	USER_INPUT            EventType = 11
	ADD_DATA              EventType = 20
	REMOVE_DATA           EventType = 21
	GENERATE_MAP          EventType = 22
	REQUEST_DATASET_NAMES EventType = 30
	DATASET_CHANGE        EventType = 31
	CHECK_DATA_REQUEST    EventType = 40
	IS_VALID_DATA_REQUEST EventType = 41

	// REQUEST_DATA       EventType = 30
)

var AllServiceInterfaces = [...]ServiceInterface{
	&(IOService{}),
	&(GISService{}),
	&(MenuService{}),
}
