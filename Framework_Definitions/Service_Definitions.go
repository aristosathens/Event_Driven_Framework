package Framework_Definitions

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
)

// ------------------------------------------- Service Definitions ------------------------------------------- //

const (
	BufferSize = 200
)

// All services must be a struct with implementations of the methods defined in this interface
type ServiceInterface interface {
	Init() string
	RunFunction(Event, chan Event) Event
}

// Every user defined ServiceStruct (in Services.go) will be embedded here in a Service struct as Locals
type Service struct {
	Name           string
	Active         bool
	ReceiveChannel chan Event
	SendChannel    chan Event
	RunFunction    func(Event, chan Event) Event
	Locals         ServiceInterface
}

// // Constructor for a Service struct
func NewService(serviceStruct ServiceInterface) Service {
	newStruct := Service{}
	newStruct.Active = false

	newStruct.Locals = serviceStruct
	newStruct.Name = newStruct.Locals.Init()
	newStruct.RunFunction = serviceStruct.RunFunction

	newStruct.ReceiveChannel = make(chan Event, BufferSize)
	newStruct.SendChannel = make(chan Event, BufferSize)
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
			returnEvent := s.RunFunction(event, (*s).SendChannel)
			if returnEvent.Type == FINISHED {
				s.Close()
				return
			}
			if returnEvent.Type != NONE {
				s.SendChannel <- returnEvent
			}
		default:
			//
		}

	}
}

// Closes the sending channel and sets status to inactive
func (s *Service) Close() {
	fmt.Println("Closing ", s.Name)
	close(s.SendChannel)
	s.Active = false
}

// ------------------------------------------- Utility Functions ------------------------------------------- //

// Modified version of this SO answer: https://stackoverflow.com/a/35213181/9463878
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

	name := fun.Name()
	start := strings.Index(name, "*") + 1
	end := strings.Index(name, ")")

	// fmt.Println("Name is: ", name)
	// fmt.Println(start)
	// fmt.Println(end)

	if start < 0 || end <= 0 {
		return ""
	} else {
		newName := name[start:end]
		return newName

	}
}

// Given a function, returns its name
func getFunctionName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}
