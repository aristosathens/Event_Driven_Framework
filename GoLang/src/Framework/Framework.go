package Framework

import (
	. "ServiceInterface"
	. "Services"
	"fmt"
	"strconv"
)

// ------------------------------------------- Framework ------------------------------------------- //

type Framework struct {
	Services        []*Service
	SendChannel     chan<- Event
	ReceiveChannel  <-chan Event
	sendChannels    []chan Event
	receiveChannels []chan Event
}

// Get list of all services and init each one
func (f *Framework) Init() {
	f.Services = make([]*Service, len(AllServiceInterfaces))
	f.receiveChannels = make([]chan Event, len(AllServiceInterfaces))
	f.sendChannels = make([]chan Event, len(AllServiceInterfaces))

	for i, serviceInterface := range AllServiceInterfaces {
		s := NewService(serviceInterface)
		f.Services[i] = &s
		f.receiveChannels[i] = f.Services[i].SendChannel
		f.sendChannels[i] = f.Services[i].ReceiveChannel
		go s.Run()
	}
}

// Runs the Framework. Monitors receiveChannel
func (f *Framework) Run() {
	fmt.Println("Starting Framework Run()...")
	f.Post(NewEvent(GLOBAL_START, ""))
	for {
		event := <-f.ReceiveChannel
		fmt.Println("event detected in Framework")
		f.Post(event)
		// f.Post(NewEvent(PING, ""))
		if event.Type == GLOBAL_EXIT {
			break
		}
	}
}

// Post an event to all Services
func (f *Framework) Post(event Event) {
	fmt.Println("Posting")
	for i, _ := range f.sendChannels {
		f.sendChannels[i] <- event
	}
}

// Waits until every service is set to inactive. This ensures all queued events are handled and channels are closed before exiting
func (f *Framework) Close() {
	f.Post(NewEvent(GLOBAL_EXIT, ""))
	for i, _ := range f.Services {
		service := f.Services[i]
		fmt.Println("Checking service for inactivity: ", service.Name)
		for {
			if service.Active == false {
				fmt.Println("Found inactive service")
				break
			}
		}
	}
}

// ------------------------------------------- Utilities ------------------------------------------- //

// Takes an array of input channels (i.e. service --> framework channels)
// Returns a single channel, which framework can use to monitor all services
// func mergeChannels(channels []chan Event) <-chan Event {
// 	output := make(chan Event, bufferSize)
// 	for _, channel := range channels {
// 		go func(output chan Event, channel <-chan Event) {
// 			for {
// 				event := <-channel
// 				output <- event
// 			}
// 		}(output, channel)
// 	}
// 	return output
// }

// func distributeChannels(channels []chan Event) chan<- Event {
// 	input := make(chan Event, bufferSize)
// 	for _, channel := range channels {
// 		go func(input chan Event, channel chan<- Event) {
// 			for {
// 				event := <-input
// 				channel <- event
// 			}
// 		}(input, channel)
// 	}
// 	return input
// }

// func (f Framework) WaitFor

// ------------------------------------------- Debugging ------------------------------------------- //

// To use, Uncomment InitBug() and RunDebug() in Main.go
// Rewrite Run() for testing rest of Framework

// var testCounter int

// func (f *Framework) InitDebug() {

// 	serviceNames := AllServiceNames()
// 	f.sendChannels = make([]chan Event, len(serviceNames))
// 	receiveChannels := make([]chan Event, len(serviceNames))
// 	f.Services = make([]*Service, len(serviceNames))
// 	for i, serviceName := range serviceNames {
// 		receiveChannel := make(chan Event, bufferSize)
// 		sendChannel := make(chan Event, bufferSize)
// 		receiveChannels[i] = make(chan Event, bufferSize)
// 		f.sendChannels[i] = sendChannel
// 		service := NewService(sendChannel, receiveChannel, serviceName)
// 		f.Services[i] = &service

// 		go service.Run()
// 	}
// }

// Runs the Framework.
func (f *Framework) RunDebug() {
	f.Post(NewEvent(GLOBAL_START, ""))
	for i := 0; i < 100; i++ {
		event := NewEvent(PING, strconv.Itoa(i))
		f.Post(event)
	}
}
