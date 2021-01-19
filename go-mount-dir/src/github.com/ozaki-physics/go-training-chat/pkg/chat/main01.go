package main

import (
	"log"
	"net/http"
)

func main() {
	// http.HandleFunc(a, b)
	// a は path, b は path のリクエストがきたときに実行する関数
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
			<html>
				<head>
					<title>チャット</title>
				</head>
				<body>
					チャットしましょう!直接版
				</body>
			</html>
		`))
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
