package udp_server

import (
	"log"
	"net"
	"strings"
)

// TCP server
type server struct {
	address           string // Address to open connection: localhost:9999
	onNewMessage      func(message string)
	close             chan bool
	ReadBuffer        int
	MessageTerminator rune
}

// Start network server
func (s *server) Listen() {
	address, addressError := net.ResolveUDPAddr("udp", s.address)
	if addressError != nil {
		log.Println(addressError)
		return
	}

	listener, err := net.ListenUDP("udp", address)
	if err != nil {
		log.Println(err)
		return
	}

	go func() {
		<-s.close
		listener.Close()
	}()

	listener.SetReadBuffer(s.ReadBuffer)

	buffer := make([]byte, s.ReadBuffer)
	var message strings.Builder
	for {
		len, _, readErr := listener.ReadFrom(buffer)
		if readErr != nil {
			log.Println(readErr)
			listener.Close()
			return
		}
		partial := string(buffer[0:len])
		endOfMessage := strings.IndexRune(partial, s.MessageTerminator)
		if endOfMessage >= 0 {
			message.Write(buffer[0:endOfMessage])
			s.onNewMessage(message.String())
			message.Reset()
			message.Write(buffer[endOfMessage+1 : len])
		} else {
			message.Write(buffer[0:len])
		}
	}
}

func (s *server) OnNewMessage(callback func(message string)) {
	s.onNewMessage = callback
}

func (s *server) Close() {
	s.close <- true
}

func New(address string) *server {
	log.Println("Creating server with address", address)
	server := &server{
		address:           address,
		close:             make(chan bool, 1),
		MessageTerminator: '\n',
		ReadBuffer:        1024 * 8,
	}

	server.OnNewMessage(func(message string) {})

	return server
}
