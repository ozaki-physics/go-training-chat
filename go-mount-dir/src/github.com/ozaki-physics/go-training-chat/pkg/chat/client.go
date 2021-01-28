package chat

import (
	"github.com/gorilla/websocket"
)

// クライアントが WebSocket から ReadMessage を使ってデータを読み込む処理をする
func (c *client) read() {
	for {
		// 読み込んだ msg はすぐ forward チャネルに送られる
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

// 継続的に send チャネルから値を受け取り WriteMessage メソッドで書き出す
func (c *client) write() {
	// チャネルに溜まっているだけ繰り返す
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
