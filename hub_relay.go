package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
	"github.com/nbd-wtf/go-nostr"
)

type HubRelay struct {
	storage *HubStorage
}

func (r *HubRelay) Name() string {
	return "RelayHub"
}

func (r *HubRelay) Storage() Storage {
	return r.storage
}

func (r *HubRelay) OnInitialized(*Server) {}

func (r *HubRelay) Init() error {
	err := envconfig.Process("", r)
	if err != nil {
		return fmt.Errorf("couldn't process envconfig: %w", err)
	}

	return nil
}
func (r *HubRelay) UserAuth(key string) {
	log.Println("User auth ..............")
}
func (r *HubRelay) UserExit(key string) {
	log.Println("User exit................")
}

func (r *HubRelay) AcceptEvent(evt *nostr.Event) bool {
	// block events that are too large
	jsonb, _ := json.Marshal(evt)
	if len(jsonb) > 10000 {
		return false
	}

	return true
}

func (r *HubRelay) BeforeSave(evt *nostr.Event) {
	// do nothing
}

func (r *HubRelay) AfterSave(evt *nostr.Event) {
	// delete all but the 100 most recent ones for each key

}
