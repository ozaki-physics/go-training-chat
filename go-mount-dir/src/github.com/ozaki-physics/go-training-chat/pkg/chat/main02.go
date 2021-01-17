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
func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 1度だけ実行する処理
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
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
