package telestrator

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = 10 * time.Second
	maxMessageSize = 1024
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
		}
	}()

	o.conn.SetReadLimit(maxMessageSize)
	if err := o.conn.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Printf("[WARN] unable to set read deadline: %s", err)
	}
	o.conn.SetPongHandler(func(string) error {
		return o.conn.SetReadDeadline(time.Now().Add(pongWait))
	})

	for {
		_, message, err := o.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
				log.Printf("[ERROR] unexpected close error: %s", err)
			}
			break
		}

		o.overlay.OnMessageReceived(o, message)
	}
}

// WritePump writes outgoing socket data on the overlay connection.
func (o *Connection) WritePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		if o.conn != nil {
			o.conn.Close()
		}
	}()

	for {
		select {
		case <-ticker.C:
			if err := o.conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				log.Printf("[WARN] unable to set write deadline: %s", err)
			}

			if err := o.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				log.Printf("[ERROR] sending ping: %s", err)
				return
			}
		case message, ok := <-o.Send:
			if err := o.conn.SetWriteDeadline(time.Now().Add(writeWait)); err != nil {
				log.Printf("[WARN] unable to set write deadline: %s", err)
			}

			if !ok {
				// close connection
				if err := o.conn.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
					log.Printf("[WARN] unable to write message: %s", err)
				}
				return
			}

			w, err := o.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}

			if _, err := w.Write(message); err != nil {
				log.Printf("[WARN] unable to write message: %s", err)
			}

			if err := w.Close(); err != nil {
				return
			}
		}
	}
}
