package main

import (
	"context"
	"log"

	"github.com/nbd-wtf/go-nostr"
)

type HubServer struct {
	Address  string
	IsVip    bool
	publish  *nostr.Relay
	listener map[string]*nostr.Subscription
}

func (h *HubServer) Run() error {
	relay, e := nostr.RelayConnect(context.Background(), h.Address)
	if e != nil {
		return e
	}
	h.publish = relay
	return nil
}

func (h *HubServer) Save(event nostr.Event) {
	h.publish.Publish(context.Background(), event)
}
func (h *HubServer) Delete(id, pubkey string) {}
func (h *HubServer) QueryEvents(filter nostr.Filters) (events []nostr.Event, err error) {
	log.Println("print filters ", filter)
	// sub := h.publish.Subscribe(context.Background(), filter)
	// go func() {
	// 	<-sub.EndOfStoredEvents
	// 	// handle end of stored events (EOSE, see NIP-15)
	// }()
	// for ev := range sub.Events {
	// 	log.Println("Subscribe ================= ", ev)
	// 	events = append(events, *ev)
	// }
	// log.Println("============== in ", events)
	return events, nil
}
