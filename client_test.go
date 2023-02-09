package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/kelseyhightower/envconfig"
	"github.com/nbd-wtf/go-nostr"
	"github.com/nbd-wtf/go-nostr/nip19"
)

func TestMain(m *testing.M) {
	go runServer()
	os.Exit(m.Run())
}
func runServer() {
	r := HubRelay{}
	if err := envconfig.Process("", &r); err != nil {
		log.Fatalf("failed to read from env: %v", err)
		return
	}
	r.storage = &HubStorage{}
	// r.storage = &postgresql.PostgresBackend{DatabaseURL: r.PostgresDatabase}
	if err := Start(&r); err != nil {
		log.Fatalf("server terminated: %v", err)
	}
}
func TestClient(t *testing.T) {

	relay, err := nostr.RelayConnect(context.Background(), "ws://localhost:7447")
	if err != nil {
		panic(err)
	}
	log.Println(relay)
	defer relay.Close()

}

func TestGetAUser(t *testing.T) {
	sk := nostr.GeneratePrivateKey()
	pk, _ := nostr.GetPublicKey(sk)
	nsec, _ := nip19.EncodePrivateKey(sk)
	npub, _ := nip19.EncodePublicKey(pk)

	fmt.Println("sk:", sk)
	fmt.Println("pk:", pk)
	fmt.Println(nsec)
	fmt.Println(npub)
}
