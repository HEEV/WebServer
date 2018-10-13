package chat

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/satori/go.uuid"

	"../packets"
	"../sql"

	log "github.com/sirupsen/logrus"
)

// Websocket code adapted from https://github.com/gorilla/websocket/tree/master/examples/chat

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

// Client keeps track of the information related to the websocket connection
type Client struct {
	hub *Hub

	// Websocket connection
	conn *websocket.Conn

	// Buffered channel of messages
	send chan []byte

	// UUID string for this client
	uid string
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(
		func(string) error {
			c.conn.SetReadDeadline(time.Now().Add(pongWait))
			return nil
		},
	)

	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(
				err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Errorf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))

		// Packet identification information, all packets must include these
		var ident packets.Identification

		// Unmarshal the data into the identification packet
		err = json.Unmarshal(message, &ident)
		if err != nil {
			// There was an error unmarshalling the packet identification
			log.Warnf("Invalid packet from client %s!", c.uid)
			log.Warn("Identification information not included!")
			log.Warn("Data received:")
			log.Warn(string(message))

			// Drop the error after logging and wait for next message
			continue
		}

		// Generate the client's UUID
		uid, err := uuid.NewV4()
		if err != nil {
			log.Error(err)
			return
		}

		response := packets.Identification{
			AndroidID: ident.AndroidID,
			MessageID: uid.String(),
		}

		switch ident.MessageType {
		case "GetNextRunNumber":
			// Identification/initial communication
			// Create packet w/ run number value from server
			nextRun, err := sql.GetNextRunNumber(ident.AndroidID)
			if err != nil {
				log.Error("Unable to retrieve next run ID")
				log.Error(err)
			}

			data := packets.NextRunNumberData{
				RunNumber: nextRun,
			}
			response.Data = data
			response.MessageType = "NextRunNumberResponse"

			// Marshal response for client
			bytes, err := json.Marshal(response)
			if err != nil {
				log.Error("Unable to marshal GetNextRunNumber response")
				log.Error(err)
			}

			log.Infof(
				"Sending %+v response to connection %s",
				string(message),
				c.uid,
			)

			// Send response to client
			c.send <- bytes

		case "Log":
			// Log value to database

		}

		// Number of receiving clients = Total # of clients - 1 for sender
		numRecv := c.hub.NumClients() - 1

		log.Infof(
			"Connection %s sending message \"%+v\" to %d other clients",
			c.uid,
			string(message),
			numRecv,
		)

		// Broadcast the message out to all other connected clients
		c.hub.broadcast <- BroadcastMessage{message, c}
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)

	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// ServeWs handles websocket requests
func ServeWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to websocket
	conn, err := upgrader.Upgrade(w, r, nil)

	// Ensure nothing went wrong
	if err != nil {
		log.Error(err)
		return
	}

	// Create the new client object
	// Generate the client's UUID
	uid, err := uuid.NewV4()
	if err != nil {
		log.Error(err)
		return
	}

	// Generate client object
	client := &Client{
		hub:  hub,
		conn: conn,
		send: make(chan []byte, 256),
		uid:  uid.String(),
	}
	// Queue the new client to be registered with the hub
	client.hub.register <- client

	// Create a goroutine for the client to watch for new messages from client
	go client.readPump()
	// Create a goroutine for the client to send new messages to client
	go client.writePump()
}
