package main

import "chatgo/internal/chatgo"

func main() {
	app := chatgo.NewContainer().App()
	app.Run()
}
