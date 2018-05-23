# UDPServer

Package udp_server created to help build UDP servers faster.

Inspired by [firstrow/tcp_server](https://github.com/firstrow/tcp_server).

### Install package

``` bash
go get -u github.com/diogok/udp_server
```

### Usage:

NOTICE: `OnNewMessage` callback will receive new message only if it's ending with `\n` or the setup MessageTerminator

``` go
package main

import "github.com/diogok/udo_server"

func main() {
	server := udp_server.New("localhost:9999")
  server.MessageTerminator='\n' // Optional end of message byte, default to newline.

	server.OnNewMessage(func(c *tcp_server.Client, message string) {
		// new message received
	})
	server.Listen()
  server.Close()
}
```

## License

MIT

