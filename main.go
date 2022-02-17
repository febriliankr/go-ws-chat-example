package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	server := socketio.NewServer(nil)

	server.OnConnect("/socket-io", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("WS connected!", s.ID())
		return nil
	})

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("WS connected!", s.ID())
		return nil
	})

	server.OnEvent("/socket-io", "chat message", func(s socketio.Conn, msg string) {
		fmt.Println(msg)
		s.Emit("chat message", "chat: "+msg)
	})
	server.OnEvent("/", "chat message", func(s socketio.Conn, msg string) {
		fmt.Println(msg)
		s.Emit("chat message", "chat: "+msg)
	})

	go server.Serve()

	http.Handle("/socket-io/", server)
	http.Handle("/", http.FileServer(http.Dir("./")))
	log.Println("Serving at http://localhost:3000 ...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
