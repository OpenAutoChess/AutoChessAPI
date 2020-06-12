package main

import (
	"encoding/json"
	"fmt"
	engineio "github.com/googollee/go-engine.io"
	"github.com/googollee/go-engine.io/transport"
	"github.com/googollee/go-engine.io/transport/polling"
	"github.com/googollee/go-engine.io/transport/websocket"
	socketio "github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowHeaders := "Access-Control-Allow-Headers, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, authorization"

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

		w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, PATCH, GET, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", allowHeaders)

		if r.Method == "OPTIONS" {
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	pt := polling.Default
	wt := websocket.Default

	wt.CheckOrigin = func(req *http.Request) bool {
		return true
	}

	server, err := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			pt,
			wt,
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	server.OnConnect("/", func(conn socketio.Conn) error {
		client, err := NewClient(conn)

		if err != nil {
			return err
		}
		conn.SetContext(client)
		return nil
	})

	games := map[string]*Game{}

	server.OnEvent("/", "join", func(conn socketio.Conn, room string) {
		if _, ok := games[room]; !ok {
			games[room] = NewGame(room, conn)
		}

		conn.LeaveAll()
		conn.Join(room)
		client, err := NewClient(conn)
		conn.SetContext(client)

		if err != nil {
			return
		}

		games[room].Join(client.User)

		if server.RoomLen("/", room) == getMode(room).cnt {
			var resp []byte
			resp, err := json.Marshal(games[room])
			if err == nil {
				server.BroadcastToRoom("/", room, "game-start", games[room], resp)
			}
		}
	})

	server.OnEvent("/", "move", func(conn socketio.Conn, data map[string]map[string]int) {
		user := conn.Context().(Client).User
		room := conn.Rooms()[0]
		game := games[room]

		if game.CheckMove(user, data["from"], data["to"]) {
			game.Move(data["from"], data["to"])

			var resp []byte
			resp, err := json.Marshal(game.Pieces)
			if err == nil {
				server.BroadcastToRoom("/", room, "game-move", string(resp))
			}
		}
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", corsMiddleware(server))
	log.Println("Serving at localhost:3002...")
	log.Fatal(http.ListenAndServe(":3002", nil))
}