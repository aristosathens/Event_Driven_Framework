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
			"add":   MenuItem{"Add websites", &addWebsites},
			"print": MenuItem{"Show all websites", (*Menu).displayData},
			"exit":  MenuItem{"Exit", NewEvent(GLOBAL_EXIT, "", "")},
		},
		// InputHandler: nil,
		// data:         make([]string, 1),
	}

	addWebsites = Menu{
		// MenuType: ADD_WEBSITES,
		Elements: map[string]MenuItem{
			"print": MenuItem{"Show all websites", (*Menu).displayData},
			"exit":  MenuItem{"Main Menu", &main},
		},
		InputHandler: MenuItem{"Enter url: ", (*Menu).addWebsitesInputHandler},
	}

	return main
}

// ------------------------------------------- Public ------------------------------------------- //

// ------------------------------------------- Private ------------------------------------------- //

func (m *Menu) addWebsitesInputHandler(input string) {
	fmt.Println("Adding data to addWebsites menu: ", input)
	if IsValidUrl(input) {
		(*m).data = append(m.data, input)
		// This adds the website to the data in Main Menu
		// (*m).Elements["exit"].Response.data = append(((*m).Elements["exit"].Response).(*Menu).data.([]string), input)
	}
}

// ------------------------------------------- Utilities ------------------------------------------- //

func (m *Menu) displayMenu() {
	for key, item := range m.Elements {
		fmt.Println(" ( " + key + " ) " + item.Label)
	}
}

func (m *Menu) displayData() {
	fmt.Println("Menu data: ")
	for _, elem := range m.data {
		fmt.Println(elem)
	}
}
