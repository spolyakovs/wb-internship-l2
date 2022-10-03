package main

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/spolyakovs/wb-internship-l2/develop/dev11/internal/server"
)

func main() {
	configPath := "./configs/local.toml"

	config, err := server.MakeConfigFromFile(configPath)
	if err != nil {
		log.Fatalf("couldn't read config from file:%v\n\tpath:%s", err, configPath)
	}

	if err := server.Start(config); err != nil {
		log.Fatalf("couldn't start server:%v\n\t", err)
	}
}
