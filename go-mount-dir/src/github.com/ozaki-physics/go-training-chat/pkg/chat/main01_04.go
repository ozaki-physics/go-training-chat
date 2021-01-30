package chat

import (
	"flag"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"text/template"
	// 自作パッケージの import
	"github.com/ozaki-physics/go-training-chat/pkg/trace"
)

type templateHandler01_04 struct {
	filename string
	once     sync.Once
	templ    *template.Template
}

func (t *templateHandler01_04) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(template.ParseFiles(filepath.Join("pkg/chat/templates", t.filename)))
	})
	// テンプレートに書かれている置換用の変数を r の値を使って置換する
	t.templ.Execute(w, r)
}

func Main01_04() {
	// flag.String は *string 型を返す つまりフラグの値が保持されている場所を返す
	var addr = flag.String("addr", ":8080", "アプリケーションのアドレス")
	// フラグを解釈する コマンドラインで指定された文字列から情報を取り出して *addr に入れる
	flag.Parse()
	r := newRoom()
	// tracer パッケージを使うために room struct のフィールドを nil じゃなくするため
	// ターミナルに出力されるために os.Stdout を使う
	r.tracer = trace.New(os.Stdout)
	http.Handle("/", &templateHandler01_04{filename: "chat01_04.html"})
	http.Handle("/room", r)
	// バックグラウンドで動くため メインスレッドはサーバを起動していられる
	go r.run()
	// Web サーバの起動
	// フラグの値そのものを知りたいから 間接演算子`*`を使う
	log.Println("Webサーバーを開始します。ポート: ", *addr)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServer:", err)
	}

}
