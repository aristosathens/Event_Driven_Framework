package Services_GIS

import (
	// "encoding/json"
	. "Framework_Definitions"
	// "JSON_Saver"
	"fmt"
	"strings"
	// "github.com/kardianos/osext"
	// "os"
	// "path/filepath"
)

const ()

// ------------------------------------------- Public ------------------------------------------- //

type MenuService struct {
	currentMenu Menu
	path        string
	websites    []string
}

func (s *MenuService) Init() string {
	return "MenuService"
}

func (s *MenuService) RunFunction(event Event, sendChannel chan Event) Event {

	returnEvent := NewEvent(NONE, "", "")

	switch event.Type {

	case GLOBAL_START:
		s.currentMenu = InitMenus()
		returnEvent.Type = REQUEST_USER_INPUT

	case GLOBAL_EXIT:
		return NewEvent(FINISHED, "", "")

	case USER_INPUT:
		sendChannel <- s.respondToInput(event, sendChannel)
		returnEvent = NewEvent(REQUEST_USER_INPUT, s.currentMenu.InputHandler.Label, "")

	case IS_VALID_DATA_REQUEST:
		if event.Parameter == true {
			for _, request := range **(s.currentMenu.tempData) {
				**(s.currentMenu.requestData) = mergeRequests(**(s.currentMenu.requestData), request)
			}
		}
		s.currentMenu.clearTempData()
		return returnEvent

	default:
		return returnEvent
	}

	s.currentMenu.displayMenu()
	return returnEvent
}

// ------------------------------------------- Private ------------------------------------------- //

func (s *MenuService) respondToInput(event Event, sendChannel chan Event) Event {

	msg := strings.ToLower(event.Parameter.(string))

	for key, elem := range s.currentMenu.Elements {
		item := elem.Response
		if msg == key || msg == string(key[0]) {
			switch item.(type) {
			case *Menu:
				s.currentMenu = *(item.(*Menu))
			case func(*Menu):
				item.(func(*Menu))(&s.currentMenu)
			case Event:
				sendChannel <- item.(Event)
			default:
				fmt.Println("Default case - Shouldn't be here !")
			}
			return NewEvent(NONE, "", "")
		}
	}

	if s.currentMenu.InputHandler.Response != nil {
		return s.currentMenu.InputHandler.Response.(func(*Menu, string) Event)(&s.currentMenu, msg)
	}
	return NewEvent(NONE, "", "")
}
