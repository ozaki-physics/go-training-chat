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

コードが複数ある場合 package 全体を build してバイナリを生成する必要がある
`go build -o ファイル名`で build
`./ファイル名`で実行
ただし`_test.gp`や`_OS名.go`は build から除外される
`go install`にすると`$GOPATH/bin`にバイナリが生成される
build がうまくできないので、また後日ちゃんと調べる
→ package を正しく使うようにしたら build できた
ただ install はまだ使い方がよく分からない

go では同じパッケージ内に同じ名前の struct を type で定義することができないみたい

#### チャットルームとクライアントをサーバ側でモデル化する
オブジェクト指向のクラスを Go では 型
インスタンスは 型の値 に相当する
型 = クラスだと思っていたから違和感が無い笑
また インスタンスも従来`Hello hello = new Hello()`の hello のことを指しているけど 変数という認識だったから go でインスタンスは 型の値 つまり変数という認識で違和感が無い

チャットアプリの全ユーザー(クライアント)は 自動で大きな公開チャットルールに配置されるとする
*room 型は クライアントとの接続管理やメッセージのルーティングを受け持つ
*client 型は 1つのクライアントへの接続を表す

go で WebSocket を扱うライブラリは
golang.org/x/net/websocket
github.com/gorilla/websocket
があるらしい
net/websocket のドキュメントを読みに行ったら gorilla みたいに活発に保守されているパッケージに劣る的なことが書いてあった
また github.com/trevex/golem は軽量な Websocket フレームワーク
gorilla には無い イベントと関数のルーティング, JSON エンコード・デコード, ルーム機能, 接続型の拡張 が実装されている
gorilla を使うために`go get github.com/gorilla/websocket`する
→ go.mod や go.sum に書き加えられる

クライアントのモデル化は`func (c *client) read()`や`func (c *client) write()`で実現した
次はチャットルームのモデル化をするために クライアントがルームに参加 退室する仕組みを作る
同時アクセスによる競合を防ぐために 2つのチャネルでそれぞれ入室と退室を受け持つ
在室しているクライアントの保持のマップは 複数の goroutine からアクセスされても大丈夫なように チャネル経由で操作する
具体的には select 文を使う
select によって map への同時アクセスを防げる
閉じたチャネルに送信すると runtime panic を起こす

チャットルームを HTTP ハンドラにする
ServeHTTP メソッドを持たせたことで *room は HTTP ハンドラとして扱えるようになった
WebSocket を利用するためには websocket.Upgrader 型を使って HTTP 接続をアップグレードする必要がある
websocket.Upgrader 型の値(upgrader)は再利用できるため1個生成するだけで良い

HTTP リクエスト
→ ServeHTTP メソッドの呼び出し
→ upgrader.Upgrader メソッドの呼び出し
→ WebSocket コネクションの取得

ヘルパー関数を使って複雑さを下げる
現状 コードを使うためには room の struct を生成してもらう必要がある
生成を簡単にするためにヘルパー関数を使う

ブラウザで jQuery を使ってクライアントの作成

ポート番号8080が2ヵ所にハードコードされているから整える
テンプレートの埋め込み機能を使って JS が適切なホスト名とポートが分かるようにする

テンプレートの`{{`と`}}`で囲まれた部分はアノテーションを表していて データを埋め込む場所と値を示す
`{{.Host}}`なら`Template.Execute(w, data)`より`data.Host`の値で置換される
build や run するときに -addr=":8090"とかでポートを変えられる
でも今 main.go と chat の package main がズレているから ちゃんと渡すということをしないと使えない
とりあえず初期値の8080で困らないし放置する

### 以下は内容から逸れる話
#### Go Modules の使い方
`go mod init パス`で初期設定
`go build`で勝手に go.mod に書かれたパッケージもダウンロードして build してくれる
`go mod download`で go.mod や go.sum をもとに依存パッケージをダウンロードしてくれるらしい
__build も成功するが ダウンロードした package がどこにも見当たらない__
`$GOPATH/pkg/mod`の配下に置かれるらしい
#### Go にインスタンスという概念はあるのか
変数に type で宣言した struct を入れることが実質インスタンスを作っているのと同義
クラスに属する操作(通称メソッド)は レシーバを使って struct に属させるメソッドと同じ
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
#### `http.FileServer()`を調べる
`Handler 型`を返す関数 つまり`ServeHTTP()`関数だけを持ち HTTP リクエストを受けてレスポンスを返す
引数に`FileSystem 型`を渡す
#### `http.Dir()`を調べる
`Dir 型`があある
限定的なディレクトリツリーの中でホストOSのファイルシステムを使用して`FileSystem 型`を返す
#### FileSystem を調べる
`Opne()`インタフェースを持つ
```go
type FileSystem interface {
    Open(name string) (File, error)
}
```
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
#### `flag.String()`を調べる
```go
func String(name string, value string, usage string) *string
```
name: 指定する名前, value: デフォルトの値, usage: 使用方法などの説明 を使って文字列フラグを定義する
戻り値は フラグの値を格納する文字列変数のアドレス
あくまでアドレスだから 指定する名前の中の値が欲しいときは`*`を使う必要がある
#### `template.Template.Execute()`を調べる
```go
func (t *Template) Execute(wr io.Writer, data interface{}) error
```
data を wr に適用する
つまりテンプレートの中の変数を置換してくれる
