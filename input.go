package main

import "fmt"

// Returns the row and column index of the selected piece in the board
func inputStartPiece(team bool) (int, int) {
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
						if board[row][column].Piece.Team == team {
							return row, column
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

func inputDestinationTile(team bool) (int, int) {
	for {
		fmt.Print("destination tile > ")

		var location string
		fmt.Scanln(&location)

		if len(location) == 2 {
			if arrayContains(validColumns, location[0:1]) {
				if arrayContains(validRows, location[1:]) {
					row := rowToIndex[location[1:]]
					column := columnToIndex[location[0:1]]

					if !board[row][column].Occupied {
						if board[row][column].Piece.Team == team {
							return row, column
						} else {
							fmt.Printf("\"%s\" does not belong to the right team.\n", location)
						}
					} else {
						fmt.Printf("\"%s\" is occupied.\n", location)
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
