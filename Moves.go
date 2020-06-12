package main

func isEmptyCellClassical(game *Game, row int, col int) bool {
	if _, ok := game.Pieces[row]; ok {
		if _, ok2 := game.Pieces[row][col]; ok2 {
			return false
		}
	}
	return true
}

func canMoveToClassical(game *Game, piece Piece, row int, col int) bool {
	if !(row >= 1 && row <= 8) {
		return false
	}
	if isEmptyCellClassical(game, row, col) {
		return true
	}
	targetPiece := game.Pieces[row][col]
	return targetPiece.Color != piece.Color
}

func tryToAddMoveClassical(moves *[]map[string]int, game *Game, piece Piece, row int, col int) {
	if canMoveToClassical(game, piece, row, col) {
		*moves = append(*moves, map[string]int{"row": row, "col": col})
	}
}

func getClassicalPawnMoves(game *Game, row int, col int) []map[string]int {
	piece := game.Pieces[row][col]
	var moves []map[string]int
	if piece.Color == 2 {
		if canMoveToClassical(game, piece, row-1, col) {
			tryToAddMoveClassical(&moves, game, piece, row-1, col)
			if row == 7 {
				tryToAddMoveClassical(&moves, game, piece, row-2, col)
			}
		}
	} else {
		if canMoveToClassical(game, piece, row+1, col) {
			tryToAddMoveClassical(&moves, game, piece, row+1, col)
			if row == 2 {
				tryToAddMoveClassical(&moves, game, piece, row+2, col)
			}
		}
	}
	return moves
}


func getClassicalKingMoves(game *Game, row int, col int) []map[string]int {
	piece := game.Pieces[row][col]
	var moves []map[string]int
	tryToAddMoveClassical(&moves, game, piece, row-1, col)
	tryToAddMoveClassical(&moves, game, piece, row+1, col)
	tryToAddMoveClassical(&moves, game, piece, row, col-1)
	tryToAddMoveClassical(&moves, game, piece, row, col+1)
	tryToAddMoveClassical(&moves, game, piece, row-1, col+1)
	tryToAddMoveClassical(&moves, game, piece, row+1, col-1)
	tryToAddMoveClassical(&moves, game, piece, row-1, col-1)
	tryToAddMoveClassical(&moves, game, piece, row+1, col+1)
	return moves
}

func getClassicalRookMoves(game *Game, row int, col int) []map[string]int {
	piece := game.Pieces[row][col]
	var moves []map[string]int
	for i, j := row-1, col; i > 0 && canMoveToClassical(game, piece, i, j);i-- {
		tryToAddMoveClassical(&moves, game, piece, i, j)
	}
	for i, j := row+1, col; i <= 8 && canMoveToClassical(game, piece, i, j);i++ {
		tryToAddMoveClassical(&moves, game, piece, i, j)
	}
	for i, j := row, col-1; j > 0 && canMoveToClassical(game, piece, i, j);j-- {
		tryToAddMoveClassical(&moves, game, piece, i, j)
	}
	for i, j := row, col+1; j <= 8 && canMoveToClassical(game, piece, i, j);j++ {
		tryToAddMoveClassical(&moves, game, piece, i, j)
	}

	return moves
}


func getClassicalBishopMoves(game *Game, row int, col int) []map[string]int {
	piece := game.Pieces[row][col]
	var moves []map[string]int
	for i, j := row-1, col-1; i > 0 && j > 0 && canMoveToClassical(game, piece, i, j);i, j = i-1, j-1 {
		tryToAddMoveClassical(&moves, game, piece, i, j)
	}
	for i, j := row+1, col+1; i <= 8 && j <= 8 && canMoveToClassical(game, piece, i, j);i, j = i+1, j+1 {
		tryToAddMoveClassical(&moves, game, piece, i, j)
	}
	for i, j := row+1, col-1; i <= 8 && j > 0 && canMoveToClassical(game, piece, i, j);i, j = i+1, j-1 {
		tryToAddMoveClassical(&moves, game, piece, i, j)
	}
	for i, j := row-1, col+1; i > 0 && j <= 8 && canMoveToClassical(game, piece, i, j);i, j = i-1, j+1 {
		tryToAddMoveClassical(&moves, game, piece, i, j)
	}

	return moves
}

func getClassicalQueenMoves(game *Game, row int, col int) []map[string]int {
	return append(getClassicalBishopMoves(game, row, col), getClassicalRookMoves(game, row, col)...)
}


func getClassicalKnightMoves(game *Game, row int, col int) []map[string]int {
	piece := game.Pieces[row][col]
	var moves []map[string]int
	if row + 1 <= 8 && col + 2 <= 8 && canMoveToClassical(game, piece, row + 1, col + 2) {
		tryToAddMoveClassical(&moves, game, piece, row + 1, col + 2)
	}
	if row + 1 <= 8 && col - 2 > 0 && canMoveToClassical(game, piece, row + 1, col - 2) {
		tryToAddMoveClassical(&moves, game, piece, row + 1, col - 2)
	}
	if row - 1 > 0 && col + 2 <= 8 && canMoveToClassical(game, piece, row - 1, col + 2) {
		tryToAddMoveClassical(&moves, game, piece, row - 1, col + 2)
	}
	if row - 1 > 0 && col - 2 > 0 && canMoveToClassical(game, piece, row - 1, col - 2) {
		tryToAddMoveClassical(&moves, game, piece, row - 1, col - 2)
	}

	if row + 2 <= 8 && col + 1 <= 8 && canMoveToClassical(game, piece, row + 2, col + 1) {
		tryToAddMoveClassical(&moves, game, piece, row + 2, col + 1)
	}
	if row + 2 <= 8 && col - 1 > 0 && canMoveToClassical(game, piece, row + 2, col - 1) {
		tryToAddMoveClassical(&moves, game, piece, row + 2, col - 1)
	}
	if row - 2 > 0 && col + 1 <= 8 && canMoveToClassical(game, piece, row - 2, col + 1) {
		tryToAddMoveClassical(&moves, game, piece, row - 2, col + 1)
	}
	if row - 2 > 0 && col - 1 > 0 && canMoveToClassical(game, piece, row - 2, col - 1) {
		tryToAddMoveClassical(&moves, game, piece, row - 2, col - 1)
	}

	return moves
}
