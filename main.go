package main

import (
	"fmt"
	"go-pacMan/game"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/buger/goterm"
)

func getGridMap(x, y int) *game.GameStatus {
	rand.Seed(time.Now().UnixNano())
	newGame := game.GameStatus{}
	newGame.GridLen.X = x
	newGame.GridLen.Y = y
	for iy := 0; iy < y; iy++ {
		row := []string{}
		for ix := 0; ix < x; ix++ {
			if rand.Intn(99)%2 == 0 {
				newGame.DotsInGrid++
				row = append(row, ".")
			} else {
				row = append(row, " ")

			}
		}
		newGame.Grid = append(newGame.Grid, row)
	}

	newGame.GenerateManStartPosition()

	return &newGame
}

func playGame(game *game.GameStatus) string {

	for game.RefreshGrid() {
		// TODO: Sistema de refresh do jogo
	}

	return "PARABÉNS VOCÊ TERMINOU!!!"
}

func main() {
	xLen := 0
	yLen := 0

	if args := os.Args[1:]; len(args) == 0 {
		yLen = goterm.Height() - 7
		xLen = goterm.Width() / 2
	} else {
		xLen, _ = strconv.Atoi(args[0])
		yLen, _ = strconv.Atoi(args[1])
	}

	fmt.Println(xLen, yLen)

	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()

	fmt.Println(playGame(getGridMap(xLen, yLen)))
}
