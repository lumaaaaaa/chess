package main

const (
	whiteBackground = "\033[48;2;255;243;205m"
	oliveBackground = "\033[48;2;152;118;84m"
	resetBackground = "\033[0m"
)

var (
	board     [8][8]Tile
	turn      int
	whiteMove bool
)

func main() {
	newGame()
	printBoard()
	inputMove()
}
