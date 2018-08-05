package ServiceInterface

import (
	"fmt"
	// . "container/list"
	// "unsafe"
	"reflect"
	"runtime"
)

// ------------------------------------------- Event Definitions ------------------------------------------- //

// The use of EventType alias and the constants is like an enumerate type
type EventType int

// These are the EventTypes necessary in every application.
// Every EventType listed here must also be put in the array in the GetGlobalEventTypes() function
// To create custom event types, create a similar list in the Services.go file corresponding to your application. Use positive int values
const (
	NONE         EventType = -1
	GLOBAL_START EventType = -2
	GLOBAL_EXIT  EventType = -3
	FINISHED     EventType = -4
	KEY_DOWN     EventType = -5
	KEY_UP       EventType = -6
)

// Define the Event struct, which is used to pass messages between services
type Event struct {
	Type      EventType
	Parameter string
	Origin    string
	Seen      bool
}

func NewEvent(eventType EventType, param string) Event {
	newEvent := Event{}
	newEvent.Type = eventType
	newEvent.Parameter = param
	newEvent.Origin = myCaller()
	newEvent.Seen = false
	return newEvent
}

// ------------------------------------------- Service Definitions ------------------------------------------- //

// Every service must embed a ServiceFields struct
type Service struct {
	Name           string
	Active         bool
	ReceiveChannel chan Event
	SendChannel    chan Event
	RunFunction    func(Event, *chan Event) Event
}

// Constructor for a Service struct
func NewService(receiveChannel chan Event, sendChannel chan Event, runFunction func(Event, *chan Event) Event) Service {
	newStruct := Service{}
	newStruct.Name = getFunctionName(runFunction)
	newStruct.Active = false
	newStruct.ReceiveChannel = receiveChannel
	newStruct.SendChannel = sendChannel
	newStruct.RunFunction = runFunction
	return newStruct
}

// Continuously checks the ReceiveChannel for events
func (s *Service) Run() {
	s.Active = true
	fmt.Println("In Service: ", s.Name)

	var event Event
	for {
		select {
		case event = <-s.ReceiveChannel:
			fmt.Println("event detected in service: ", event.Type)
			fmt.Println("event origin: ", event.Origin)
			fmt.Println("event detected in service: ", event.Type)

			returnEvent := s.RunFunction(event, &((*s).SendChannel))
			if returnEvent.Type == FINISHED {
				fmt.Println("Received FINISHED event")
				s.Close()
				return
			}
			if returnEvent.Type != NONE {
				fmt.Println("Received normal event")
				s.SendChannel <- returnEvent
			}
		default:
			// continue
		}

	}
}

// Closes the service's channels
func (s *Service) Close() {
	fmt.Println("Closing ", s.Name)
	close(s.ReceiveChannel)
	close(s.SendChannel)
	s.Active = false
}

// ------------------------------------------- Utility Functions ------------------------------------------- //

// https://stackoverflow.com/a/35213181/9463878
// MyCaller returns the caller of the function that called it. Use to set name of Service
func myCaller() string {

	fpcs := make([]uintptr, 1) // we get the callers as uintptrs - but we just need 1

	// skip 3 levels to get to the caller of whoever called Caller()
	n := runtime.Callers(3, fpcs)
	if n == 0 {
		return "n/a" // proper error her would be better
	}

	// get the info of the actual function that's in the pointer
	fun := runtime.FuncForPC(fpcs[0] - 1)
	if fun == nil {
		return "n/a"
	}

	return fun.Name() // return its name
}

func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
