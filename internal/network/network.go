package network

import (
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
)

//System holds the socket information for the network functions
type System struct {
	Host   string
	PSList []string
}

const maxBufferSize = 1024

//Search searches the network for devices
func (sys *System) Search(host string) {
	log.Debug("Searching for devices")
	pc, err := net.ListenPacket("udp", ":0")
	if err != nil {
		log.Error(err)
		return
	}

	go func() {
		log.Debug("Listening for SRCH messages ...")
		buffer := make([]byte, maxBufferSize)

		n, addr2, err := pc.ReadFrom(buffer)
		if err != nil {

			return
		}

		fmt.Printf("packet-received: bytes=%d from=%s\n", n, addr2.String())
		log.Debug(string(buffer))
	}()

	msg := fmt.Sprintf("%s * HTTP/1.1\ndevice-discovery-protocol-version:%s\n", "SRCH", "00020020")
	addr, err := net.ResolveUDPAddr("udp", "255.255.255.255:987")
	if err != nil {
		log.Error(err)
	}
	_, err = pc.WriteTo([]byte(msg), addr)
	if err != nil {
		log.Error(err)
	}
	log.Info("sent")

}
