package ddp

import (
	"fmt"
	"net"
	"strings"

	log "github.com/sirupsen/logrus"
)

//System holds the socket information for the network functions
type System struct {
	Host   string
	PSList []string
}

//Packet basic DDP structure
type Packet struct {
	Status  string
	Data    interface{}
	Version string
}

//Host host information
type Host struct {
	ID          string
	Type        string
	Name        string
	RequestPort string
}

//PS4 ps4 type
type PS4 struct {
	IP      string
	Status  string
	Host    Host
	Version string
}

//TYPES - TODO: Refactor to enum

//SRCH Search type
const SRCH string = "SRCH"

//LAUNCH Search type
const LAUNCH string = "LAUNCH"

//WAKEUP Search type
const WAKEUP string = "WAKEUP"

//STANDBY Search type
const STANDBY string = "STANDBY"

const version string = "00020020"

const maxBufferSize = 1024

//Search searches the network for devices
func (sys *System) Search() {
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

//Credential attempts to retrieve credentials from the ps4 second screen app on Android or IOS
func (sys *System) Credential() {
	log.Debug("Searching for devices")
	pc, err := net.ListenPacket("udp", ":987")
	if err != nil {
		log.Error(err)
		return
	}

	go func() {
		log.Debug("Listening for DDP messages ...")
		buffer := make([]byte, maxBufferSize)

		n, addr2, err := pc.ReadFrom(buffer)
		if err != nil {

			return
		}

		fmt.Printf("packet-received: bytes=%d from=%s\n", n, addr2.String())
		log.Debug(string(buffer))
		ddp := parse_ddp_response(buffer)
		log.Debugf("Parsed response: %+v", ddp)
	}()

	// msg := fmt.Sprintf("%s * HTTP/1.1\ndevice-discovery-protocol-version:%s\n", "SRCH", "00020020")
	// addr, err := net.ResolveUDPAddr("udp", "255.255.255.255:987")
	// if err != nil {
	// 	log.Error(err)
	// }
	// _, err = pc.WriteTo([]byte(msg), addr)
	// if err != nil {
	// 	log.Error(err)
	// }
	// log.Info("sent")
}

// func getDDPMessage(b []byte) string {
// 	header:=fmt.Sprintf("HTTP/1.1 %s\n",status)
// }

// parse_ddp_response parsees the ddp response
func parse_ddp_response(bs []byte) *Packet {
	str := string(bs)
	if strings.Contains(str, SRCH) {
		log.Debug("Found a SRCH message")
		return &Packet{
			Status: SRCH,
		}
	}
	log.Debug("Unknown message")
	return &Packet{}
}
