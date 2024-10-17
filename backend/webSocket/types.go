package webSocket

import (
	"pc3r/prisma/db"

	"github.com/gorilla/websocket"
)

type responseRegisterToChat struct {
	Message string `json:"message"`
	Chat_id 	string `json:"chat_id"`
}

// We consider Hub as a chat room,
type Hub struct {
	// many clients could be connected to a hub
	clients map[*Client]bool

	// In order to broadcast a message to all clients connected to the hub
	broadcast chan Message

	// Register requests from clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

// Message is a object used to pass data on sockets.
type Message struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}

// FindHandler is a type that defines handler finding functions.
type FindHandler func(Event) (Handler, bool)

// Client is a type that reads and writes on sockets.
type Client struct {
	send           chan Message
	socket         *websocket.Conn
	rt             *Router
	findHandler    FindHandler
	user           *db.UserModel
	SubscribedHubs []*Hub // Ajouter un champ pour stocker les hubs auxquels le client est abonné
}