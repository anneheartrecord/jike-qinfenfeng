package models

import (
	"github.com/gorilla/websocket"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

const month = 60 * 60 * 24 * 30

type Rooms struct {
	gorm.Model
	RoomID   int64  `gorm:"unique" `
	RoomName string `gorm:"unique"`
	RoomCap  int    ``
	RoomHost string ``
}

type SendMsg struct {
	Type    int    `json:"type"`
	Content string `json:"content"`
}

type ReplyMsg struct {
	From    string `json:"from"`
	Code    int    `json:"code"`
	Content string `json:"content"`
}
type Client struct {
	ID     string
	SendID string
	Socket *websocket.Conn
	Send   chan []byte
}
type Broadcast struct {
	Client  *Client
	Message []byte
	Type    int
}
type Message struct {
	Sender    string `json:"sender,omitempty"`
	Recipient string `json:"recipient,omitempty"`
	Content   string `json:"content,omitempty"`
}

type ClientManager struct {
	Clients    map[string]*Client
	Broadcast  chan *Broadcast
	Reply      chan *Client
	Register   chan *Client
	Unregister chan *Client
}

var Manager = ClientManager{
	Clients:    make(map[string]*Client), // 参与连接的用户，出于性能的考虑，需要设置最大连接数
	Broadcast:  make(chan *Broadcast),
	Register:   make(chan *Client),
	Reply:      make(chan *Client),
	Unregister: make(chan *Client),
}

func (c *Client) Read() {
	defer func() {
		Manager.Unregister <- c
		_ = c.Socket.Close()
	}()
	for {
		c.Socket.PongHandler()
		sendMsg := new(SendMsg)
		err := c.Socket.ReadJSON(&sendMsg)
		if err != nil {
			zap.L().Error("read json failed", zap.Error(err))
			Manager.Unregister <- c
			_ = c.Socket.Close()
			break
		}

	}
}

/*func(c *Client) Write()  {
	for id, conn := range Manager.Clients {
		if id != "1" {
			continue
		}
		select {
		case conn.Send <- message:
			flag = true
		default:
			close(conn.Send)
			delete(Manager.Clients, conn.ID)
		}
	}

}*/
