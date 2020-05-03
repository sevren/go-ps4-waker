package main

import (
	"github.com/sevren/go-ps4-waker/internal/ddp"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Info("Initalizing ps4-waker")
	log.SetLevel(log.DebugLevel)
	sys := &ddp.System{}
	//sys.Search()
	sys.Credential()
	for {
	}

}
