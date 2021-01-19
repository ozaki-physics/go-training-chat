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

### 内容
Go アプリケーションや Go の標準ライブラリは パッケージ単位で構成されている
各パッケージはそれぞれのフォルダに配置される

Java や Node.js では同期状態を保つのは複雑なスレッド管理がいる
Go では並列処理向けの仕組みが言語に実装されているため チャネルなどを利用して容易に実装できる

Webサーバの2個の役割
1. html, css, JavaScript を返す
2. WebSocket 通信

```go:main01.go
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

コード内に html を直接書くこともできるが 拡張性が無い
テンプレートを使うと 汎用的なテキストの中に 固有のテキストを埋め込むことができる
そのためテンプレートを利用するのが一般的
Go 標準パッケージに テキスト向けの`text/template`と html 向けの`html/template`がある
`html/template`だと挿入するコンテキストに不正なスクリプトの確認や URLで使用できない文字をエンコードすることが可能

テンプレートを使う場合はテンプレートのコンパイルが必要
テンプレートのコンパイルとは テンプレートを解釈してデータを埋め込める状態にすること

`sync.Once.Do()`を使うことで 複数の goroutine から `ServeHTTP()` メソッドを呼び出されてもコンパイルは1回しか実行しないことを保証する
`ServeHTTP()`メソッドの中でテンプレートをコンパイルすると 必要になるまで処理を後回しにできる
このことを 遅延初期化(lazy initialization)という。
めったに呼ばれない処理の中で遅延初期化が使われているとエラーに気づかないという問題もある

### 以下は内容から逸れる話
#### sync.Once を調べる
sync パッケージは 排他制御の Mutex でも使った並列処理を簡単に扱うためのパッケージ
`sync.Once`は関数を1度だけ呼び出したいときに使う
並列で走らせても1度だけしか実行されないから初期化等で使う
```go
import (
  "fmt"
  "sync"
)

func main() {
    var once sync.Once

    once.Do(func() {fmt.Println("A")})
    once.Do(func() {fmt.Println("B")}) // こちらの関数は呼び出されない
}
```
`once.Do(func() {処理})`の形で使うことが多いみたい
公式(https://golang.org/pkg/sync/#example_Once)より
>`once.Do(f)` が複数回呼び出された場合 f の値が呼び出しごとに異なっていても、最初の呼び出しのみがfを呼び出します。各関数を実行するには、Onceの新しいインスタンスが必要です。
>Do への呼び出しはfへの1回の呼び出しが戻るまで戻らないため f によって Do が呼び出されると、デッドロックが発生します。

#### `http.Handle()`を調べる
その前に`http.Handler`から理解した方が良い
`http.Handler`とは `ServeHTTP()`関数だけを持つインタフェース
HTTP リクエストを受けてレスポンスを返すことが責務
```go
type Handler interface {
  ServeHTTP(ResponseWriter, *Request)
}
```

本命の`http.Handle()`とは 表示する URL と URL に対応する`http.Handler`を`DefaultServeMux`に登録する関数
`DefaultServeMux`とはデフォルトで`http.ServeMux 型`の構造体(https://golang.org/pkg/net/http/#ServeMux)
4個のメソッドを持つ
1. `func NewServeMux() *ServeMux`
2. `func (mux *ServeMux) Handle(pattern string, handler Handler)`
3. `func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request))`
4. `func (mux *ServeMux) Handler(r *Request) (h Handler, pattern string)`

`http.ServeMux 型`の構造体は`http.Handler`インタフェースを持つ
つまり`ServeHTTP`メソッドを持ち、HTTP リクエストを受けてレスポンスを返す
レスポンスを返すときに URL に対応した`http.Handler`を実行する
つまりルータの役割を担う

`http.ListenAndServe()`の第2引数が nil の場合`DefaultServeMux`が Handler として指定される

#### `http.HandleFunc()`をついでに調べる
その前に`http.HandlerFunc`から理解した方が良い
`http.HandlerFunc 型`とは
`ServeHTTP()`関数を持つ`func(ResponseWriter, *Request)`の別名の型
関数を定義して`http.HandlerFunc 型`にキャストするだけで構造体を宣言することなく `http.Handler`を用意することができる
つまり struct を宣言しなくても HTTP リクエストを受けてレスポンスを返すことができる

本命の`http.HandleFunc()`とは 
URL と`func(ResponseWriter, *Request)`を引数で渡すと`DefaultServeMux`に登録してくれる関数
```go
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
    DefaultServeMux.HandleFunc(pattern, handler)
}
```
内部で`func(ResponseWriter, *Request)`から`http.HandlerFunc 型`へのキャストが行われる

#### `Template.Execute()`を調べる
公式 https://golang.org/pkg/text/template/#Template.Execute を読む
```go
func (t *Template) Execute(wr io.Writer, data interface{}) error
```
Execute は解析されたテンプレートに 指定されたデータオブジェクトを適用し 出力を wr に書き込む

#### `template.Must()`を調べる
関数呼び出しをラップして error が nil じゃないなら panic を起こすペルパー関数
var t = template.Must(template.New("name").Parse("text"))

#### `template.ParseFiles()`を調べる
名前付きファイルからテンプレートを解析して新しいテンプレートを作成する
引数にファイルの名前が必要
