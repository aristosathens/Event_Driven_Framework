package Services

import (
	. "ServiceInterface"
	"bufio"
	"fmt"
	"os"
)

// ------------------------------------------- Service Names ------------------------------------------- //

// Put custom event types here. Assign them to positive integers
const (
	PING EventType = 1
	PONG EventType = 2
)

// Returns array of all Services. All services MUST be put in this array
func AllServiceNames() []func(event Event, sendChannel *chan Event) Event {
	return []func(event Event, sendChannel *chan Event) Event{
		IOService,
		// TestService,
	}
}

// ------------------------------------------- Services ------------------------------------------- //

// To create new services, implement ServiceNameRun() functions.
// All ServiceNameRun() functions must be included in the array in the AllServiceNames() function

// ------------- IO Service ------------- //

func IOService(event Event, sendChannel *chan Event) Event {
	fmt.Println("IO service got event: ", event.Type)
	returnEvent := NewEvent(NONE, "")

	switch eventType := event.Type; eventType {
	case GLOBAL_EXIT:
		returnEvent.Type = FINISHED
	case GLOBAL_START:
		fmt.Println("IOService sees GLOBAL_START event")
		go inputMonitor(sendChannel)
	case KEY_DOWN:
		fmt.Println("Key stroked detected: ", event.Parameter)
	}
	return returnEvent
}

func inputMonitor(sendChannel *chan Event) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		*sendChannel <- NewEvent(KEY_DOWN, text)
		// fmt.Println(text)
	}
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
	default:
		//
	}
	return returnEvent
}
