package models

type Table struct {
	Name    string
	Clients map[*Client]bool
	// register:   make(chan *client.Client),
	// unregister: make(chan *client.Client),

}
