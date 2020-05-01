package network

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"net"
	"time"

	log "github.com/sirupsen/logrus"
)

//System holds the socket information for the network functions
type System struct {
	Host   string
	PSList []string
}

/*
   """Init."""
   self.sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
   self.sock.settimeout(6.0)
   self.sock.setsockopt(socket.SOL_SOCKET, socket.SO_BROADCAST, 1)
   self.msg = get_ddp_search_message()
   self.host = '255.255.255.255'
   self.ps_list = []*/

const maxBufferSize = 1024

//Search searches the network for devices
func (sys *System) Search(host string) {
	log.Debug("Searching for devices")
	pc, err := net.ListenPacket("udp", ":0")
	if err != nil {
		log.Error(err)
		return
	}
	defer pc.Close()
	buffer := make([]byte, maxBufferSize)

	go func() {
		for {

			n, addr, err := pc.ReadFrom(buffer)
			if err != nil {

				return
			}

			fmt.Printf("packet-received: bytes=%d from=%s\n",
				n, addr.String())

			deadline := time.Now().Add(1 * time.Minute)
			err = pc.SetWriteDeadline(deadline)
			if err != nil {

				return
			}

			// Write the packet's contents back to the client.
			n, err = pc.WriteTo(buffer[:n], addr)
			if err != nil {

				return
			}

			fmt.Printf("packet-written: bytes=%d to=%s\n", n, addr.String())
		}
	}()

	msg := fmt.Sprintf("%s * HTTP/1.1\ndevice-discovery-protocol-version:%s\n", "SRCH", "00020020")
	var buffer2 bytes.Buffer
	encoder := gob.NewEncoder(&buffer2)
	encoder.Encode(msg)
	addr, err := net.ResolveUDPAddr("udp", "255.255.255.255:987")
	if err != nil {
		log.Error(err)
	}

	for i := 0; i < 1; i++ {
		_, err = pc.WriteTo(buffer2.Bytes(), addr)
		if err != nil {
			log.Error(err)
		}
		log.Info("sent")
	}

}

/*def search(self, host):
  """Search for Devices."""
  if host is None:
      host = self.host
  null_responses = 0
  try:
      self.send(host)
  except (socket.error, socket.timeout):
      self.sock.close()
      return self.ps_list

  while null_responses < 3:
      try:
          device = self.receive()
          if device is not None:
              if device not in self.ps_list:
                  self.ps_list.append(device)
                  continue
          null_responses += 1
      except (socket.error, socket.timeout):
          self.sock.close()
          return self.ps_list

  return self.ps_list*/
