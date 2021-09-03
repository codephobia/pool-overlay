package telestrator

import (
	"sync"
)

// BroadcastExcept contains a payload and a connection to skip during broadcast.
type BroadcastExcept struct {
	Connection *Connection
	Message    []byte
}

// Overlay managed multiple websocket connections to the overlay.
type Overlay struct {
	Connections map[*Connection]struct{}

	BroadcastExcept chan BroadcastExcept
	Register        chan *Connection
	Unregister      chan *Connection

	lock sync.Mutex
}

// NewOverlay returns a new Overlay.
func NewOverlay() *Overlay {
	return &Overlay{
		Connections: make(map[*Connection]struct{}),

		BroadcastExcept: make(chan BroadcastExcept),
		Register:        make(chan *Connection),
		Unregister:      make(chan *Connection),

		lock: sync.Mutex{},
	}
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

// OnMessageReceived takes an incoming message and broadcasts it to everyone
// else connected.
func (o *Overlay) OnMessageReceived(connection *Connection, message []byte) {
	// broadcast the message out to all other connections
	o.BroadcastExcept <- BroadcastExcept{
		Connection: connection,
		Message:    message,
	}
}
