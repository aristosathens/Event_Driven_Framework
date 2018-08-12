package Services_Crawler

import (
	. "Framework_Definitions"
)

// ------------------------------------------- Names ------------------------------------------- //

// To create new services, in a new file create a struct and implement the methods found in ServiceInterface in ServiceInterface.go
// All service structs must be in the AllServiceInterfaces array at the top of Services.go

// Put custom event types here. Assign them to positive integers
const (
	PING                 EventType = 1
	PONG                 EventType = 2
	PRINT_LINE           EventType = 5
	REQUEST_USER_INPUT   EventType = 10
	USER_INPUT           EventType = 11
	ADD_WEBSITE          EventType = 20
	REMOVE_WEBSITE       EventType = 21
	REQUEST_WEBSITE_LIST EventType = 30
	WEBSITE_LIST         EventType = 31
	// UPDATE_MENU        EventType = 10
)

var AllServiceInterfaces = [...]ServiceInterface{
	&(IOService{}),
	&(CrawlerService{}),
	&(MenuService{}),
}
