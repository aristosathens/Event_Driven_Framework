package Services_Crawler

import (
	. "Framework_Definitions"
)

// ------------------------------------------- Names ------------------------------------------- //

// To create new services, in a new file create a struct and implement the methods found in ServiceInterface in ServiceInterface.go
// All service structs must be in the AllServiceInterfaces array at the top of Services.go

// Put custom event types here. Assign them to positive integers
const (
	PING               EventType = 1
	PONG               EventType = 2
	REQUEST_USER_INPUT EventType = 3
	USER_INPUT         EventType = 4
)

var AllServiceInterfaces = [...]ServiceInterface{
	&(IOService{}),
	&(CrawlerService{}),
}

// ------------------------------------------- Services ------------------------------------------- //

// ------------- Test Service ------------- //

// type TestService struct {
// }

// func (s *TestService) Init() {
// }

// func (s *TestService) RunFunction(event Event, sendChannel chan Event) Event {
// 	returnEvent := NewEvent(NONE, "", "")
// 	switch eventType := event.Type; eventType {
// 	case GLOBAL_EXIT:
// 		returnEvent.Type = FINISHED
// 	case PING:
// 		fmt.Println(fmt.Sprintf("TestService received %s %d event from %s ", event.Parameter, event.Type, event.Origin))
// 	case KEY_DOWN:
// 		if event.Parameter == "exit" {
// 			fmt.Println("Detected exit key stroke")
// 			returnEvent.Type = GLOBAL_EXIT
// 		}
// 	}
// 	return returnEvent
// }
