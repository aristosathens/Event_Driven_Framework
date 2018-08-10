package Services_Crawler

import (
	// "encoding/json"
	. "Framework_Definitions"
	"Submenus_Crawler"
	// "JSON_Saver"
	"fmt"
	"strings"
	// "github.com/kardianos/osext"
	// "os"
	// "path/filepath"
)

const ()

type MenuType int

const (
	MAIN         MenuType = 0
	ADD_WEBSITES MenuType = 1
)

// ------------------------------------------- Public ------------------------------------------- //

type MenuService struct {
	currentMenu MenuType
	path        string
	menus       []Menu
	websites    []string
}

func (s *MenuService) Init() string {
	return "MenuService"
}

func (s *MenuService) RunFunction(event Event, sendChannel chan Event) Event {
	returnEvent := NewEvent(NONE, "", "")

	switch event.Type {

	case GLOBAL_START:
		s.currentMenu = MAIN
		sendChannel <- NewEvent(REQUEST_WEBSITE_LIST, "", "")
		return NewEvent(REQUEST_USER_INPUT, "Enter command: ", "")

	case GLOBAL_EXIT:
		returnEvent.Type = FINISHED

	case USER_INPUT:
		s.respondToInput(event, sendChannel)

	case WEBSITE_LIST:
		s.websites = event.Parameter.([]string)
	}

	s.updateDisplay()
	return returnEvent
}

// ------------------------------------------- Private ------------------------------------------- //

// Takes USER_INPUT event and updates service according to message and current status
func (s *MenuService) respondToInput(event Event, sendChannel chan Event) Event {
	msg := strings.ToLower(event.Parameter.(string))

	switch s.currentMenu {

	case MAIN:
		if msg == "a" || msg == "add" {
			s.currentMenu = ADD_WEBSITES
			sendChannel <- NewEvent(REQUEST_USER_INPUT, "Enter urls you would like to add: ", "")
			break

			// return NewEvent(UPDATE_MENU, "", "self")
		} else if msg == "c" || msg == "crawl" {
			// TO DO: Implement crawling
		} else if msg == "e" || msg == "exit" {
			fmt.Println("Posting global exit event")
			sendChannel <- NewEvent(GLOBAL_EXIT, "", "")
		} else {
			fmt.Println("Invalid input.")

		}

	case ADD_WEBSITES:
		if isValidUrl(msg) {
			fmt.Println("Added " + msg + " to websites list.")
			sendChannel <- NewEvent(ADD_WEBSITE, msg, "")
			sendChannel <- NewEvent(REQUEST_USER_INPUT, "Enter urls you would like to add: ", "")
			sendChannel <- NewEvent(REQUEST_WEBSITE_LIST, "", "")

		} else if msg == "exit" {
			s.currentMenu = MAIN
			break

		} else {
			fmt.Println("Not a valid url.")
			sendChannel <- NewEvent(REQUEST_USER_INPUT, "Enter urls you would like to add: ", "")
		}
	}

	return NewEvent(NONE, "", "")
}

// Prints menu to the console
func (s *MenuService) updateDisplay() {
	switch s.currentMenu {

	case MAIN:
		fmt.Println("(add) Add new websites")
		fmt.Println("(exit) Exit")

	case ADD_WEBSITES:
		fmt.Println("(exit) Back")
		fmt.Println("Current websites:")
		for _, site := range s.websites {
			fmt.Println("- " + site)
		}

	}

}

// func (s *MenuService) exit() {

// }

// func (s *MenuService) defineMenus() {
// 	s.menus = []Menu{
// 		Menu{
// 			menuType: MAIN,
// 			elements: []MenuItem{
// 				MenuItem{"Exit: ", "exit"},
// 				MenuItem{"Add websites", "add"},
// 			},
// 		},
// 		Menu{
// 			menuType: ADD_WEBSITES,
// 			elements: []MenuItem{
// 				MenuItem{"Enter new websites names: ", ""},
// 			},
// 		},
// 	}
// }
