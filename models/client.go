package models

import ws "github.com/gorilla/websocket"

type Client struct {
	Name  string
	Score int
	Conn  *ws.Conn
	// send  chan []byte
}
