package advanced

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

func SocketExample() {

	server, err := socketio.NewServer(nil)
	handleError(err)
	log.Println("1")
	server.OnConnect("/chatRoom", func(conn socketio.Conn) error {
		log.Println("on connection")
		conn.Join("chat")
		return nil
	})

	server.OnEvent("/chatRoom","messageReceived", func(s socketio.Conn, msg string) {
		fmt.Println("notice:", msg)
		server.BroadcastToRoom("chatRoom","messageReceived",msg)
	})

	server.OnDisconnect("/chatRoom", func(conn socketio.Conn,msg string) {
		log.Println("on disconnect")
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	log.Println("Serving at localhost:5000...")
	log.Fatal(http.ListenAndServe(":5000", nil))
}

func handleError(err error){
	if err != nil {
		log.Println(err.Error())
	}
}