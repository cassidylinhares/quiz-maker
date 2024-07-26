package services

// Purpose: keep track of name, connection, score

import (
	"fmt"
	"log"
	"net/http"

	ws "github.com/gorilla/websocket"
)

var upgrader = ws.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// allow any connection
	CheckOrigin: func(r *http.Request) bool { return true },
}

type Client struct {
	Name  string
	Score int
	Conn  *ws.Conn
	// send  chan []byte
}

func newClient(conn *ws.Conn) *Client {
	return &Client{
		Name:  "Person",
		Score: 0,
		Conn:  conn,
	}
}

// Handle websocket req from clients
func ServeWs(server *Server, w http.ResponseWriter, r *http.Request) (*Client, error) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal("Failed: ", err)
		return nil, err
	}

	client := newClient(conn)
	fmt.Println("New Client joined the game.")
	fmt.Println(client)

	server.Register <- RegisterClient{client, "bird"}

	return client, nil
}

func (c *Client) IncrementScore() {
	c.Score++
}

func (c *Client) PrintPlayerInfo() {
	fmt.Printf("Name: %s\nScore:%d\n", c.Name, c.Score)
}
