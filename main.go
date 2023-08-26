package main

import (
	"math/rand"
	"strconv"
	"time"
)

const (
	boardSize              = 10
	aliveMarker            = "X"
	deadMarker             = " "
	initialAlivePercentage = 0.3
)

var board [boardSize][boardSize]string

func createDefaultBoard() [boardSize][boardSize]string {
	var defaultBoard [boardSize][boardSize]string
	for i := range defaultBoard {
		for j := range defaultBoard[i] {
			defaultBoard[i][j] = deadMarker
		}
	}
	return defaultBoard
}

func copyBoard(src [boardSize][boardSize]string) [boardSize][boardSize]string {
	var dest [boardSize][boardSize]string

	// option 1
	// for i := range src {
	// 	for j := range src[i] {
	// 		dest[i][j] = src[i][j]
	// 		copy(dest[i][j], src[i][j])
	// 	}
	// }

	// option 2
	// copy(dest[:], src[:])

	// option 3
	for i := range src {
		copy(dest[i][:], src[i][:])
	}

	return dest
}

func drawBoard() {
	for _, row := range board {
		for _, cell := range row {
			print(cell)
		}
		println()
	}
}

func checkNeighbours(x int, y int, b [boardSize][boardSize]string) string {
	aliveNeighbours := 0
	max := len(board) - 1

	for nx := x - 1; nx <= x+1; nx++ {
		for ny := y - 1; ny <= y+1; ny++ {
			if nx >= 0 && nx <= max && ny >= 0 && ny <= max {
				if b[nx][ny] == aliveMarker && !(nx == x && ny == y) {
					aliveNeighbours++
				}
			}
		}
	}

	if b[x][y] == aliveMarker {
		if aliveNeighbours < 2 || aliveNeighbours > 3 {
			return deadMarker
		} else {
			return aliveMarker
		}
	}

	if aliveNeighbours == 3 {
		return aliveMarker
	}

	return deadMarker
}

func lifeLoop() {
	newBoard := copyBoard(board)

	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			newBoard[i][j] = checkNeighbours(i, j, board)
		}
	}

	board = newBoard
}

func randomizeStartValueOnBoard() {
	for i := range board {
		for j := range board[i] {
			if rand.Float64() < initialAlivePercentage {
				board[i][j] = aliveMarker
			}
		}
	}
}

func main() {
	board = createDefaultBoard()
	timesToLoop := 100
	randomizeStartValueOnBoard()
	println("--- Frame 0 --------------------")
	drawBoard()

	for i := 0; i < timesToLoop; i++ {
		lifeLoop()
		frameNrAsString := strconv.Itoa(i + 1)
		println("--- Frame " + frameNrAsString + " --------------------")
		drawBoard()
		time.Sleep(1 * time.Second)
	}
}
