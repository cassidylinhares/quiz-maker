package services

// Purpose: Make tables, add clients to tables,

type RegisterClient struct {
	Client    *Client
	TableName string
}

type Server struct {
	Clients    map[*Client]bool
	Tables     map[string]*Table
	Register   chan RegisterClient
	Unregister chan RegisterClient
}

func NewServer(tableNames []string) *Server {
	server := &Server{
		Clients:    make(map[*Client]bool),
		Tables:     make(map[string]*Table),
		Register:   make(chan RegisterClient),
		Unregister: make(chan RegisterClient),
	}

	for _, tableName := range tableNames {
		server.Tables[tableName] = NewTable(tableName)
	}

	return server
}

func (s *Server) Run() {
	for {
		select {
		case registerClient := <-s.Register:
			s.registerClient(registerClient)
		case registerClient := <-s.Unregister:
			s.unregisterClient(registerClient)
		}
	}
}

func (s *Server) registerClient(registerClient RegisterClient) {
	s.Clients[registerClient.Client] = true
	table := s.Tables[registerClient.TableName]
	table.AddClientToTable(registerClient.Client)
}

func (s *Server) unregisterClient(registerClient RegisterClient) {
	if _, exist := s.Clients[registerClient.Client]; exist {
		delete(s.Clients, registerClient.Client)
	}

	if _, exist := s.Tables[registerClient.TableName]; exist {
		table := s.Tables[registerClient.TableName]
		table.DeleteClientFromTable(registerClient.Client)
	}
}
