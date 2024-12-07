package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type newGame struct {
	board [3][3]string
}

func (game *newGame) printBoard() {
	for i := 0; i < 3; i++ {
		// print the row numbers
		if i == 0 {
			println("  1   2  3")
			println("------------")
		}
		for j := 0; j < 3; j++ {
			//fmt.Printf("%d| ", i+1)
			if j == 0 {
				print(fmt.Sprint(i+1) + "| ")
			}
			print(fmt.Sprint(game.board[i][j]) + "  ")
		}

		if i < 2 {
			println("\n------------")
		}
	}
	println("\n\n")
}

func (game *newGame) makeMove() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your move (X or O): ")
	move, _ := reader.ReadString('\n')
	move = strings.TrimSpace(move)

	var row, col int
	fmt.Print("Enter the row (1-3): ")
	fmt.Scanf("%d", &row)
	fmt.Print("Enter the column (1-3): ")
	fmt.Scanf("%d", &col)

	if row >= 1 && row <= 3 && col >= 1 && col <= 3 && game.board[row-1][col-1] == "_" {
		game.board[row-1][col-1] = move
	} else {
		fmt.Println("Invalid move, try again.")
		game.makeMove()
	}
}

func (game *newGame) checkEndOfGame() bool {
	end := false
	//check horizontal win
	for i := 0; i < 3; i++ {
		if game.board[i][0] == game.board[i][1] && game.board[i][1] == game.board[i][2] && game.board[i][0] != "_" {
			return true
		}
	}

	//check vertical win
	for i := 0; i < 3; i++ {
		if game.board[0][i] == game.board[1][i] && game.board[1][i] == game.board[2][i] && game.board[0][i] != "_" {
			return true
		}
	}

	//check diagonal win
	if game.board[0][0] == game.board[1][1] && game.board[1][1] == game.board[2][2] && game.board[0][0] != "_" {
		return true
	}

	return end
}

func main() {
	game := newGame{}
	game.board = [3][3]string{{"_", "_", "_"}, {"_", "_", "_"}, {"_", "_", "_"}}
	game.printBoard()

	endOfGame := false
	for !endOfGame {
		game.makeMove()
		game.printBoard()
		endOfGame = game.checkEndOfGame()
	}
}
