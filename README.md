# go-training-chat
Go言語の基本構文勉強とチャットアプリの作成

## 目的 Overview
Go言語の基本構文を勉強する<br>
書籍を見ながらWebアプリを作ってみる

## インストール方法 Install
```bash
$ git clone git@github.com:ozaki-physics/go-training-chat.git
```

## 環境 Requirement
```bash
$ docker-compose up -d
$ docker-compose exec go_training_chat bash
root@hoge:/go/src/github.com/ozaki-physics/go-training-chat# 自由に使う
$ docker-compose down
```
`go mod download`で依存ライブラリをダウンロードする
なるべく推奨されている感じのディレクトリ構成にしたが GOPATH との共存がうまくいっていないからそのままでは使えない

## 使い方 Usage
1. 公式で基本構文を学習する<br>
https://go-tour-jp.appspot.com/list

2. 書籍に取り組む<br>
本当なら feature/内容 ブランチを作って develop ブランチにマージして開発する。<br>
面倒だから develop ブランチに直接コミットして 区切りが良いと main ブランチにマージする

## 参考文献 References
『Go言語によるWebアプリケーション開発』<br>
原書名『Go Programming Blueprints』<br>
著者:Mat Ryer<br>
訳者:鵜飼 文敏, 牧野 聡<br>
出版社:株式会社オライリー・ジャパン<br>
2016年01月21日：初版第01刷<br>
2016年02月12日：初版第02刷<br>
ISBN:978-4-87311-752-2 C3055

サンプルコードのリポジトリ<br>
https://github.com/oreilly-japan/go-programming-blueprints

## 目次 Table of contents
- <b>CHAPTER 01 WebSocket を使ったチャットアプリケーション</b>
net/http パッケージを使った html 送信や WebSocket を使ってブラウザと接続する方法
- <b>CHAPTER 02 認証機能の追加</b>
OAuth 認証を利用したユーザーの識別とソーシャルログインの実装
- <b>CHAPTER 03 プロフィール画像を追加する3つの方法</b>
画像をユーザーがアップロード, 認証サービス, Webサービス Gravatar から取得
- <b>CHAPTER 04 ドメイン名を検索するコマンドラインツール</b>
標準入出力とパイプの解説
- <b>CHAPTER 05 分散システムと柔軟なデータの処理</b>
NSQ, MongoDB の利用
- <b>CHAPTER 06 REST形式でデータや機能を公開する</b>
5章の内容をJSON形式にして公開
http.HandlerFunc の機能をラップしてパイプライン形式実装
- <b>CHAPTER 07 ランダムなおすすめを提示するWebサービス</b>
内部データを適切に公開する方法, Go で列挙型を実装する方法
- <b>CHAPTER 08 ファイルシステムのバックアップ</b>
os パッケージを使ったファイルシステムの操作, Go のインタフェースの解説
- <b>appendix A 安定した開発環境のためのベストプラクティス</b>
環境構築の方法など
- <b>appendix B Goらしいコードの書き方</b>
日本語版オリジナル記事で Go 言語のイディオムの解説
