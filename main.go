package main

import (
	"bufio"
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
	userOption := bufio.NewReader(os.Stdin)

	for game.DotsInGrid > 1 {
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		fmt.Printf("PONTOS: %v\n", game.Points)
		for _, line := range game.Grid {
			fmt.Printf("%v\n", line)
		}
		fmt.Print("Use:\n D Para virar para Direita\n C Para virar para Cima\n B Para virar para Baixo\n E Para virar para Esquerda\n")
		opt, _ := userOption.ReadString('\n')
		fmt.Println(opt)
		switch opt {
		case "D\n":
			game.PacRight()
		case "C\n":
			game.PacUp()
		case "B\n":
			game.PacDown()
		case "E\n":
			game.PacLeft()
		default:
			fmt.Println("Informe uma opção valida")
			time.Sleep(time.Second * 1)
		}
		c.Run()
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
