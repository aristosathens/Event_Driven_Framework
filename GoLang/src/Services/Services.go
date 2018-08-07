package Services

import (
	. "ServiceInterface"
	"bufio"
	"fmt"
	"os"
	"strings"
	// "unsafe"
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
}

// ------------------------------------------- Services ------------------------------------------- //

// To create new services, implement ServiceNameRun() functions.
// All ServiceNameRun() functions must be included in the array in the AllServiceNames() function

// ------------- IO Service ------------- //
type IOService struct {
	// Service specific variables here
	exampleLocal int
}

func (s *IOService) Init() {
	// s.newLocal = 5
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
		go inputMonitor(sendChannel)
	case KEY_DOWN:
		fmt.Println("Key stroked detected: ", event.Parameter)
	}
	fmt.Println("returning from IO Service")
	return returnEvent
}

func inputMonitor(sendChannel chan Event) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	// if strings.Compare(text, "exit") == 0 {
	// 	*sendChannel <- NewEvent(GLOBAL_EXIT, "")
	// 	fmt.Println("comparison passed")
	// 	return
	// }
	sendChannel <- NewEvent(KEY_DOWN, text)
}

// ------------- Test Service ------------- //

func TestService(event Event, sendChannel chan Event) Event {
	returnEvent := NewEvent(NONE, "")

	fmt.Println("In Service run function")
	switch eventType := event.Type; eventType {
	case GLOBAL_EXIT:
		returnEvent.Type = FINISHED
	case PING:
		fmt.Println(fmt.Sprintf("TestService received %s %d event from %s ", event.Parameter, event.Type, event.Origin))
	case KEY_DOWN:
		if event.Parameter == "exit" {
			returnEvent.Type = GLOBAL_EXIT
		}
	default:
		//
	}
	return returnEvent
}
