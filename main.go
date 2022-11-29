package main

import "prts-backend/src/server"

func main() {
	var port = ":33333"
	server.Run(&port, nil)
}
