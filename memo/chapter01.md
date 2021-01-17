## 第01章 WebSocket を使ったチャットアプリケーション
筆者は機能するコードが作れたらそれを最初から作り直すようにしている
多くの小説家やジャーナリストは「物書きの真髄は書き直すことにある」と述べており、ソフトウェアにも当てはまると思われるから
初めて書くときは 問題や解決アプローチを学び、頭で考えていることを書き出すだけで精一杯
同じプログラムをもう一度作ろうと思ったときに 問題解決のための新しい知識を適用できるようになる
### この章でやること
- net/http パッケージを使った http リクエストの応答
- テンプレートを使ったコンテンツを返す
- http.Handler 型のインタフェースに適合させる
- goroutine を使ってアプリケーションが複数の作業を同時に行う
- チャネルを使った goroutine 間の情報共有
- WebSocket の実装
- ログの記録
- テスト駆動開発で 完全なパッケージ構成を作成
- 非公開の型を 公開されているインタフェースで返す

Go アプリケーションや Go の標準ライブラリは パッケージ単位で構成されている
各パッケージはそれぞれのフォルダに配置される

Java や Node.js では同期状態を保つのは複雑なスレッド管理がいる
Go では並列処理向けの仕組みが言語に実装されているため チャネルなどを利用して容易に実装できる

Webサーバの2個の役割
1. html, css, JavaScript を返す
2. WebSocket 通信

```go:main.go
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
					チャットしましょう!
				</body>
			</html>
		`))
	})
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServer:", err)
	}
}
```
