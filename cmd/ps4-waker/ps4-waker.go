package main

import (
	"github.com/sevren/go-ps4-waker/internal/network"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Initalizing ps4-waker")
	log.SetLevel(log.DebugLevel)
	sys := &network.System{}
	sys.Search("")
	for {
	}

}
