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
最初の1回目だけ
root@hoge:/go/src/github.com/ozaki-physics/go-training-chat# go mod init github.com/ozaki-physics/go-training-chat
root@hoge:/go/src/github.com/ozaki-physics/go-training-chat# 自由に使う
$ docker-compose down
```
`go get`が保存されないから、毎回実行する必要がある
`go.mod`に記述があれば大丈夫とかにならないかな

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
- CHAPTER 01 WebSocket を使ったチャットアプリケーション
- CHAPTER 02 認証機能の追加
- CHAPTER 03 プロフィール画像を追加する3つの方法
- CHAPTER 04 ドメイン名を検索するコマンドラインツール
- CHAPTER 05 分散システムと柔軟なデータの処理
- CHAPTER 06 REST形式でデータや機能を公開する
- CHAPTER 07 ランダムなおすすめを提示するWebサービス
- CHAPTER 08 ファイルシステムのバックアップ
- appendix A 安定した開発環境のためのベストプラクティス
- appendix B Goらしいコードの書き方
