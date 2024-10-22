```
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	boardSize = 3
	maxTurns  = 9
)

type Board [boardSize][boardSize]byte

func drawBoard(board Board) {
	for i := 0; i < boardSize; i++ {
		fmt.Println("-------------")
		for j := 0; j < boardSize; j++ {
			fmt.Printf("| %c ", board[i][j])
		}
		fmt.Println("\n-------------")
	}
}

func checkWin(board Board, player byte) bool {
	for i := 0; i < boardSize; i++ {
		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
			return true
		}
		if board[0][i] == player && board[1][i] == player && board[2][i] == player {
			return true
		}
	}
	if board[0][0] == player && board[1][1] == player && board[2][2] == player {
		return true
	}
	if board[0][2] == player && board[1][1] == player && board[2][0] == player {
		return true
	}
	return false
}

func main() {
	var board Board
	player := 'X'
	turn := 0

	fmt.Println("Welcome to Tic-Tac-Toe!")

	scanner := bufio.NewScanner(os.Stdin)
	for turn < maxTurns {
		drawBoard(board)

		fmt.Printf("Player %c, enter row (0-2) and column (0-2): ", player)
		scanner.Scan()
		input := scanner.Text()
		inputs := strings.Split(input, " ")

		row, _ := strconv.Atoi(inputs[0])
		col, _ := strconv.Atoi(inputs[1])

		if board[row][col] != ' ' || row < 0 || row > 2 || col < 0 || col > 2 {
			fmt.Println("Invalid move. Try again.")
			continue
		}

		board[row][col] = player

		if checkWin(board, player) {
			drawBoard(board)
			fmt.Printf("Player %c wins!\n", player)
			break
		}

		player = 'O'
		if player == 'O' {
			player = 'X'
		}

		turn++
	}

	drawBoard(board)
	if turn == maxTurns && !checkWin(board, 'X') && !checkWin(board, 'O') {
		fmt.Println("It's a draw!")
	}
}
```
