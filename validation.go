package main

func validMovesAtIndices(row int, column int) []string {
	tile := board[row][column]

	if !tile.Occupied {
		return nil
	}

	var validMoves []string

	switch tile.Piece.Type {
	case Pawn:
		// starting row for white
		if !tile.Piece.Team && row == 6 {
			// pawn is not blocked immediately
			if !board[row-1][column].Occupied {
				validMoves = append(validMoves, indexToColumn[column]+indexToRow[row-1])
			}

			if !board[row-2][column].Occupied {
				validMoves = append(validMoves, indexToColumn[column]+indexToRow[row-2])
			}
		}

		// starting row for black
		if tile.Piece.Team && row == 1 {
			// pawn is not blocked immediately
			if !board[row+1][column].Occupied {
				validMoves = append(validMoves, indexToColumn[column]+indexToRow[row+1])
			}

			if !board[row+2][column].Occupied {
				validMoves = append(validMoves, indexToColumn[column]+indexToRow[row+2])
			}
		}

		// diagonal takes
		if !tile.Piece.Team {
			if column != 0 && board[row-1][column-1].Occupied {
				validMoves = append(validMoves, indexToColumn[column-1]+indexToRow[row-1])
			}

			if column != 7 && board[row-1][column+1].Occupied {
				validMoves = append(validMoves, indexToColumn[column-1]+indexToRow[row-1])
			}
		}

		// TODO: handle en passant

	}
	return validMoves
}
