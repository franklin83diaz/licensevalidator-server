package main

import (
	"log"
	"os"

	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/toml"
)

func main() {
	configFile := "/etc/licvs/default.conf"
	if os.Getenv("ENV") == "debug" {
		log.SetFlags(log.LstdFlags | log.Lshortfile)
		configFile = "default.conf"
	}

	// Add driver to load TOML files
	config.WithOptions(config.ParseEnv)
	config.AddDriver(toml.Driver)

	// Load config from files
	err := config.LoadFiles(configFile)
	if err != nil {
		log.Fatal(err)
	}
}
