# go-training-chat
書籍を読みながらチャットアプリの作成  

## 目的 Overview
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
1. 書籍に取り組む  
本当なら feature/内容 ブランチを作って develop ブランチにマージして開発する。  
面倒だから develop ブランチに直接コミットして 区切りが良いと main ブランチにマージする  

## 参考文献 References
『Go言語によるWebアプリケーション開発』  
原書名『Go Programming Blueprints』  
著者:Mat Ryer  
訳者:鵜飼 文敏, 牧野 聡  
出版社:株式会社オライリー・ジャパン  
2016年01月21日：初版第01刷  
2016年02月12日：初版第02刷  
ISBN:978-4-87311-752-2 C3055  

サンプルコードのリポジトリ  
https://github.com/oreilly-japan/go-programming-blueprints  

## 目次 Table of contents
- CHAPTER 01 WebSocket を使ったチャットアプリケーション  
net/http パッケージを使った html 送信や WebSocket を使ってブラウザと接続する方法  
[メモ](./memo/chapter01.md)  
- CHAPTER 02 認証機能の追加  
OAuth 認証を利用したユーザーの識別とソーシャルログインの実装  
[メモ](./memo/chapter02.md)  
- CHAPTER 03 プロフィール画像を追加する3つの方法  
画像をユーザーがアップロード, 認証サービス, Webサービス Gravatar から取得  
[メモ](./memo/chapter03.md)  
- CHAPTER 04 ドメイン名を検索するコマンドラインツール  
標準入出力とパイプの解説  
[メモ](./memo/chap04r01.md)  
- CHAPTER 05 分散システムと柔軟なデータの処理  
NSQ, MongoDB の利用  
[メモ](./memo05hapter01.md)  
- CHAPTER 06 REST形式でデータや機能を公開する  
5章の内容をJSON形式にして公開  
[メモ](./memo/chapter06.md)  
http.HandlerFunc の機能をラップしてパイプライン形式実装
- CHAPTER 07 ランダムなおすすめを提示するWebサービス  
内部データを適切に公開する方法, Go で列挙型を実装する方法  
[メモ](./memo/chapter07.md)  
- CHAPTER 08 ファイルシステムのバックアップ  
os パッケージを使ったファイルシステムの操作, Go のインタフェースの解説  
[メモ](./memo/chapter08.md)  
- appendix A 安定した開発環境のためのベストプラクティス  
環境構築の方法など  
- appendix B Goらしいコードの書き方  
日本語版オリジナル記事で Go 言語のイディオムの解説  
