package Services_GIS

import (
	. "Framework_Definitions"
	"strconv"
	// . "Utilities"
	// "container/list"
	"errors"
	"fmt"
	"strings"
	// "text/tabwriter"
)

const ()

type Menu struct {
	Name         string
	Elements     map[string]MenuItem
	InputHandler MenuItem
	requestData  **[]dataRequest
	tempData     **[]dataRequest
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

	sharedRequests := &[]dataRequest{}
	sharedTemp := &[]dataRequest{}

	main = Menu{
		Name: "main",
		Elements: map[string]MenuItem{
			"data": MenuItem{"Change data", &changeData},
			"help": MenuItem{"Help", &help},
			"gen":  MenuItem{"Generate map", NewEvent(GENERATE_MAP, &sharedRequests, "")},
			"exit": MenuItem{"Exit", NewEvent(GLOBAL_EXIT, "", "")},
		},
		requestData: &sharedRequests,
		tempData:    &sharedTemp,
	}

	changeData = Menu{
		Name: "changeData",
		Elements: map[string]MenuItem{
			"exit":    MenuItem{"Main Menu", &main},
			"display": MenuItem{"Display current data", (*Menu).displayData},
			"help":    MenuItem{"Help", (*Menu).changeDataDisplayHelp},
		},
		InputHandler: MenuItem{"What data would you like to add? ", (*Menu).changeDataInputHandler},
		requestData:  &sharedRequests,
		tempData:     &sharedTemp,
	}

	help = Menu{
		Name: "help",
		Elements: map[string]MenuItem{
			"exit": MenuItem{"Main Menu", &main},
		},
		requestData: &sharedRequests,
		tempData:    &sharedTemp,
	}

	return main
}

// ------------------------------------------- Change Data Menu ------------------------------------------- //

// Create empty dataRequest, then fill it with parsed user input.
func (m *Menu) changeDataInputHandler(input string) Event {
	request := dataRequest{"", map[string][]dataRange{}}
	request, err := parseDataRequest(input, request)
	if err != nil {
		fmt.Println(err)
		return NewEvent(NONE, "", "")
	} else {
		**(m.tempData) = append(**(m.tempData), request)
		return NewEvent(CHECK_DATA_REQUEST, request, "")
	}
}

// Parses one line of user input. Exepects a single data type and one or more ranges..
// Takes in a preexisint request and fills out a single map key/element in the dataRequest.requested
func parseDataRequest(input string, request dataRequest) (dataRequest, error) {
	input = trim(input)

	// Get dataset name (e.g. cities, countries)
	i := strings.Index(input, " ")
	if i < 0 {
		return request, errors.New("Improperly formatted input.")
	}
	request.datasetName = trim(input[:i])
	input = trim(input[i:])

	// Get data type (e.g. population, age)
	i = strings.Index(input, "[")
	if i < 0 {
		return request, errors.New("Improperly formatted input.")
	}
	dataTypeName := trim(input[:i])

	// Get data ranges (e.g. [0, 200])
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

		request.requested[dataTypeName] = append(request.requested[dataTypeName], r)
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

// ------------------------------------------- General ------------------------------------------- //

func (m *Menu) displayMenu() {
	for key, item := range m.Elements {
		fmt.Println(" ( " + key + " ) " + item.Label)
	}
}

// Prints currently defined aliases
// func displayAliases() {
// 	tabWriter := tabwriter.NewWriter(os.Stdout, 0, 0, 0, ' ', tabwriter.Debug)
// 	fmt.Fprintln(tabWriter, "NAME \t COMMAND(S)")
// 	fmt.Fprintln(tabWriter, "---- \t ----------")

// 	for _, alias := range currentAliases {
// 		command := getAliasCommand(alias)
// 		leftCol := alias
// 		for {
// 			i := strings.Index(command, "\n")
// 			if i == -1 {
// 				i = len(command)
// 				fmt.Fprintln(tabWriter, leftCol+" \t "+command[:i])
// 				break
// 			}
// 			column := leftCol + " \t " + command[:i]
// 			fmt.Fprintln(tabWriter, column)
// 			command = command[i+1:]
// 			leftCol = ""
// 		}
// 	}
// 	tabWriter.Flush()
// }

func (m *Menu) displayData() {
	for _, elem := range **m.requestData {
		fmt.Println(elem.datasetName)
		for name, val := range elem.requested {
			fmt.Print(name)
			for _, set := range val {
				fmt.Print(" [" + strconv.FormatFloat(set.lower, 'f', 2, 64) + ", " + strconv.FormatFloat(set.upper, 'f', 2, 64) + "]")
			}
			fmt.Print("\n")
		}
	}
}

func (m *Menu) clearTempData() {
	req := &[]dataRequest{}
	*m.tempData = req
}

// ------------------------------------------- Utility ------------------------------------------- //

func trim(input string) string {
	return strings.TrimSpace(input)
}

// Merge request into array of requests. Properly handles dataset name and data type
func mergeRequests(requests []dataRequest, req dataRequest) []dataRequest {
	for _, request := range requests {
		// if req is for a dataset already in requests
		if req.datasetName == request.datasetName {
			for key, _ := range req.requested {
				// if key exists in requests
				if _, ok := request.requested[key]; ok {
					// add all ranges from req to slice of ranges in request
					for _, myRange := range req.requested[key] {
						request.requested[key] = append(request.requested[key], myRange)
					}
				} else {
					request.requested[key] = req.requested[key]
				}
			}
			return requests
		}
	}
	requests = append(requests, req)
	return requests
}
