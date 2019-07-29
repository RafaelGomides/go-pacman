package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/buger/goterm"
)

// GameStatus é a estrutura do jogo em execução
type GameStatus struct {
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

// Up movimenta o personagem para cima
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

// Down movimenta o personagem para baixo
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

// Left movimenta o personagem para esquerda
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

// Right movimenta o personagem para direita
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
	g.Grid[yPos][xPos] = "O"
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
			game.Right()
		case "C\n":
			game.Up()
		case "B\n":
			game.Down()
		case "E\n":
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
