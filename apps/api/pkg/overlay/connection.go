package overlay

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 1 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = 10 * time.Second
	maxMessageSize = 512
)

// Connection is a websocket connection for the OBS Overlay.
type Connection struct {
	overlay *Overlay

	conn *websocket.Conn
	Send chan []byte
}

// NewOverlayConnection returns a new overlay connection.
func NewOverlayConnection(overlay *Overlay, conn *websocket.Conn) *Connection {
	return &Connection{
		overlay: overlay,

		conn: conn,
		Send: make(chan []byte, 256),
	}
}

// ReadPump reads incoming socket data on the overlay connection.
func (o *Connection) ReadPump() {
	defer func() {
		o.overlay.Unregister <- o
		if o.conn != nil {
			o.conn.Close()
			o.conn = nil
		}
	}()

	o.conn.SetReadLimit(maxMessageSize)
	o.conn.SetReadDeadline(time.Now().Add(pongWait))
	o.conn.SetPongHandler(func(string) error { return o.conn.SetReadDeadline(time.Now().Add(pongWait)) })

	for {
		_, _, err := o.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				log.Printf("[ERROR] unexpected close error: %s", err)
			}
			break
		}
	}
}

// WritePump writes outgoing socket data on the overlay connection.
func (o *Connection) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		if o.conn != nil {
			o.conn.Close()
			o.conn = nil
		}
	}()

	for {
		select {
		case <-ticker.C:
			o.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := o.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				log.Printf("[ERROR] sending ping: %s", err)
				return
			}
		case message, ok := <-o.Send:
			o.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// close connection
				o.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := o.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			w.Write(message)

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}
