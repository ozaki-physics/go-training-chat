package chat

import (
	"github.com/gorilla/websocket"
)

// client はチャットを行っている一人のユーザーを表す
type client struct {
	// soket はこのクライアントとの通信を行う WebSocket への参照
	socket *websocket.Conn
	// send はメッセージが送られるバッファ付きチャネル
	// 待ち行列のように蓄積され WebSocket を通じてユーザーのブラウザに送られるのを待機
	send chan []byte
	// room はこのクライアントが参加しているチャットルームの参照を保持
	room *room
}

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
