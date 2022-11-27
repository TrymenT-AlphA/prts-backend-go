package main

import "prts-backend/src/server"

func main() {
	var port = ":3000"
	server.Run(&port, nil)
}
