package main

type Hub struct {
	// A connection mapped to a GameId
	connections map[Connection]int

	//broadcast chan []byte

	// Register connection into hub
	register chan Connection

	// Remove connection from hub
	unregister chan Connection
}

var noGameHub = NewHub()

func NewHub() Hub {
	return Hub{
		//broadcast:   make(chan []byte),
		register:    make(chan Connection),
		unregister:  make(chan Connection),
		connections: make(map[Connection]int),
	}
}

// Listen for events and manage connections
func (h *Hub) Run() {
	for {
		select {
		case c := <-h.register:
			h.connections[c] = 0
		case c := <-h.unregister:
			if h.checkForConnection(c) {
				delete(h.connections, c)
				if err := c.CloseChannel(); err != nil {
					panic("aw fudge")
				}
			}
		}
	}
}

// Change a connections game id if they move to a new game
func (h *Hub) ChangeGame(Conn Connection, GameId int) {
	h.connections[Conn] = GameId
}

// Removes game id when a player leaves a game
func (h *Hub) LeaveGame(Conn Connection) {
	h.connections[Conn] = 0
}

func (h *Hub) checkForConnection(conn Connection) bool {
	if _, ok := h.connections[conn]; ok {
		return true
	}

	return false
}
