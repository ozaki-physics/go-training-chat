## Appendix A
### Go のインストール
Go のインストールは、ソースコードからのコンパイルがオススメ
Go 1.5からセルフホスティング(コンパイラでそのコンパイラ自身のソースコードをコンパイルすることなど)ができるようになっているから?
バイナリをインストールしたら`export PATH=$PATH:/user/local/go/bin`に path を通すらしい
### Go の設定
ツールを呼び出しやすくするために 環境変数 PATH に go/bin を追加
path が通っているかの確認は`$ go version`

GOPATH
Go を使ったプログラムのソースコードやコンパイルされたバイナリのパッケージを置く場所を指定する
つまり Go は $GOPATH/src より下をパスとして認識するっぽい?
ソースコード内の import の記述があると GOPATH の場所を探す
`go get`コマンドでダウンロードされる場所も GOPATH になる
プロジェクトごとに GOPATH を変えることもできるが、すべてのプロジェクトで1つの GOPATH しか使わないことが推奨されている
書籍では`export GOPATH=$HOME/Work/go`にしていた
もし`go get github.com/stretchr/powerwalk`を実行すると
$GOPATH/src/github.com/stretchr/powerwalk というディレクトリ構造が生成されてダウンロードされる
他のパッケージと import path が被らないように github.com/自分のアカウント名/package名 などにしておくと良いらしい
調べると 最近プロジェクト毎に必要なパッケージは`glide get`を使うらしい? __調べる必要あり__
最近はプロジェクトを作るときは Go Modules がディファクトスタンダードらしい __調べる必要あり__

### Go のツール
コードフォーマットのツール`$ go fmt`は フォーマットしたいコードがあるディレクトリに移動して使うっぽい
静的解析ツール`go vet`
import 文を自動で修正してくれる外部パッケージ goimports がある
https://godoc.org/golang.org/x/tools/cmd/goimports
`go get golang.org/x/tools/cmd/goimports`
しばらく使う予定はない

## Appendix B


## パッケージ管理
そもそもコード内の import 文は異なるパッケージを利用するために使う
import文がパッケージを探すディレクトリ
- 標準パッケージを置くための $GOROOT
- ユーザが任意でダウンロードするパッケージ(準標準パッケージ(golang.org/x/なんとか), その他のサードパーティパッケージ)を配置するディレクトリ(GOPATHモード, モジュールモード)

Go のコードはライブラリも含めてすべて $GOPATH/src 以下に置くというルールがある
そのため、もともと Go に備わっていたパッケージを利用する方法は GOPATH を使って、どのプロジェクトでも同じパッケージ群を参照していた
しかし、パッケージをプロジェクトごとに管理したい要望が増えて GOPATH をプロジェクトごとにコントロールする GOVENDOR が出てきた
その GOVENDOR を使いやすくするサードパーティで glide や dep や goenv などが登場した
そしてバージョン1.13からは modules が標準で登場し GOVENDOR の代わりに使われるようになった
これで Go 言語の依存モジュール管理ツールは modules が主流になった
だが、いまでも Go 自体のバージョン管理には goenv が使われたりしているらしい

つまり、パッケージを管理する方法は
- GOPATHモードの $GOPATH/src に保存する方式
- モジュールモード(module-aware mode)の Go modulesを利用する方式

## Go modules
パッケージの依存関係やバージョンを保存を go.mod ファイルで管理する
Go modules によるパッケージ管理では go.mod と同じ階層に go.sum と呼ばれるファイルが作られる
go.sum はパッケージごとのチェックサムを記録したファイルで 2つともリポジトリにいれることが推奨されている

`$ go mod init`で go.mod ファイルが作られる
`$ go get`すると $GOPATH/src 配下ではなく $GOPATH/pkg/mod 配下にバージョンを含めて配置する
そして import も $GOPATH/pkg/mod から探す

GOPATH モードの go get では最新のパッケージを取得
モジュールモードでは go get golang.org/x/text@v0.3.1 のようにバージョンを指定してインストールできる

## Go modules の使い方
1. `go mod init`で、初期化する
引数にパッケージ名(example.com/go-mod-test)を書くことができる
書いたパッケージは`go.mod`ファイルに追記される
2. `go build`などのビルドコマンドで、依存モジュールを自動インストールする
`go.mod`ファイルに追記される
3. `go list -m all`で、現在の依存モジュールを表示する
4. `go get`で、依存モジュールの追加やバージョンアップを行う
5. `go mod tidy`で、使われていない依存モジュールを削除する

`go mod init (importpath)`で importpath を書かないと怒られた
importpath は世界で一意なものにする
コードが置かれる予定のリポジトリ を書けば良い
github.com って書く人が多いが go.modファイルは あなたのプロジェクトがGitHubにホストされていることを想定しているらしいし、一意になるから
sampleでも良いが公開しても使えない

go.modファイル のある配下の .goファイル から importpath に書いたパッケージが使えるようになる

GOPATHを設定している場合は`go mod`が使えないらしい
`$GOPATH/go.mod exists but should not`ってエラー言われた
`export GOPATH=`で GOPATH を削除したら問題なく`go mod`としての`go get`を使えた

## 雑多に調べたこと
src はソースコード(source code)の略
bin はバイナリ(binary)の略
`$ go build`とは、コンパイルすること
`$ go run`とは、コンパイルする かつ 実行すること
godoc.org とは、様々なGo言語のライブラリのAPIドキュメントを生成するサービス

どんなディレクトリ構成にするのがいい?

Docker で Go を使うなら何をマウントするかを考えたい
Dockerは環境だけ用意して、標準ライブラリはDockerの方で良いが、外部ライブラリはローカルの方に出したい
コンパイルしたコードは Docker 側で良い
