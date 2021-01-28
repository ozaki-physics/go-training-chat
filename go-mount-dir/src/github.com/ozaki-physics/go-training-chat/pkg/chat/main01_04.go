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
