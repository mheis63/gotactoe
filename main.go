package main

import "fmt"

type newGame struct {
	board [3][3]int
}

func (game *newGame) printBoard() {
	for i := 0; i < 3; i++ {
		// print the row numbers
		if i == 0 {
			println("  1   2  3")
			println("------------")
		}
		for j := 0; j < 3; j++ {
			fmt.Printf("%d| ", i+1)
			if j == 0 {
				print(fmt.Sprint(i+1) + "| ")
				fmt.Printf("%d", game.board[i][j])
				// print the board contents
				print(fmt.Sprint(game.board[i][j]))
				// print the column separator
			}
		}
		println("\n\n")

		if i < 2 {
			println("\n------------")
		}
	}
	println("\n\n")
}

func (game *newGame) makeMove() {
	// Implement the logic for making a move
	// For now, let's just make a dummy move
	game.board[0][0] = 1
}

func (game *newGame) checkEndOfGame() bool {
	end := false
	//check horizontal win
	for i := 0; i < 3; i++ {
		if game.board[i][0] == game.board[i][1] && game.board[i][1] == game.board[i][2] && game.board[i][0] != 0 {
			return true
		}
	}

	//check vertical win

	//check diagonal win

	return end
}

func main() {
	game := newGame{}
	game.board = [3][3]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	game.printBoard()

	// endOfGame := false
	// for !endOfGame {
	// 	game.makeMove()
	// 	game.printBoard()
	// 	endOfGame = game.checkEndOfGame()
	// }
}
