package chat

type room struct {
	// すべてのクライアントに転送するためのメッセージを保持するチャネル
	forward chan []byte
}
