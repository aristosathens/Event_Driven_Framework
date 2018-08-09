package Services_Crawler

import (
	. "Framework_Definitions"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ------------- IO Service ------------- //

// ------------------------------------------- Public ------------------------------------------- //

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
	case KEY_DOWN:
		fmt.Println("Key stroked detected: ", event.Parameter)
	}
	fmt.Println("Exiting from IO Service Run Function")
	return returnEvent
}

// ------------------------------------------- Private ------------------------------------------- //

func (s *IOService) readUserInput(e Event, sendChannel chan Event) {
	fmt.Print("Enter text: ")
	text, _ := s.reader.ReadString('\n')
	text = strings.TrimSpace(text)
	newEvent := NewEvent(USER_INPUT, text, e.Origin)
	sendChannel <- newEvent
}
