package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	// ファイル名の格納
	filename string
	// コンパイルするために使う
	once sync.Once
	// templ コンパイルされたテンプレートの参照を保持
	templ *template.Template
}

// ServerHTTP は HTTP リクエストを処理する *templateHandler 型のレシーバを持つメソッド
// sync.Once の値は常に同じものを使う必要があるため レシーバがポインタである必要がある
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 1度だけ実行する処理
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	// テンプレートに はめ込むデータ(今回は nil)を適用する
	// t.templ.Execute(w, nil)
	// 訳者の注意書きで 戻り値をチェックすべきらしいので実装する
	// しかし うまくエラーを起こさせる方法が分からなかった
	if err := t.templ.Execute(w, nil); err != nil {
		log.Fatal("テンプレートに適用するとき", err)
	}
}

func main() {
	// 第2引数には Handler 型じゃなくても ServeHTTP() メソッドを持っている struct なら良い
	// ここで URL に対応する http.Handler を DefaultServeMux に登録
	http.Handle("/", &templateHandler{filename: "chat.html"})
	// Web サーバを開始
	// ListenAndServe() の第2引数が nil なら DefaultServeMux が Handler として指定される
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
