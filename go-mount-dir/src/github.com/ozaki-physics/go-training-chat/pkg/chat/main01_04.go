package chat

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
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
	t.templ.Execute(w, r)
}

func Main01_04() {
	r := newRoom()
	http.Handle("/", &templateHandler01_04{filename: "chat01_04.html"})
	http.Handle("/room", r)
	// バックグラウンドで動くため メインスレッドはサーバを起動していられる
	go r.run()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServer:", err)
	}

}
