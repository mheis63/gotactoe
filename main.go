package main

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
			//print the column numbers
			if j == 0 {
				print(string(i+1+'0') + "| ")
			}
			// print the board contents
			print(string(game.board[i][j] + '0'))
			// print the column separator
			if j < 2 {
				print(" | ")
			}

		}
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
	// Implement the logic to check if the game has ended
	// For now, let's just return false to avoid an infinite loop
	return false
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
