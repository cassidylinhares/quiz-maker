package services

import (
	"fmt"
)

// Purpose: hold guest in each table and keep track of the highest score
type Table struct {
	Name    string
	Clients map[*Client]bool
	// register:   make(chan *client.Client),
	// unregister: make(chan *client.Client),

}

func NewTable(name string) *Table {
	return &Table{
		Name:    name,
		Clients: make(map[*Client]bool),
		// register:   make(chan *client.Client),
		// unregister: make(chan *client.Client),
	}
}

func (t *Table) AddClientToTable(client *Client) {
	t.Clients[client] = true
}

func (t *Table) DeleteClientFromTable(client *Client) {
	if _, exist := t.Clients[client]; exist {
		delete(t.Clients, client)
	}
}

func (t *Table) getClients() []Client {
	// get slice of comment
	clients := make([]Client, 0, len(t.Clients))
	for c := range t.Clients {
		clients = append(clients, *c)
	}
	return clients
}

func (t *Table) GetClientWithHighestScore() Client {
	clients := t.getClients()

	cHighest := clients[0]
	for _, c := range clients {
		if c.Score > cHighest.Score {
			cHighest = c
		}
	}
	return cHighest
}

func (t *Table) PrintTableInfo() {
	clients := t.getClients()
	highestScore := t.GetClientWithHighestScore()

	fmt.Printf("Table Name: %s\nPlayer with Highest Score:%s\nHighest Score:%d\n", t.Name, highestScore.Name, highestScore.Score)
	fmt.Printf("Guest at table %s:\n", t.Name)

	for _, c := range clients {
		fmt.Printf("Name: %s\nScore:%d\n", c.Name, c.Score)
	}
}
