package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"time"
)

type GameStatus struct {
	Direction  string
	DotsInGrid int
	Grid       [][]string
	GridLen    struct {
		x int
		y int
	}
	PlayerPos struct {
		x int
		y int
	}
	Points int
}

func (g *GameStatus) checkIfHasDot(x, y int) {
	if g.Grid[x][y] == "." {
		g.Points++
		g.DotsInGrid--
	}
}

func (g *GameStatus) Up() {
	g.Grid[g.PlayerPos.x][g.PlayerPos.y] = " "
	if g.PlayerPos.x-1 >= 0 {
		g.PlayerPos.x--
	} else {
		g.PlayerPos.x = g.GridLen.x - 1
	}
	g.checkIfHasDot(g.PlayerPos.x, g.PlayerPos.y)
	g.Grid[g.PlayerPos.x][g.PlayerPos.y] = "V"
}

func (g *GameStatus) Down() {
	g.Grid[g.PlayerPos.x][g.PlayerPos.y] = " "
	if g.PlayerPos.x+1 < g.GridLen.x {
		g.PlayerPos.x++
	} else {
		g.PlayerPos.x = 0
	}
	g.checkIfHasDot(g.PlayerPos.x, g.PlayerPos.y)
	g.Grid[g.PlayerPos.x][g.PlayerPos.y] = "A"
}

func (g *GameStatus) Left() {
	g.Grid[g.PlayerPos.x][g.PlayerPos.y] = " "
	if g.PlayerPos.y-1 >= 0 {
		g.PlayerPos.y--
	} else {
		g.PlayerPos.y = g.GridLen.y - 1
	}
	g.checkIfHasDot(g.PlayerPos.x, g.PlayerPos.y)
	g.Grid[g.PlayerPos.x][g.PlayerPos.y] = ">"
}

func (g *GameStatus) Right() {
	g.Grid[g.PlayerPos.x][g.PlayerPos.y] = " "
	if g.PlayerPos.y+1 < g.GridLen.y {
		g.PlayerPos.y++
	} else {
		g.PlayerPos.y = 0
	}
	g.checkIfHasDot(g.PlayerPos.x, g.PlayerPos.y)
	g.Grid[g.PlayerPos.x][g.PlayerPos.y] = "<"
}

func (g *GameStatus) generateManStartPosition() {
	rand.Seed(time.Now().UnixNano())
	xPos := rand.Intn(g.GridLen.x - 1)
	yPos := rand.Intn(g.GridLen.y - 1)
	g.Grid[xPos][yPos] = "O"
	g.PlayerPos.x = xPos
	g.PlayerPos.y = yPos
}

func getGridMap(x, y int) *GameStatus {
	rand.Seed(time.Now().UnixNano())
	newGame := GameStatus{}
	newGame.GridLen.x = x
	newGame.GridLen.y = y
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

	newGame.generateManStartPosition()

	return &newGame
}

func playGame(game *GameStatus) string {

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
			game.Direction = "D"
			game.Right()
		case "C\n":
			game.Direction = "C"
			game.Up()
		case "B\n":
			game.Direction = "B"
			game.Down()
		case "E\n":
			game.Direction = "E"
			game.Left()
		default:
			fmt.Println("Informe uma opção valida")
			time.Sleep(time.Second * 1)
		}
		c.Run()
	}

	return "PARABÉNS VOCÊ TERMINOU!!!"
}

func main() {

	if len(os.Args[1:]) == 0 {
		fmt.Println("Insira os argumentos para o funcionamento da aplicação")
		return
	}

	args := os.Args[1:]
	xLen, _ := strconv.Atoi(args[0])
	yLen, _ := strconv.Atoi(args[1])

	fmt.Println(playGame(getGridMap(xLen, yLen)))

}
