package main

import (
	"fmt"
	"log"

	"github.com/dandytron/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("could not read config file: %v", err)
	}

	setErr := cfg.SetUser("Dandy")
	if setErr != nil {
		log.Fatalf("was not able to set username: %v", setErr)
	}

	newCfg, err := config.Read()
	if err != nil {
		log.Fatalf("could not read config file a second time: %v", err)
	}
	fmt.Printf("%s\n", newCfg.DbUrl)
}
