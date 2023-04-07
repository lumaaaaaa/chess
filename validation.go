package main

import (
	"math"
)

func validMovesAtIndices(row int, column int) []string {
	tile := board[row][column]

	if !tile.Occupied {
		return nil
	}

	var validMoves []string

	switch tile.Piece.Type {
	case Pawn:
		// starting row for white
		if !tile.Piece.Team {
			if row == 6 {
				// pawn is not blocked immediately
				if !board[row-1][column].Occupied {
					validMoves = append(validMoves, indexToColumn[column]+indexToRow[row-1])
				}

				if !board[row-2][column].Occupied {
					validMoves = append(validMoves, indexToColumn[column]+indexToRow[row-2])
				}
			} else {
				// check for normal pawn move
				if !board[row-1][column].Occupied {
					validMoves = append(validMoves, indexToColumn[column]+indexToRow[row-1])
				}
			}
		}

		// starting row for black
		if tile.Piece.Team {
			if row == 1 {
				// pawn is not blocked immediately
				if !board[row+1][column].Occupied {
					validMoves = append(validMoves, indexToColumn[column]+indexToRow[row+1])
				}

				if !board[row+2][column].Occupied {
					validMoves = append(validMoves, indexToColumn[column]+indexToRow[row+2])
				}
			} else {
				if !board[row+1][column].Occupied {
					validMoves = append(validMoves, indexToColumn[column]+indexToRow[row+1])
				}
			}
		}

		// diagonal takes
		if !tile.Piece.Team {
			if column != 0 && board[row-1][column-1].Occupied {
				validMoves = append(validMoves, indexToColumn[column-1]+indexToRow[row-1])
			}

			if column != 7 && board[row-1][column+1].Occupied {
				validMoves = append(validMoves, indexToColumn[column+1]+indexToRow[row-1])
			}
		}
		if tile.Piece.Team {
			if column != 0 && board[row+1][column-1].Occupied {
				validMoves = append(validMoves, indexToColumn[column-1]+indexToRow[row+1])
			}

			if column != 7 && board[row+1][column+1].Occupied {
				validMoves = append(validMoves, indexToColumn[column+1]+indexToRow[row+1])
			}
		}

		// handle en passant
		if tile.Piece.Team != latestMove.Piece.Team && latestMove.Piece.Type == Pawn && math.Abs(float64(latestMove.DestinationRow-latestMove.StartRow)) == 2.0 {
			if !tile.Piece.Team && row == 3 {
				if latestMove.DestinationColumn == column-1 {
					validMoves = append(validMoves, indexToColumn[column-1]+indexToRow[row-1])
				} else if latestMove.DestinationColumn == column+1 {
					validMoves = append(validMoves, indexToColumn[column+1]+indexToRow[row-1])
				}
			}
			if tile.Piece.Team && row == 4 {
				if latestMove.DestinationColumn == column-1 {
					validMoves = append(validMoves, indexToColumn[column-1]+indexToRow[row+1])
				} else if latestMove.DestinationColumn == column+1 {
					validMoves = append(validMoves, indexToColumn[column+1]+indexToRow[row+1])
				}
			}
		}
	}
	return validMoves
}
