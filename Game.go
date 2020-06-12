package main

import (
	socketio "github.com/googollee/go-socket.io"
)


type Game struct {
	room string
	mode Mode
	Pieces map[int]map[int]Piece
	NextMove int
	Users []int
}

func NewGame(room string, conn socketio.Conn)*Game {
	mode := getMode(room)
	return &Game{room,mode, mode.getPieces(), 0, []int{}}
}

func (game *Game)Join(user User) {
	game.Users = append(game.Users, user.Id)
}

func (game *Game)CheckMove(user User, from map[string]int, to map[string]int) bool {
	if user.Id != game.Users[game.NextMove] {
		return false
	}

	return game.mode.checkMove(game, from ,to)
}

func (game *Game)Move(from map[string]int, to map[string]int) {
	game.NextMove = (game.NextMove + 1) % len(game.Users)

	piece := game.Pieces[from["row"]][from["col"]]
	delete(game.Pieces[to["row"]], to["col"])
	game.Pieces[to["row"]][to["col"]] = piece
	delete(game.Pieces[from["row"]], from["col"])
}