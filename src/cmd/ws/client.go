package ws

import (
	"github.com/gorilla/websocket"
)

type Client struct {
	ID     int
	Ws     *websocket.Conn
	Server Server
	Ch     chan *Message
	DoneCh chan bool
}
