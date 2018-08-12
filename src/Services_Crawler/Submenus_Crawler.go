package Services_Crawler

import (
	. "Framework_Definitions"
	. "Utilities"
	// "container/list"
	"fmt"
)

// ------------------------------------------- Public ------------------------------------------- //

const ()

type Menu struct {
	// MenuType MenuType
	Elements     map[string]MenuItem
	InputHandler MenuItem
	data         **[]interface{}
}

type MenuItem struct {
	Label    string
	Response interface{}
}

// Define all menus
func InitMenus() Menu {
	var main Menu
	var addWebsites Menu
	var crawler Menu

	var shared *[]interface{}

	main = Menu{
		Elements: map[string]MenuItem{
			"add":   MenuItem{"Add websites", &addWebsites},
			"crawl": MenuItem{"Crawl websites", &crawler},
			"print": MenuItem{"Show all websites", (*Menu).displayData},
			"exit":  MenuItem{"Exit", NewEvent(GLOBAL_EXIT, "", "")},
		},
		data: &shared,
	}

	addWebsites = Menu{
		Elements: map[string]MenuItem{
			"print": MenuItem{"Show all websites", (*Menu).displayData},
			"exit":  MenuItem{"Main Menu", &main},
		},
		InputHandler: MenuItem{"Enter url: ", (*Menu).addWebsitesInputHandler},
		data:         &shared,
	}

	crawler = Menu{
		Elements: map[string]MenuItem{
			"exit": MenuItem{"Main Menu", &main},
		},
		InputHandler: MenuItem{"Choose which website to crawl. Type 'all' to crawl all websites", (*Menu).crawlerInputHandler},
		data:         &shared,
	}

	return main
}

// ------------------------------------------- Private ------------------------------------------- //

func (m *Menu) addWebsitesInputHandler(input string) {
	fmt.Println("Adding % to addWebsites menu.", input)
	if IsValidUrl(input) {
		m.appendSharedData(input)
		// This adds the website to the data in Main Menu
		// (*m).Elements["exit"].Response.data = append(((*m).Elements["exit"].Response).(*Menu).data.([]string), input)
	}
}

func (m *Menu) crawlerInputHandler(input string) {
	fmt.Println("Crawling: ", input)

}

// ------------------------------------------- Utilities ------------------------------------------- //

func (m *Menu) appendSharedData(input interface{}) {
	var newArray []interface{}
	if *m.data == nil {
		newArray = []interface{}{input}
	} else {
		newArray = append(**m.data, input)
	}
	*m.data = &newArray
}

func (m *Menu) displayMenu() {
	for key, item := range m.Elements {
		fmt.Println(" ( " + key + " ) " + item.Label)
	}
}

func (m *Menu) displayData() {
	fmt.Println("Menu data: ")
	for _, elem := range **m.data {
		fmt.Println(elem)
	}
}
