package Services_GIS

import (
	. "Framework_Definitions"
	// . "Utilities"
	// "container/list"
	"fmt"
)

const ()

type Menu struct {
	Elements     map[string]MenuItem
	InputHandler MenuItem
	data         **[]dataRequest
}

type MenuItem struct {
	Label    string
	Response interface{}
}

// ------------------------------------------- Public ------------------------------------------- //

// Input handler functions should return an Event

// Define all menus
func InitMenus() Menu {
	var main Menu
	var changeData Menu
	var help Menu

	shared := &[]dataRequest{}

	main = Menu{
		Elements: map[string]MenuItem{
			"data": MenuItem{"Add/remove data", &changeData},
			"help": MenuItem{"Help", &help},
			"exit": MenuItem{"Exit", NewEvent(GLOBAL_EXIT, "", "")},
		},
		data: &shared,
	}

	changeData = Menu{
		Elements: map[string]MenuItem{
			"print": MenuItem{"Show all websites", (*Menu).displayData},
			"exit":  MenuItem{"Main Menu", &main},
		},
		InputHandler: MenuItem{"Enter url: ", (*Menu).changeDataInputHandler},
		data:         &shared,
	}

	help = Menu{
		Elements: map[string]MenuItem{
			"exit": MenuItem{"Main Menu", &main},
		},
		data: &shared,
	}

	return main
}

// ------------------------------------------- Change Data Menu ------------------------------------------- //

func (m *Menu) changeDataInputHandler(input string) Event {
	var returnEvent Event
	returnEvent.Type = GENERATE_MAP
	returnEvent.Parameter = dataRequest{
		datasetName: "city",
		requested:   map[string][]dataRange{"population": []dataRange{NewRange(10e6, -1)}},
	}
	return returnEvent
}

// ------------------------------------------- Help Menu ------------------------------------------- //

func (m *Menu) helpInputHandler(input string) Event {
	return NewEvent(NONE, "", "")
}

// ------------------------------------------- Utility ------------------------------------------- //

// func (m *Menu) appendSharedData(input interface{}) {
// 	var newArray []interface{}
// 	if *m.data == nil {
// 		newArray = []interface{}{input}
// 	} else {
// 		newArray = append(**m.data, input)
// 	}
// 	*m.data = &newArray
// }

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
