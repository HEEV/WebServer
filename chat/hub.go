package chat

import (
	log "github.com/sirupsen/logrus"
)

// BroadcastMessage represents a data/client pair that will be passed to the
// hub to broadcast messages to all connected clients other than sending one
type BroadcastMessage struct {
	Data   []byte
	Client *Client
}

// Hub keeps track of all the information for the connected clients
type Hub struct {
	// Registered clients, by UUID string
	clients map[string]*Client

	// Inbound messages from the clients
	broadcast chan BroadcastMessage

	// Queue of client registration requests
	register chan *Client

	// Queue of client unregister requests
	unregister chan *Client
}

// NewHub creates a new websocket hub
func NewHub() *Hub {
	return &Hub{
		broadcast:  make(chan BroadcastMessage),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[string]*Client),
	}
}

// NumClients returns the number of clients currently active
func (h *Hub) NumClients() int {
	return len(h.clients)
}

// Run runs the websocket hub
func (h *Hub) Run() {
	// Server loop, will be run in a goroutine
	for {
		select {
		case client := <-h.register:
			// Register a new client and set it as active
			h.clients[client.uid] = client

			log.Infof("New connection! (%s) from %s", client.uid, client.conn.RemoteAddr().String())

		case client := <-h.unregister:
			// Check if the client exists
			if _, ok := h.clients[client.uid]; ok {
				// Remove the client
				delete(h.clients, client.uid)

				// Close the client's message channel
				close(client.send)

				log.Infof("Connection %s has disconnected", client.uid)
			}

		case message := <-h.broadcast:
			for uid, client := range h.clients {
				// Avoid broadcasting message back to sending client
				if client == message.Client {
					continue
				}

				select {
				// Attempt to put the message onto the client's message channel
				case client.send <- message.Data:
				default:
					// If the message write failed then assume client is dead
					close(client.send)
					// Remove the client from the clients list
					delete(h.clients, uid)
				}
			}
		}
	}
}
