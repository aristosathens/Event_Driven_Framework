package Services_GIS

import (
	. "Framework_Definitions"
	"strconv"
	// . "Utilities"
	// "container/list"
	"errors"
	"fmt"
	"strings"
)

const ()

type Menu struct {
	Name         string
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
		Name: "main",
		Elements: map[string]MenuItem{
			"data": MenuItem{"Change data", &changeData},
			"help": MenuItem{"Help", &help},
			"exit": MenuItem{"Exit", NewEvent(GLOBAL_EXIT, "", "")},
		},
		data: &shared,
	}

	changeData = Menu{
		Name: "changeData",
		Elements: map[string]MenuItem{
			"exit": MenuItem{"Main Menu", &main},
			"help": MenuItem{"Help", (*Menu).changeDataDisplayHelp},
		},
		InputHandler: MenuItem{"What data would you like to add?", (*Menu).changeDataInputHandler},
		data:         &shared,
	}

	help = Menu{
		Name: "help",
		Elements: map[string]MenuItem{
			"exit": MenuItem{"Main Menu", &main},
		},
		data: &shared,
	}

	return main
}

// ------------------------------------------- Change Data Menu ------------------------------------------- //

// Create empty dataRequest, then fill it with parsed user input. Return it as an ADD_DATA event
func (m *Menu) changeDataInputHandler(input string) Event {
	request := dataRequest{"", map[string][]dataRange{}}
	request, err := parseDataRequest(input, request)
	checkError(err)
	returnEvent := NewEvent(ADD_DATA, request, "")
	return returnEvent
}

// Parses one line of user input. Exepects a single data type and one or more ranges..
// Takes in a preexisint request and fills out a single map key/element in the dataRequest.requested
func parseDataRequest(input string, request dataRequest) (dataRequest, error) {
	input = trim(input)
	i := strings.Index(input, "[")
	if i < 0 {
		return request, errors.New("Improperly formatted input.")
	}
	dataType := trim(input[:i])
	for {
		r := dataRange{}

		iL := strings.Index(input, "[")
		iR := strings.Index(input, ",")
		if iL < 0 || iR < 0 || iL > iR {
			return request, errors.New("Improperly formatted input.")
		}
		var err error
		r.lower, err = strconv.ParseFloat(trim(input[iL+1:iR]), 64)
		if err != nil {
			return request, errors.New("Range must be composed of numbers.")
		}

		iL = iR
		iR = strings.Index(input, "]")
		if iR < 0 || iL > iR {
			return request, errors.New("Improperly formatted input.")
		}
		r.upper, err = strconv.ParseFloat(trim(input[iL+1:iR]), 64)
		if err != nil {
			return request, errors.New("Range must be composed of numbers.")
		}

		fmt.Println("Attempting to write to requested map.")
		request.requested[dataType] = append(request.requested[dataType], r)
		fmt.Println("Wrote to requested map.")

		input = trim(input[iR+1:])

		if len(input) <= 0 {
			break
		}
	}
	return request, nil
}

func (m *Menu) changeDataDisplayHelp() {

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
	if m.Name == "changeData" {
		m.displayData()
	}
}

func (m *Menu) displayData() {
	// fmt.Println("Menu data: ")
	for _, elem := range **m.data {
		fmt.Println(elem)
	}
}

func trim(input string) string {
	return strings.TrimSpace(input)
}
