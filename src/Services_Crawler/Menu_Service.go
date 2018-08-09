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

type MenuType int

const (
	MAIN         MenuType = 0
	ADD_WEBSITES MenuType = 1
)

type Menu struct {
	menuType MenuType
	elements []MenuItem
}

type MenuItem struct {
	label string
	key   string
	// text  func()
}

// ------------------------------------------- Public ------------------------------------------- //

type MenuService struct {
	currentMenu MenuType
	path        string
	menus       []Menu
}

func (s *MenuService) Init() string {
	s.defineMenus()
	return "MenuService"
}

func (s *MenuService) RunFunction(event Event, sendChannel chan Event) Event {
	returnEvent := NewEvent(NONE, "", "")

	switch event.Type {

	case GLOBAL_START:
		s.currentMenu = MAIN
		return NewEvent(REQUEST_USER_INPUT, "Enter command: ", "")

	case GLOBAL_EXIT:
		returnEvent.Type = FINISHED
	}

	return returnEvent
}

// ------------------------------------------- Private ------------------------------------------- //

// Takes USER_INPUT event and updates service according to message and current status
func (s *MenuService) respondToInput(event Event, sendChannel chan Event) Event {
	msg := strings.ToLower(event.Parameter)

	switch s.currentMenu {

	case MAIN:
		if msg == "a" || msg == "add" {
			s.currentMenu = ADD_WEBSITES

			// return NewEvent(UPDATE_MENU, "", "self")
		} else if msg == "c" || msg == "crawl" {
			// TO DO: Implement crawling
		} else {
			fmt.Println("Invalid input.")

		}

	case ADD_WEBSITES:
		if isValidUrl(msg) {
			fmt.Println("Added " + msg + " to websites list.")
			sendChannel <- NewEvent(ADD_WEBSITE, msg, "")
			sendChannel <- NewEvent(REQUEST_USER_INPUT, "Enter urls you would like to add: ", "")

		} else if msg == "\033" {

		} else {
			fmt.Println("Not a valid url.")
			sendChannel <- NewEvent(REQUEST_USER_INPUT, "Enter urls you would like to add: ", "")
		}
	}

	return []Event{NewEvent(NONE, "", "")}
}

// Prints menu to the console
func (s *MenuService) displayMenu() {
	switch s.currentMenu {
	case MAIN:

	case ADD_WEBSITES:
	}

}

func (s *MenuService) exit() {

}

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
