package main

import "fmt"

// Returns the row and column index of the selected piece in the board
func inputStartPiece() (int, int, []Move) {
	for {
		fmt.Print("start piece > ")

		var location string
		fmt.Scanln(&location)

		if len(location) == 2 {
			if arrayContains(validColumns, location[0:1]) {
				if arrayContains(validRows, location[1:]) {
					row := rowToIndex[location[1:]]
					column := columnToIndex[location[0:1]]

					if board[row][column].Occupied {
						if board[row][column].Piece.Team == currentTeam {
							validMoves := validMovesAtIndices(row, column)

							if len(validMoves) != 0 {
								return row, column, validMoves
							} else {
								fmt.Printf("\"%s\" does not have any available moves.\n", location)
							}
						} else {
							fmt.Printf("\"%s\" does not belong to the right team.\n", location)
						}
					} else {
						fmt.Printf("\"%s\" is unoccupied.\n", location)
					}
				} else {
					fmt.Printf("'%s' is not a valid row.\n", location[1:])
				}
			} else {
				fmt.Printf("'%s' is not a valid column.\n", location[0:1])
			}
		} else {
			fmt.Println("input should be in proper format (ex. \"e2\")")
		}
	}
}

func inputDestinationTile(validMoves []Move) Move {
	for {
		fmt.Print("destination tile > ")

		var location string
		fmt.Scanln(&location)

		if len(location) == 2 {
			if arrayContains(validColumns, location[0:1]) {
				if arrayContains(validRows, location[1:]) {
					move, valid := findMoveGivenDestination(validMoves, location)

					if valid {
						row := rowToIndex[location[1:]]
						column := columnToIndex[location[0:1]]
						if board[row][column].Occupied {
							if board[row][column].Piece.Team != currentTeam {
								return move
							} else {
								fmt.Printf("\"%s\" does not belong to the right team.\n", location)
							}
						} else {
							return move
						}
					} else {
						fmt.Printf("\"%s\" is not a valid move.\n", location)
					}

				} else {
					fmt.Printf("'%s' is not a valid row.\n", location[1:])
				}
			} else {
				fmt.Printf("'%s' is not a valid column.\n", location[0:1])
			}
		} else {
			fmt.Println("input should be in proper format (ex. \"e2\")")
		}
	}
}

func findMoveGivenDestination(moves []Move, location string) (Move, bool) {
	row := rowToIndex[location[1:]]
	column := columnToIndex[location[0:1]]

	for _, move := range moves {
		if move.DestinationRow == row && move.DestinationColumn == column {
			return move, true
		}
	}

	return Move{}, false
}
