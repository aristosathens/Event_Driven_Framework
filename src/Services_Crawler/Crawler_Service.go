package Services_Crawler

import (
	// "encoding/json"
	. "Framework_Definitions"
	"JSON_Saver"
	"fmt"
	"github.com/kardianos/osext"
	"net/url"
	// "strings"
	// "os"
	// "path/filepath"
)

// ------------------------------------------- Public ------------------------------------------- //

type CrawlerService struct {
	path     string
	websites []string
}

func (s *CrawlerService) Init() string {
	path, err := osext.ExecutableFolder()
	checkError(err)

	err = JSONSaver.Load(path+"\\websites.JSON", &s.websites)
	checkError(err)

	return "CrawlerService"
}

func (s *CrawlerService) RunFunction(event Event, sendChannel chan Event) Event {
	fmt.Println("Entering Crawler Service Run Function...")
	returnEvent := NewEvent(NONE, "", "")

	switch event.Type {

	case GLOBAL_START:

	case GLOBAL_EXIT:
		returnEvent.Type = FINISHED

	case USER_INPUT:
	}

	fmt.Println("Exiting Crawler service")
	return returnEvent
}

// ------------------------------------------- Private ------------------------------------------- //

// func (s *CrawlerService) updateScreen() {
// 	switch s.status {
// 	case MENU:
// 	case INPUT_NEW_WEBSITES:
// 	}
// }

// ------------------------------------------- Utilities ------------------------------------------- //

// Respond to errors
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

// Checks if string is a valid URL
func isValidUrl(toCheck string) bool {
	_, err := url.ParseRequestURI(toCheck)
	if err != nil {
		return false
	} else {
		return true
	}
}
