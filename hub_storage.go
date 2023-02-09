package main

import (
	"log"
	"sync"

	"github.com/nbd-wtf/go-nostr"
)

type HubStorage struct {
	clients []*HubServer
}

func (hs *HubStorage) Init() error {
	log.Println("====== HubStorage init ")
	server := HubServer{
		Address: "wss://relay.damus.io",
		IsVip:   false,
	}
	if server.Run() == nil {
		hs.clients = append(hs.clients, &server)
	}
	return nil
}

func (hs *HubStorage) QueryEvents(filter *nostr.Filter) (events []nostr.Event, err error) {
	eventschan := make(chan nostr.Event, 1000)
	var rg sync.WaitGroup
	rg.Add(len(hs.clients))
	for _, relay := range hs.clients {
		go func(relay *HubServer) {
			if events, err = relay.QueryEvents([]nostr.Filter{*filter}); err == nil {
				for _, event := range events {
					eventschan <- event
				}
			}
			rg.Done()
		}(relay)
	}
	go func() {
		rg.Wait()
		close(eventschan)
	}()
	log.Println("env................... end ...")
	for {
		if event, ok := <-eventschan; ok {
			for _, e := range events {
				if e.ID == event.ID {
					continue
				}
			}
			events = append(events, event)
		} else {
			log.Println("exit .........", len(events))
			return events, nil
		}
	}

}
func (hs *HubStorage) DeleteEvent(id string, pubkey string) error {
	log.Println("Delete event", id, pubkey)
	for _, relay := range hs.clients {
		go relay.Delete(id, pubkey)
	}
	return nil
}
func (hs *HubStorage) SaveEvent(event *nostr.Event) error {
	for _, relay := range hs.clients {
		go relay.Save(*event)
	}

	log.Println("Save event", event)
	return nil
}
