package Services_Crawler

import (
	// "encoding/json"
	. "Framework_Definitions"
	"JSON_Saver"
	. "Utilities"

	"github.com/kardianos/osext"
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
	CheckError(err)

	err = JSONSaver.Load(path+"\\websites.JSON", &s.websites)
	CheckError(err)

	return "CrawlerService"
}

func (s *CrawlerService) RunFunction(event Event, sendChannel chan Event) Event {
	// fmt.Println("Entering Crawler Service Run Function...")
	returnEvent := NewEvent(NONE, "", "")

	switch event.Type {

	case GLOBAL_START:

	case GLOBAL_EXIT:
		returnEvent = NewEvent(FINISHED, "", "")

	case REQUEST_WEBSITE_LIST:
		returnEvent = NewEvent(WEBSITE_LIST, s.websites, "")

	}

	// fmt.Println("Exiting Crawler service")
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
