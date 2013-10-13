package lib

import (
    "code.google.com/p/go.net/websocket"
)

type Client struct {
    url string
    origin string
    conn *websocket.Conn
}

func NewClient(url, origin string) *Client {
    return &Client{url, origin, nil}
}

func (self *Client) Connect() (err error) {
    self.conn, err = websocket.Dial(self.url, "", self.origin)
    return
}

func (self *Client) Send(message []byte) (err error) {
    _, err = self.conn.Write(message)
    return
}
