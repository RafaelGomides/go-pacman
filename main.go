package main

import (
	"fmt"
	"go-pacMan/game"
	u "go-pacMan/utils"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/buger/goterm"
	"github.com/mgutz/ansi"
)

var (
	normalDots string = ansi.ColorCode("white+h:black")
	superDots  string = ansi.ColorCode("white+hb:black")
	blankSpace string = ansi.ColorCode("black:black")
	reset      string = ansi.ColorCode("reset")
)

func getGridMap(x, y int) *game.Status {
	rand.Seed(time.Now().UnixNano())
	newGame := game.Status{}
	newGame.GridLen.X = x
	newGame.GridLen.Y = y
	for iy := 0; iy < y; iy++ {
		row := []string{}
		for ix := 0; ix < x; ix++ {
			if rand.Intn(99)%2 == 0 {
				newGame.DotsInGrid++
				row = append(row, fmt.Sprintf("%v.%v", normalDots, reset))
			} else if rand.Intn(99)%10 == 0 {
				newGame.DotsInGrid = newGame.DotsInGrid + 5
				row = append(row, fmt.Sprintf("%v*%v", superDots, reset))
			} else {
				row = append(row, fmt.Sprintf("%v %v", blankSpace, reset))
			}
		}
		newGame.Grid = append(newGame.Grid, row)
	}

	newGame.GenerateManStartPosition()

	return &newGame
}

func playGame(pg *game.Status) string {

	for pg.RefreshGrid() {
		u.ClearScreen()
		// ATUALIZAÇÃO DOS PONTOS
		lime := ansi.ColorCode("green+h:black")
		reset := ansi.ColorCode("reset")
		fmt.Printf("%vPontos: %v\n%v", lime, pg.Points, reset)
		// ATUALIZAÇÃO DO GRID
		for _, line := range pg.Grid {
			for _, dot := range line {
				fmt.Printf("%s", dot)
			}
			fmt.Printf("\n")
		}
	}

	return "PARABÉNS VOCÊ TERMINOU!!!"
}

func main() {
	xLen := 0
	yLen := 0

	u.HideCursor()

	if args := os.Args[1:]; len(args) == 0 {
		yLen = goterm.Height() - 2
		xLen = goterm.Width()
	} else {
		xLen, _ = strconv.Atoi(args[0])
		yLen, _ = strconv.Atoi(args[1])
	}

	fmt.Println(xLen, yLen)

	u.ClearScreen()

	fmt.Println(playGame(getGridMap(xLen, yLen)))
}
