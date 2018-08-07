package Services

import (
	. "ServiceInterface"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ------------------------------------------- Service Names ------------------------------------------- //

// Put custom event types here. Assign them to positive integers
const (
	PING           EventType = 1
	PONG           EventType = 2
	GET_USER_INPUT EventType = 3
)

var AllServiceInterfaces = [...]ServiceInterface{
	&(IOService{}),
	&(TestService{}),
}

// ------------------------------------------- Services ------------------------------------------- //

// To create new services, create a struct and implement the methods found in ServiceInterface in ServiceInterface.go
// All service structs must be in the AllServiceInterfaces array at the top of Services.go

// ------------- IO Service ------------- //
type IOService struct {
	exampleLocal int
}

func (s *IOService) Init() {
	s.exampleLocal = 5
}

func (s *IOService) RunFunction(event Event, sendChannel chan Event) Event {
	fmt.Println("IO service got event: ", event.Type)
	returnEvent := NewEvent(NONE, "")

	switch eventType := event.Type; eventType {
	case GLOBAL_EXIT:
		returnEvent.Type = FINISHED
	case GLOBAL_START:
		fmt.Println("IOService sees GLOBAL_START event")
		returnEvent.Type = GET_USER_INPUT
	case GET_USER_INPUT:
		go s.inputMonitor(sendChannel)
	case KEY_DOWN:
		fmt.Println("Key stroked detected: ", event.Parameter)
	}
	fmt.Println("returning from IO Service")
	return returnEvent
}

func (s *IOService) inputMonitor(sendChannel chan Event) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	sendChannel <- NewEvent(KEY_DOWN, text)
}

// ------------- Test Service ------------- //

type TestService struct {
}

func (s *TestService) Init() {
}

func (s *TestService) RunFunction(event Event, sendChannel chan Event) Event {
	returnEvent := NewEvent(NONE, "")
	switch eventType := event.Type; eventType {
	case GLOBAL_EXIT:
		returnEvent.Type = FINISHED
	case PING:
		fmt.Println(fmt.Sprintf("TestService received %s %d event from %s ", event.Parameter, event.Type, event.Origin))
	case KEY_DOWN:
		if event.Parameter == "exit" {
			fmt.Println("Detected exit key stroke")
			returnEvent.Type = GLOBAL_EXIT
		}
	}
	return returnEvent
}
