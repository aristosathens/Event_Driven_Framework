package Services_Crawler

import (
	// "encoding/json"
	. "Framework_Definitions"
	"JSON_Saver"
	"fmt"
	"github.com/kardianos/osext"
	// "os"
	// "path/filepath"
)

const ()

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
		returnEvent.Type = REQUEST_USER_INPUT
		returnEvent.Parameter = "Enter urls you would like to add: "

	case GLOBAL_EXIT:
		returnEvent.Type = FINISHED
	}

	fmt.Println("Exiting Crawler service")

	return returnEvent
}

// ------------------------------------------- Private ------------------------------------------- //

// ------------------------------------------- Utilities ------------------------------------------- //

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
