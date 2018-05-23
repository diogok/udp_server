package udp_server

import (
	. "github.com/smartystreets/goconvey/convey"
	"log"
	"net"
	"testing"
	"time"
)

func buildTestServer() *server {
	return New("127.0.0.1:9999")
}

func Test_accepting(t *testing.T) {
	server := buildTestServer()

	var messagesText []string

	server.OnNewMessage(func(message string) {
		log.Println(message)
		messagesText = append(messagesText, message)
	})
	go server.Listen()

	// Wait for server
	time.Sleep(100 * time.Millisecond)

	conn, err := net.Dial("udp", "127.0.0.1:9999")
	if err != nil {
		t.Fatal("Failed to connect to test server")
	}
	time.Sleep(10 * time.Millisecond)
	conn.Write([]byte("Test message\nOther"))
	time.Sleep(10 * time.Millisecond)
	conn.Write([]byte(" message\n"))
	time.Sleep(10 * time.Millisecond)

	conn.Close()
	server.Close()

	Convey("Messages should be equal", t, func() {
		So(len(messagesText), ShouldEqual, 2)
		So(messagesText[0], ShouldEqual, "Test message")
		So(messagesText[1], ShouldEqual, "Other message")
	})
}
