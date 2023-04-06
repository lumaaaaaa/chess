package main

import (
	"fmt"
)

const (
	Pawn   = 0
	Bishop = 1
	Knight = 2
	Rook   = 3
	Queen  = 4
	King   = 5

	White = 0
	Black = 1
)

type Tile struct {
	Occupied bool
	Piece    Piece
}

type Piece struct {
	Team int // 0 == white, 1 == black
	Type int
}

func newGame() {
	turn = 1
	whiteMove = true

	// black row 1
	board[0] = [8]Tile{
		{
			Occupied: true,
			Piece: Piece{
				Team: Black,
				Type: Rook,
			},
		},
		{
			Occupied: true,
			Piece: Piece{
				Team: Black,
				Type: Knight,
			},
		},
		{
			Occupied: true,
			Piece: Piece{
				Team: Black,
				Type: Bishop,
			},
		},
		{
			Occupied: true,
			Piece: Piece{
				Team: Black,
				Type: Queen,
			},
		},
		{
			Occupied: true,
			Piece: Piece{
				Team: Black,
				Type: King,
			},
		},
		{
			Occupied: true,
			Piece: Piece{
				Team: Black,
				Type: Bishop,
			},
		},
		{
			Occupied: true,
			Piece: Piece{
				Team: Black,
				Type: Knight,
			},
		},
		{
			Occupied: true,
			Piece: Piece{
				Team: Black,
				Type: Rook,
			},
		},
	}

	// black pawns
	for i := 0; i < len(board[1]); i++ {
		board[1][i] = Tile{
			Occupied: true,
			Piece:    Piece{Team: Black, Type: Pawn},
		}
	}

	// white pawns
	for i := 0; i < len(board[6]); i++ {
		board[6][i] = Tile{
			Occupied: true,
			Piece:    Piece{Team: White, Type: Pawn},
		}
	}

	// white row 1
	board[7] = [8]Tile{
		{
			Occupied: true,
			Piece: Piece{
				Team: White,
				Type: Rook,
			},
		},
		{
			Occupied: true,
			Piece: Piece{
				Team: White,
				Type: Knight,
			},
		},
		{
			Occupied: true,
			Piece: Piece{
				Team: White,
				Type: Bishop,
			},
		},
		{
			Occupied: true,
			Piece: Piece{
				Team: White,
				Type: Queen,
			},
		},
		{
			Occupied: true,
			Piece: Piece{
				Team: White,
				Type: King,
			},
		},
		{
			Occupied: true,
			Piece: Piece{
				Team: White,
				Type: Bishop,
			},
		},
		{
			Occupied: true,
			Piece: Piece{
				Team: White,
				Type: Knight,
			},
		},
		{
			Occupied: true,
			Piece: Piece{
				Team: White,
				Type: Rook,
			},
		},
	}
}

func printBoard() {
	color := whiteBackground
	for i, row := range board {
		fmt.Printf("%d ", 8-i)

		for _, tile := range row {
			if tile.Occupied {
				printColor := "\033[30m"

				if tile.Piece.Team == White {
					printColor = "\033[37m"
				}

				switch tile.Piece.Type {
				case Pawn:
					fmt.Printf("%s%s♟ %s", printColor, color, resetBackground)
				case Bishop:
					fmt.Printf("%s%s♝ %s", printColor, color, resetBackground)
				case Knight:
					fmt.Printf("%s%s♞ %s", printColor, color, resetBackground)
				case Rook:
					fmt.Printf("%s%s♜ %s", printColor, color, resetBackground)
				case Queen:
					fmt.Printf("%s%s♛ %s", printColor, color, resetBackground)
				case King:
					fmt.Printf("%s%s♚ %s", printColor, color, resetBackground)
				}
			} else {
				fmt.Printf("%s  %s", color, resetBackground)
			}

			if color == whiteBackground {
				color = oliveBackground
			} else {
				color = whiteBackground
			}
		}

		if color == whiteBackground {
			color = oliveBackground
		} else {
			color = whiteBackground
		}

		fmt.Println()
	}

	fmt.Println("  a b c d e f g h")
}

func inputMove() {
	if whiteMove {
		fmt.Printf("Turn %d: White's move.\n", turn)
	} else {
		fmt.Printf("Turn %d: Black's move.\n", turn)
	}

	fmt.Print(" > ")

}
