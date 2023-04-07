package main

const (
	whiteBackground = "\033[48;2;255;243;205m"
	oliveBackground = "\033[48;2;152;118;84m"
	resetBackground = "\033[0m"

	Size = 8
)

var (
	board       [Size][Size]Tile
	turn        int
	currentTeam bool // false for white, true for black
)

func main() {
	newGame()

	/*board[5][5].Piece = Piece{
		Team: Black,
		Type: Queen,
	}

	board[5][5].Occupied = true*/

	printBoard()
	handleMove()
}
