package Services

import (
	. "ServiceInterface"
	"container/list"
	"fmt"
)

// ------------------------------------------- Public ------------------------------------------- //

// Put custom event types here. Assign them to positive integers
const (
	UPDATE_GRAPHICS EventType = 1
)

// Returns array of all Services. All services MUST be put in this array
func AllServices() []Service {
	allServices := []Service{
		IOService{},
		TestService{},
	}
	return allServices
}

// ------------------------------------------- Services ------------------------------------------- //

// A service must be defined as a struct and implement the methods defined in Service{} in the ServiceInterface.go file
// IMPORTANT: All services must be included in the array in the AllServices() function

// ------------- IO Service ------------- //

type IOService struct {
	ServiceFields
}

func (s IOService) Init() {
}

func (s IOService) Post(event Event) {
	s.EventQueue.PushBack(event)
}

func (s IOService) Run() {
	for {
		// next and event are of type list Element
		// To get underlying Event struct we use: event.Value.(Event)
		var next *list.Element
		for event := s.EventQueue.Front(); event != nil; event = next {
			next = event.Next()
			s.EventQueue.Remove(event)

			switch eventType := event.Value.(Event).Type; eventType {
			case GLOBAL_EXIT:
				return
			case GLOBAL_START:
				//
			}
		}
	}
}

// ------------- Test Service ------------- //

type TestService struct {
	ServiceFields
}

func (s TestService) Init() {

}

func (s TestService) Post() {

}

func (s TestService) Run() {

}
