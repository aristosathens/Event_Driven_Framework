package Services_Crawler

import (
	. "Framework_Definitions"
	. "Utilities"
	// "container/list"
	"fmt"
)

// ------------------------------------------- Definitions ------------------------------------------- //

type MenuType int

const (
	MAIN         MenuType = 1
	ADD_WEBSITES MenuType = 2
)

type Menu struct {
	// MenuType MenuType
	Elements     map[string]MenuItem
	InputHandler MenuItem
	data         []interface{}
}

type MenuItem struct {
	Label    string
	Response interface{}
}

// Define all menus
func InitMenus() Menu {
	var addWebsites Menu
	var main Menu

	main = Menu{
		// MenuType: MAIN,
		Elements: map[string]MenuItem{
			"exit": MenuItem{"Exit", NewEvent(GLOBAL_EXIT, "", "")},
			"add":  MenuItem{"Add websites", &addWebsites},
		},
	}

	addWebsites = Menu{
		// MenuType: ADD_WEBSITES,
		Elements: map[string]MenuItem{
			"exit": MenuItem{"Main Menu", &main},
		},
		InputHandler: MenuItem{"Enter url: ", (*Menu).addWebsitesInputHandler},
	}

	return main
}

// ------------------------------------------- Public ------------------------------------------- //

// ------------------------------------------- Private ------------------------------------------- //

func (m *Menu) addWebsitesInputHandler(input string) {
	if IsValidUrl(input) {
		m.data = append(m.data, input)
	}
}

// ------------------------------------------- Utilities ------------------------------------------- //

func (m *Menu) displayMenu() {
	for key, item := range m.Elements {
		fmt.Println(" ( " + key + " ) " + item.Label)
	}
}
