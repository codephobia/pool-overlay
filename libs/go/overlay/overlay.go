package overlay

import (
	"sync"
)

type BroadcastExcept struct {
	Connection *Connection
	Message    []byte
}

type OnMessageReceivedFunc func(*Overlay, *Connection, []byte)

// Overlay managed multiple websocket connections to the overlay.
type Overlay struct {
	Connections map[*Connection]struct{}

	Broadcast       chan []byte
	BroadcastExcept chan BroadcastExcept
	Register        chan *Connection
	Unregister      chan *Connection

	OnMessageReceived func(*Overlay, *Connection, []byte)

	lock sync.Mutex
}

// NewOverlay returns a new Overlay.
func NewOverlay() *Overlay {
	return &Overlay{
		Connections: make(map[*Connection]struct{}),

		Broadcast:       make(chan []byte),
		BroadcastExcept: make(chan BroadcastExcept),
		Register:        make(chan *Connection),
		Unregister:      make(chan *Connection),

		lock: sync.Mutex{},
	}
}

func (o *Overlay) AssignOnMessageReceived(onMessageReceived OnMessageReceivedFunc) {
	o.OnMessageReceived = onMessageReceived
}

// Init starts an Overlay. Should be run in a separate go routine.
func (o *Overlay) Init() {
	for {
		select {
		// register connection with the overlay
		case connection := <-o.Register:
			o.RegisterConnection(connection)

		// unregister connection with the overlay
		case connection := <-o.Unregister:
			o.UnregisterConnection(connection)

		// broadcast event to all connections
		case message := <-o.Broadcast:
			o.SendConnections(message)

		case except := <-o.BroadcastExcept:
			o.SendConnectionsExcept(except.Connection, except.Message)
		}
	}
}

// RegisterConnection registers a connection with the overlay.
func (o *Overlay) RegisterConnection(connection *Connection) {
	o.lock.Lock()
	o.Connections[connection] = struct{}{}
	o.lock.Unlock()
}

// UnregisterConnection unregisters a connection with the overlay.
func (o *Overlay) UnregisterConnection(connection *Connection) {
	o.lock.Lock()
	o.ConnectionClose(connection)
	o.lock.Unlock()
}

// SendConnections sends all connections an event.
func (o *Overlay) SendConnections(message []byte) {
	o.lock.Lock()
	defer o.lock.Unlock()

	for connection := range o.Connections {
		select {
		case connection.Send <- message:
		default:
			o.ConnectionClose(connection)
		}
	}
}

// SendConnectionsExcept sends all connections except one an event.
func (o *Overlay) SendConnectionsExcept(except *Connection, message []byte) {
	o.lock.Lock()
	defer o.lock.Unlock()

	for connection := range o.Connections {
		if connection == except {
			continue
		}

		select {
		case connection.Send <- message:
		default:
			o.ConnectionClose(connection)
		}
	}
}

// ConnectionClose closes a connection with the overlay.
func (o *Overlay) ConnectionClose(connection *Connection) {
	// make sure the connection still exists with the overlay
	if _, ok := o.Connections[connection]; !ok {
		return
	}

	// close connection channel
	close(connection.Send)

	// remove connection from overlay
	delete(o.Connections, connection)
}
