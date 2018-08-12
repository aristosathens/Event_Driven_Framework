package Services_Crawler

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

// type MenuType int

// const (
// 	MAIN         MenuType = 0
// 	ADD_WEBSITES MenuType = 1
// )

// ------------------------------------------- Public ------------------------------------------- //

type MenuService struct {
	currentMenu Menu
	path        string
	// menus       []Menu
	websites []string
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
		returnEvent.Type = FINISHED

	case USER_INPUT:
		sendChannel <- NewEvent(REQUEST_USER_INPUT, "", "")
		returnEvent = s.respondToInput(event, sendChannel)

	case REQUEST_USER_INPUT:
		// ignore this event type
		return returnEvent
	}

	s.currentMenu.displayMenu()
	// sendChannel <- NewEvent(REQUEST_USER_INPUT, "", "")
	return returnEvent
}

// ------------------------------------------- Private ------------------------------------------- //

func (s *MenuService) respondToInput(event Event, sendChannel chan Event) Event {

	msg := strings.ToLower(event.Parameter.(string))

	for key, elem := range s.currentMenu.Elements {
		item := elem.Response
		if msg == key || msg == string(key[0]) {
			// fmt.Println("Key found: ", key)
			switch item.(type) {
			case *Menu:
				// fmt.Println("Switching menus")
				s.currentMenu = *(item.(*Menu))
			case func(*Menu):
				fmt.Println(s.currentMenu)
				item.(func(*Menu))(&s.currentMenu)
			case Event:
				sendChannel <- item.(Event)
			default:
				fmt.Println("Default case - Shouldn't be here !")
				// break
			}
			return NewEvent(NONE, "", "")
		}
	}

	if s.currentMenu.InputHandler.Response != nil {
		s.currentMenu.InputHandler.Response.(func(*Menu, string))(&s.currentMenu, msg)
	}
	return NewEvent(NONE, "", "")
}
