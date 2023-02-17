package main

import (
	v1 "dbielawski/blog/internal/v1"
	"flag"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	var (
		clientServerAddr = flag.String("clientlisten", "127.0.0.1:8080", "listen address for client API")
	)
	flag.Parse()

	logger := log.New()
	logger.Trace("Logger initialized")

	server := v1.NewServer(logger)

	err := server.Init()
	if err != nil {
		logger.Error("Couldn't initialize the server", err)
		os.Exit(-1)
	}
	err = server.Run(*clientServerAddr)
	if err != nil {
		logger.Error("Couldn't start the server", err)
		os.Exit(-1)
	}
}
