package chat

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	// 自作パッケージの import
	"github.com/ozaki-physics/go-training-chat/pkg/trace"
)

type room struct {
	// すべてのクライアントに転送するためのメッセージを保持するチャネル
	forward chan []byte
	// 参加しようとしているクライアントのためのチャネル
	join chan *client
	// 退室しようとしているクライアントのためのチャネル
	leave chan *client
	// 在室しているクライアントの保持
	// 複数の goroutine が同時に変更する可能性があるため チャネル経由で操作する
	clients map[*client]bool
	// チャットルーム上の操作ログを受け取る
	tracer trace.Tracer
	// tracer trace.Off(),
}

// ヘルパー関数を使って複雑さを下げる
func newRoom() *room {
	return &room{
		forward: make(chan []byte),
		join:    make(chan *client),
		leave:   make(chan *client),
		clients: make(map[*client]bool),
	}
}

func (r *room) run() {
	// 無限ループするが goroutine はバックグラウンドで実行されるため問題ない
	for {
		// select によって r.clients への同時アクセスを防げる
		select {
		case client := <-r.join:
			// 参加
			r.clients[client] = true
			r.tracer.Trace("新しいクライアントが参加しました")
		case client := <-r.leave:
			// 退室
			delete(r.clients, client)
			r.tracer.Trace("クライアントが退室しました")
			// client.send チャネルを close しているのは client の write メソッドの for ループを終了させるため
			close(client.send)
		case msg := <-r.forward:
			r.tracer.Trace("メッセージを受信しました:", string(msg))
			// すべてのクライアントにメッセージを転送
			for client := range r.clients {
				select {
				case client.send <- msg:
					// メッセージを送信
					r.tracer.Trace("-- クライアントに送信されました")
				default:
					// 送信失敗のとき ルームから削除するなどの掃除をする
					// この default はclient.send チャネルに msg が送信できなかったときに動作する
					// しかし毎回 client.send チャネルを close するから閉じた send チャネルを持ったクライアントは r.clients にいない
					// にもかかわらず delete を実行していう
					// よってこのメソッドの組み方は参考にならないかも
					delete(r.clients, client)
					close(client.send)
					r.tracer.Trace("-- 送信に失敗しました クライアントをクリーンアップします")
				}
			}
		}
	}
}

const (
	soketBufferSize   = 1024
	messageBufferSize = 256
)

// WebSocket を利用するためには websocket.Upgrader 型を使って HTTP 接続をアップグレードする必要がある
// websocket.Upgrader 型の値(upgrader)は再利用できるため1個生成するだけで良い
var upgrader = &websocket.Upgrader{
	ReadBufferSize:  soketBufferSize,
	WriteBufferSize: soketBufferSize,
}

// HTTP ハンドラとして扱えるように *room に ServerHTTP()メソッドを実装
func (r *room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP:", err)
		return
	}
	// クライアントの struct 生成
	client := &client{
		socket: socket,
		send:   make(chan []byte, messageBufferSize),
		room:   r,
	}
	// room に client を参加させる
	r.join <- client
	// defer は defer へ渡した関数の実行を 呼び出し元の関数の終わりまで遅延させる
	defer func() { r.leave <- client }()
	// goroutine で呼び出され続ける
	go client.write()
	// read で接続が保持され 終了まで他の処理はブロックされる
	client.read()
}
