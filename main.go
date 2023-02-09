package main

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

func main() {
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
