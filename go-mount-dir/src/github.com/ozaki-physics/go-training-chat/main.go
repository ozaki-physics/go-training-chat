package main

import (
	"fmt"
	"github.com/ozaki-physics/go-training-chat/pkg/training"
	"github.com/ozaki-physics/go-training-chat/pkg/chat_app"
	// "net/http"
)

func main() {
	fmt.Println("hello")
	fmt.Println(training.Message)
	fmt.Println(chat_app.Message)
	// 動かない
	// fmt.Println(training.message)
	// fmt.Println(chat_app.message)


	// http.Handle("/", http.FileServer(http.Dir("web")))
	// http.ListenAndServe(":8080", nil)
}