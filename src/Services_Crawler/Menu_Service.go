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
		returnEvent = s.respondToInput(event, sendChannel)
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
			// fmt.Println("Key found: ", key)
			switch item.(type) {
			case *Menu:
				// fmt.Println("Switching menus")
				fmt.Println(s.currentMenu)
				s.currentMenu = *(item.(*Menu))
				fmt.Println(s.currentMenu)
				return NewEvent(REQUEST_USER_INPUT, "", "")
			case func():
				item.(func())()
			case Event:
				sendChannel <- item.(Event)
			default:
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
