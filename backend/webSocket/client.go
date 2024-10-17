package webSocket

import (
	"fmt"
	"log"
	"pc3r/prisma/db"

	"github.com/gorilla/websocket"
)

// NewClient accepts a socket and returns an initialized Client.
func NewClient(rt *Router, socket *websocket.Conn, findHandler FindHandler, user *db.UserModel) *Client {
	return &Client{
		send:           make(chan Message),
		socket:         socket,
		findHandler:    findHandler,
		rt:             rt,
		SubscribedHubs: []*Hub{},
		user:           user,
	}
}

// Write receives messages from the channel and writes to the socket.
func (c *Client) Emit(msg Message) {
	go func() {
		c.send <- msg
	}()
	c.Write()
}

// Write receives messages from the channel and writes to the socket.
func (c *Client) Write() {
	msg := (<-c.send)
	err := c.socket.WriteJSON(msg)
	if err != nil {
		log.Printf("socket write error: %v\n", err)
	}
}

// Read intercepts messages on the socket and assigns them to a handler function.
func (c *Client) Read() {
	var msg Message
	for {
		// read incoming message from socket
		if err := c.socket.ReadJSON(&msg); err != nil {
			log.Printf("socket read error: %v\n", err)
			break
		}
		// assign message to a function handler
		if handler, found := c.findHandler(Event(msg.Event)); found {
			handler(c, msg.Data)
		}
	}

	// close interrupted socket connection
	for _, hub := range c.SubscribedHubs {
		go func(h *Hub) {
			h.unregister <- c
		}(hub)
	}
	c.SubscribedHubs = []*Hub{}
	c.socket.Close()
}

// Ajouter un hub aux hubs abonnés du client
func (c *Client) AddSubscribedHub(hub *Hub) {
	c.SubscribedHubs = append(c.SubscribedHubs, hub)
}
func (c *Client) IsSubscribedToHub(hubToCheck *Hub) bool {
	for _, hub := range c.SubscribedHubs {
		if hub == hubToCheck {
			return true
		}
	}
	return false
}

// Enlever un hub des hubs abonnés du client
func (c *Client) RemoveSubscribedHub(hub *Hub) {
	for i, subscribedHub := range c.SubscribedHubs {
		if subscribedHub == hub {
			// Supprimer le hub de la liste des hubs abonnés
			c.SubscribedHubs = append(c.SubscribedHubs[:i], c.SubscribedHubs[i+1:]...)
			fmt.Println("Client unsubscribed from hub", c.SubscribedHubs)
			break
		}
	}
}
