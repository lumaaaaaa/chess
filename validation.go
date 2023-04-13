package main

import (
	"math"
)

func validMovesAtIndices(row int, column int) []Move {
	tile := board[row][column]

	if !tile.Occupied {
		return nil
	}

	var validMoves []Move

	switch tile.Piece.Type {
	case Pawn:
		// starting row for white
		if !tile.Piece.Team {
			if row == 6 {
				// pawn is not blocked immediately
				if !board[row-1][column].Occupied {
					move := Move{
						Piece:             board[row][column].Piece,
						StartRow:          row,
						StartColumn:       column,
						DestinationRow:    row - 1,
						DestinationColumn: column,
						Capture:           false,
						CapturedRow:       -1,
						CapturedColumn:    -1,
					}
					validMoves = append(validMoves, move)
				}

				if !board[row-2][column].Occupied {
					move := Move{
						Piece:             board[row][column].Piece,
						StartRow:          row,
						StartColumn:       column,
						DestinationRow:    row - 2,
						DestinationColumn: column,
						Capture:           false,
						CapturedRow:       -1,
						CapturedColumn:    -1,
					}
					validMoves = append(validMoves, move)
				}
			} else {
				// check for normal pawn move
				if !board[row-1][column].Occupied {
					move := Move{
						Piece:             board[row][column].Piece,
						StartRow:          row,
						StartColumn:       column,
						DestinationRow:    row - 1,
						DestinationColumn: column,
						Capture:           false,
						CapturedRow:       -1,
						CapturedColumn:    -1,
					}
					validMoves = append(validMoves, move)
				}
			}
		}

		// starting row for black
		if tile.Piece.Team {
			if row == 1 {
				// pawn is not blocked immediately
				if !board[row+1][column].Occupied {
					move := Move{
						Piece:             board[row][column].Piece,
						StartRow:          row,
						StartColumn:       column,
						DestinationRow:    row + 1,
						DestinationColumn: column,
						Capture:           false,
						CapturedRow:       -1,
						CapturedColumn:    -1,
					}
					validMoves = append(validMoves, move)
				}

				if !board[row+2][column].Occupied {
					move := Move{
						Piece:             board[row][column].Piece,
						StartRow:          row,
						StartColumn:       column,
						DestinationRow:    row + 2,
						DestinationColumn: column,
						Capture:           false,
						CapturedRow:       -1,
						CapturedColumn:    -1,
					}
					validMoves = append(validMoves, move)
				}
			} else {
				if !board[row+1][column].Occupied {
					move := Move{
						Piece:             board[row][column].Piece,
						StartRow:          row,
						StartColumn:       column,
						DestinationRow:    row + 1,
						DestinationColumn: column,
						Capture:           false,
						CapturedRow:       -1,
						CapturedColumn:    -1,
					}
					validMoves = append(validMoves, move)
				}
			}
		}

		// diagonal takes
		if !tile.Piece.Team {
			if column != 0 && board[row-1][column-1].Occupied {
				move := Move{
					Piece:             board[row][column].Piece,
					StartRow:          row,
					StartColumn:       column,
					DestinationRow:    row - 1,
					DestinationColumn: column - 1,
					Capture:           true,
					CapturedRow:       row - 1,
					CapturedColumn:    column - 1,
				}
				validMoves = append(validMoves, move)
			}

			if column != 7 && board[row-1][column+1].Occupied {
				move := Move{
					Piece:             board[row][column].Piece,
					StartRow:          row,
					StartColumn:       column,
					DestinationRow:    row - 1,
					DestinationColumn: column + 1,
					Capture:           true,
					CapturedRow:       row - 1,
					CapturedColumn:    column + 1,
				}
				validMoves = append(validMoves, move)
			}
		}
		if tile.Piece.Team {
			if column != 0 && board[row+1][column-1].Occupied {
				move := Move{
					Piece:             board[row][column].Piece,
					StartRow:          row,
					StartColumn:       column,
					DestinationRow:    row + 1,
					DestinationColumn: column - 1,
					Capture:           true,
					CapturedRow:       row + 1,
					CapturedColumn:    column - 1,
				}
				validMoves = append(validMoves, move)
			}

			if column != 7 && board[row+1][column+1].Occupied {
				move := Move{
					Piece:             board[row][column].Piece,
					StartRow:          row,
					StartColumn:       column,
					DestinationRow:    row + 1,
					DestinationColumn: column + 1,
					Capture:           true,
					CapturedRow:       row + 1,
					CapturedColumn:    column + 1,
				}
				validMoves = append(validMoves, move)
			}
		}

		// handle en passant
		if tile.Piece.Team != latestMove.Piece.Team && latestMove.Piece.Type == Pawn && math.Abs(float64(latestMove.DestinationRow-latestMove.StartRow)) == 2.0 {
			if !tile.Piece.Team && row == 3 {
				if latestMove.DestinationColumn == column-1 {
					move := Move{
						Piece:             board[row][column].Piece,
						StartRow:          row,
						StartColumn:       column,
						DestinationRow:    row - 1,
						DestinationColumn: column - 1,
						Capture:           true,
						CapturedRow:       latestMove.DestinationRow,
						CapturedColumn:    latestMove.DestinationColumn,
					}
					validMoves = append(validMoves, move)
				} else if latestMove.DestinationColumn == column+1 {
					move := Move{
						Piece:             board[row][column].Piece,
						StartRow:          row,
						StartColumn:       column,
						DestinationRow:    row - 1,
						DestinationColumn: column + 1,
						Capture:           true,
						CapturedRow:       latestMove.DestinationRow,
						CapturedColumn:    latestMove.DestinationColumn,
					}
					validMoves = append(validMoves, move)
				}
			}
			if tile.Piece.Team && row == 4 {
				if latestMove.DestinationColumn == column-1 {
					move := Move{
						Piece:             board[row][column].Piece,
						StartRow:          row,
						StartColumn:       column,
						DestinationRow:    row + 1,
						DestinationColumn: column - 1,
						Capture:           true,
						CapturedRow:       latestMove.DestinationRow,
						CapturedColumn:    latestMove.DestinationColumn,
					}
					validMoves = append(validMoves, move)
				} else if latestMove.DestinationColumn == column+1 {
					move := Move{
						Piece:             board[row][column].Piece,
						StartRow:          row,
						StartColumn:       column,
						DestinationRow:    row + 1,
						DestinationColumn: column + 1,
						Capture:           true,
						CapturedRow:       latestMove.DestinationRow,
						CapturedColumn:    latestMove.DestinationColumn,
					}
					validMoves = append(validMoves, move)
				}
			}
		}
	}
	return validMoves
}
