package Services_Crawler

import (
	. "Framework_Definitions"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ------------------------------------------- Public ------------------------------------------- //

// This service handles input and output.
// To request input from the user, send an event with type REQUEST_USER_INPUT and your message in the parameter
// To print output, send an event with type PRINT_LINE, and your message in the parameter

type IOService struct {
	reader       *bufio.Reader
	exampleLocal int
}

func (s *IOService) Init() string {
	s.reader = bufio.NewReader(os.Stdin)
	return "IOService"
}

func (s *IOService) RunFunction(event Event, sendChannel chan Event) Event {
	fmt.Println("Entering IO service Run Function")

	returnEvent := NewEvent(NONE, "", "")

	switch eventType := event.Type; eventType {
	case GLOBAL_START:
	case GLOBAL_EXIT:
		returnEvent.Type = FINISHED

	case REQUEST_USER_INPUT:
		go s.readUserInput(event, sendChannel)

	}
	fmt.Println("Exiting from IO Service Run Function")
	return returnEvent
}

// ------------------------------------------- Private ------------------------------------------- //

func (s *IOService) readUserInput(event Event, sendChannel chan Event) {
	fmt.Println("Type 'exit' to quit application.")
	fmt.Println(event.Parameter)
	text, _ := s.reader.ReadString('\n')
	text = strings.TrimSpace(text)
	var newEvent Event
	if text == "exit" {
		newEvent = NewEvent(GLOBAL_EXIT, "", "")
	} else {
		newEvent = NewEvent(USER_INPUT, text, event.Origin)
	}
	sendChannel <- newEvent
}
