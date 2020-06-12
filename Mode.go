package main

import (
	"fmt"
	"strings"
)

type Mode struct {
	name string
	cnt int
	getPieces func()map[int]map[int]Piece
	checkMove func(*Game, map[string]int, map[string]int)bool
}



var modes = map[string]Mode{
	"classical": Mode{"classical",2, piecesClassical, checkMoveClassical},
	"blitz": Mode{"blitz",2,piecesClassical,checkMoveClassical},
	"four": Mode{"four",4, piecesFour,checkMoveFour},
	"puzzles": Mode{"puzzles",1, piecesFour,checkMoveClassical},
	"computer": Mode{"computer",1, piecesClassical,checkMoveClassical},
}

func getMode(room string)Mode {
	return modes[strings.Split(room, "_")[0]]
}

func checkMoveClassical(game *Game, from map[string]int, to map[string]int) bool {
	piece := game.Pieces[from["row"]][from["col"]]
	var moves []map[string]int
	if piece.Class == "Pawn" {
		moves = getClassicalPawnMoves(game, from["row"], from["col"])
	} else if piece.Class == "King" {
		moves = getClassicalKingMoves(game, from["row"], from["col"])
	} else if piece.Class == "Rook" {
		moves = getClassicalRookMoves(game, from["row"], from["col"])
	} else if piece.Class == "Bishop" {
		moves = getClassicalBishopMoves(game, from["row"], from["col"])
	} else if piece.Class == "Queen" {
		moves = getClassicalQueenMoves(game, from["row"], from["col"])
	} else {
		moves = getClassicalKnightMoves(game, from["row"], from["col"])
	}

	fmt.Println(moves)
	moveExists := false
	for _, move := range moves {
		if move["row"] == to["row"] && move["col"] == to["col"] {
			moveExists = true
			break
		}
	}
	if moveExists {
		fmt.Println("EXISTS")
	}
	return moveExists
}


func checkMoveFour(game *Game, from map[string]int, to map[string]int) bool {

	return true
}



func piecesClassical() map[int]map[int]Piece {
	return map[int]map[int]Piece{
		1: map[int]Piece{
			1: Piece{1,"Rook"},
			2: Piece{1,"Knight"},
			3: Piece{1,"Bishop"},
			4: Piece{1,"King"},
			5: Piece{1,"Queen"},
			6: Piece{1,"Bishop"},
			7: Piece{1,"Knight"},
			8: Piece{1,"Rook"},
		},
		2: map[int]Piece{
			1: Piece{1,"Pawn"},
			2: Piece{1,"Pawn"},
			3: Piece{1,"Pawn"},
			4: Piece{1,"Pawn"},
			5: Piece{1,"Pawn"},
			6: Piece{1,"Pawn"},
			7: Piece{1,"Pawn"},
			8: Piece{1,"Pawn"},
		},
		7: map[int]Piece{
			1: Piece{2,"Pawn"},
			2: Piece{2,"Pawn"},
			3: Piece{2,"Pawn"},
			4: Piece{2,"Pawn"},
			5: Piece{2,"Pawn"},
			6: Piece{2,"Pawn"},
			7: Piece{2,"Pawn"},
			8: Piece{2,"Pawn"},
		},
		8: map[int]Piece{
			1: Piece{2,"Rook"},
			2: Piece{2,"Knight"},
			3: Piece{2,"Bishop"},
			4: Piece{2,"Queen"},
			5: Piece{2,"King"},
			6: Piece{2,"Bishop"},
			7: Piece{2,"Knight"},
			8: Piece{2,"Rook"},
		},
	}
}

func piecesFour() map[int]map[int]Piece {
	return map[int]map[int]Piece{
		1: map[int]Piece{
			3: Piece{1,"Rook"},
			4: Piece{1,"Knight"},
			5: Piece{1,"Bishop"},
			6: Piece{1,"King"},
			7: Piece{1,"Queen"},
			8: Piece{1,"Bishop"},
			9: Piece{1,"Knight"},
			10: Piece{1,"Rook"},
		},
		2: map[int]Piece{
			3: Piece{1,"Pawn"},
			4: Piece{1,"Pawn"},
			5: Piece{1,"Pawn"},
			6: Piece{1,"Pawn"},
			7: Piece{1,"Pawn"},
			8: Piece{1,"Pawn"},
			9: Piece{1,"Pawn"},
			10: Piece{1,"Pawn"},
		},
		3: map[int]Piece{
			1: Piece{4,"Rook"},
			2: Piece{4,"Pawn"},
			11: Piece{2,"Rook"},
			12: Piece{2,"Pawn"},
		},
		4: map[int]Piece{
			1: Piece{4,"Knight"},
			2: Piece{4,"Pawn"},
			11: Piece{2,"Knight"},
			12: Piece{2,"Pawn"},
		},
		5: map[int]Piece{
			1: Piece{4,"Bishop"},
			2: Piece{4,"Pawn"},
			11: Piece{2,"Bishop"},
			12: Piece{2,"Pawn"},
		},
		6: map[int]Piece{
			1: Piece{4,"Queen"},
			2: Piece{4,"Pawn"},
			11: Piece{2,"Queen"},
			12: Piece{2,"Pawn"},
		},
		7: map[int]Piece{
			1: Piece{4,"King"},
			2: Piece{4,"Pawn"},
			11: Piece{2,"King"},
			12: Piece{2,"Pawn"},
		},
		8: map[int]Piece{
			1: Piece{4,"Bishop"},
			2: Piece{4,"Pawn"},
			11: Piece{2,"Bishop"},
			12: Piece{2,"Pawn"},
		},
		9: map[int]Piece{
			1: Piece{4,"Knight"},
			2: Piece{4,"Pawn"},
			11: Piece{2,"Knight"},
			12: Piece{2,"Pawn"},
		},
		10: map[int]Piece{
			1: Piece{4,"Rook"},
			2: Piece{4,"Pawn"},
			11: Piece{2,"Rook"},
			12: Piece{2,"Pawn"},
		},
		11: map[int]Piece{
			3: Piece{3,"Pawn"},
			4: Piece{3,"Pawn"},
			5: Piece{3,"Pawn"},
			6: Piece{3,"Pawn"},
			7: Piece{3,"Pawn"},
			8: Piece{3,"Pawn"},
			9: Piece{3,"Pawn"},
			10: Piece{3,"Pawn"},
		},
		12: map[int]Piece{
			3: Piece{3,"Rook"},
			4: Piece{3,"Knight"},
			5: Piece{3,"Bishop"},
			6: Piece{3,"Queen"},
			7: Piece{3,"King"},
			8: Piece{3,"Bishop"},
			9: Piece{3,"Knight"},
			10: Piece{3,"Rook"},
		},
	}}