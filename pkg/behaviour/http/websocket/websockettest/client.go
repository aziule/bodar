package websockettest

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// Client is a test struct to test the websocket behaviours.
type Client struct {
	conn  *websocket.Conn
	chErr chan error // Send errors here
}

// NewClient creates a new test client.
func NewClient(url string) (*Client, chan []byte, chan error, error) {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("dial error: %v", err)
	}

	chReadMsg := make(chan []byte)
	chErr := make(chan error)

	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				chErr <- err
				continue
			}

			chReadMsg <- msg
		}
	}()

	return &Client{
		conn:  conn,
		chErr: chErr,
	}, chReadMsg, chErr, nil
}

// WriteMessage writes a message to the server.
func (c *Client) WriteMessage(payload []byte) {
	c.writeMessage(websocket.BinaryMessage, payload)
}

// Ping the server.
func (c *Client) Ping() {
	c.writeMessage(websocket.PingMessage, nil)
}

// Pong the server.
func (c *Client) Pong() {
	c.writeMessage(websocket.PongMessage, nil)
}

// Close the connection with the server.
func (c *Client) Close() {
	c.writeMessage(websocket.CloseMessage, nil)
}

func (c *Client) writeMessage(messageType int, data []byte) {
	err := c.conn.WriteMessage(messageType, data)
	if err != nil {
		c.chErr <- err
	}
}
