package main

import "chatgo/internal/client"

func main() {
	app := client.NewContainer().App()
	app.Run()
}
