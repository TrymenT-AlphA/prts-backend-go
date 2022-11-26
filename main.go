package main

import "prts-backend/src/server"

func main() {
	var port = ":3001"
	server.Run(&port, nil)
}
