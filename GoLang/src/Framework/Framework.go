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
	ReceiveChannel  <-chan Event
	sendChannels    []chan<- Event
	receiveChannels []<-chan Event
}

// Get list of all services and init each one
func (f *Framework) Init() {
	f.Services = make([]*Service, len(AllServiceInterfaces))
	f.receiveChannels = make([]<-chan Event, len(AllServiceInterfaces))
	f.sendChannels = make([]chan<- Event, len(AllServiceInterfaces))

	for i, serviceInterface := range AllServiceInterfaces {
		s := NewService(serviceInterface)
		f.Services[i] = &s
		f.receiveChannels[i] = f.Services[i].SendChannel
		f.sendChannels[i] = f.Services[i].ReceiveChannel
		go s.Run()
	}
	f.ReceiveChannel = f.mergeChannels(f.receiveChannels)
}

// Runs the Framework. Monitors receiveChannel
func (f *Framework) Run() {
	f.Post(NewEvent(GLOBAL_START, ""))
	for {
		select {
		case event := <-f.ReceiveChannel:
			fmt.Println("event detected in Framework")
			f.Post(event)
			if event.Type == GLOBAL_EXIT {
				return
			}
		}
	}
}

// Post an event to all Services
func (f *Framework) Post(event Event) {
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
	for _, ch := range f.sendChannels {
		close(ch)
	}
}

// ------------------------------------------- Utilities ------------------------------------------- //

// Takes an array of input channels (i.e. framework <- service channels)
// Returns a single channel, which framework can use to monitor all services
func (f *Framework) mergeChannels(channels []<-chan Event) <-chan Event {
	aggregateChannel := make(chan Event)
	for _, ch := range channels {
		go func(c <-chan Event) {
			for {
				msg, flag := <-c
				if !flag {
					break
				}
				aggregateChannel <- msg
			}
			// close(aggregateChannel)
		}(ch)
	}
	return aggregateChannel
}

// ------------------------------------------- Debugging ------------------------------------------- //

// To use, Uncomment InitDebug() and RunDebug() in Main.go

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
