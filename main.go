package main

type Game struct {
	board [3][3]int
}

func main() {
	game := Game{}
	game.board = [3][3]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}}
	game.printBoard()
}

func (game *Game) printBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			print(game.board[i][j])
		}
		println()
	}
}
