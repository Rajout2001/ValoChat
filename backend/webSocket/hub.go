package webSocket

/*
	Hub to implement Join and left room in websockets
*/
// source: Youtube + Tabellout

// NewHub initialise un nouveau hub.
func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		broadcast:  make(chan Message),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
}

// Subscribe a client into our hub.
func (h *Hub) AddClient(client *Client) {
	h.clients[client] = true
	client.SubscribedHubs = append(client.SubscribedHubs, h)
}

// unsubscribe a client from our hub. 
func (h *Hub) RemoveClient(client *Client) {
	if _, ok := h.clients[client]; ok {
		delete(h.clients, client)
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.AddClient(client)
		case client := <-h.unregister:
			delete(h.clients, client)
			// close(client.send)
		case message := <-h.broadcast:
			for client := range h.clients {
				client.Emit(message)
			}
		}
	}
}