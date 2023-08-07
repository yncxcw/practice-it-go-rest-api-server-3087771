package main

import (
	"example.com/backend"
)

func main() {
	app := &backend.App{}
	app.Port = ":9001"
	app.Init()
	app.Run()
}
