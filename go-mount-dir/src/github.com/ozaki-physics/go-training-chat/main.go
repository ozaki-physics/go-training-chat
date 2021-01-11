package main

import (
	"fmt"
	"github.com/ozaki-physics/go-training-chat/training"
	"github.com/ozaki-physics/go-training-chat/chat_app"
	// "net/http"
)

func main() {
	fmt.Println("hello")
	fmt.Println(training.Message)
	fmt.Println(chat_app.Message)


	// http.Handle("/", http.FileServer(http.Dir("web")))
	// http.ListenAndServe(":8080", nil)
}
