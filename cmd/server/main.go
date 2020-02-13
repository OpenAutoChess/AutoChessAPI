package main

import (
    "log"
    "fmt"
    "net/http"
    socketio "github.com/googollee/go-socket.io"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		allowHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"

		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		w.Header().Set("Access-Control-Allow-Methods", "POST, PUT, PATCH, GET, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", allowHeaders)

		next.ServeHTTP(w, r)
	})
}

func main() {
    server, err := socketio.NewServer(nil)
    if (err != nil) {
        log.Fatal(err)
    }

    server.OnConnect("/", func(connection socketio.Conn) error {
        fmt.Println("connected:", connection.ID())
        connection.Join("game")
        return nil
    })

    server.OnEvent("/", "move", func(s socketio.Conn, msg string) {
		fmt.Println(msg)
		s.Emit("game:move", msg)
	})


    go server.Serve()
    defer server.Close()

    http.Handle("/socket.io/", corsMiddleware(server))
    http.Handle("/", http.FileServer(http.Dir("./asset")))

    log.Fatal(http.ListenAndServe(":8000", nil))


}
