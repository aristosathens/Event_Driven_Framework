package Framework

import (
	. "Services"
	"fmt"
	"reflect"
)

type Framework struct {
	allServices []Service
	eventQueue  []Event
}

func (f Framework) Init() {
	allServices = AllServices()
	eventQueue := [1000]Event
}

func (f Framework) Run() {
	for {
		for _, event := eventQueue {
			for _, service := allServices {
				service.Post(event)
			}
			if event.Type == GLOBAL_EXIT {
				break
			}
		}
	}
}

