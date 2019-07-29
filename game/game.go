package game

import (
	char "go-pacMan/characters"
	"math/rand"
	"time"
)

// GameStatus é a estrutura do jogo em execução
type GameStatus struct {
	DotsInGrid int
	Grid       [][]string
	GridLen    struct {
		X int
		Y int
	}
	Pacman   *char.Pacman
	Monsters []*char.Monster
	Points   int
}

func (g *GameStatus) checkIfHasDot(x, y int) {
	if g.Grid[x][y] == "." {
		g.Points++
		g.DotsInGrid--
	}
}

// Up movimenta o personagem para cima
func (g *GameStatus) Up() {
	g.Grid[g.Pacman.Pos.X][g.Pacman.Pos.Y] = " "
	if g.Pacman.Pos.X-1 >= 0 {
		g.Pacman.Pos.X--
	} else {
		g.Pacman.Pos.X = g.GridLen.X - 1
	}
	g.checkIfHasDot(g.Pacman.Pos.X, g.Pacman.Pos.Y)
	g.Grid[g.Pacman.Pos.X][g.Pacman.Pos.Y] = "V"
}

// Down movimenta o personagem para baixo
func (g *GameStatus) Down() {
	g.Grid[g.Pacman.Pos.X][g.Pacman.Pos.Y] = " "
	if g.Pacman.Pos.X+1 < g.GridLen.X {
		g.Pacman.Pos.X++
	} else {
		g.Pacman.Pos.X = 0
	}
	g.checkIfHasDot(g.Pacman.Pos.X, g.Pacman.Pos.Y)
	g.Grid[g.Pacman.Pos.X][g.Pacman.Pos.Y] = "A"
}

// Left movimenta o personagem para esquerda
func (g *GameStatus) Left() {
	g.Grid[g.Pacman.Pos.X][g.Pacman.Pos.Y] = " "
	if g.Pacman.Pos.Y-1 >= 0 {
		g.Pacman.Pos.Y--
	} else {
		g.Pacman.Pos.Y = g.GridLen.Y - 1
	}
	g.checkIfHasDot(g.Pacman.Pos.X, g.Pacman.Pos.Y)
	g.Grid[g.Pacman.Pos.X][g.Pacman.Pos.Y] = ">"
}

// Right movimenta o personagem para direita
func (g *GameStatus) Right() {
	g.Grid[g.Pacman.Pos.X][g.Pacman.Pos.Y] = " "
	if g.Pacman.Pos.Y+1 < g.GridLen.Y {
		g.Pacman.Pos.Y++
	} else {
		g.Pacman.Pos.Y = 0
	}
	g.checkIfHasDot(g.Pacman.Pos.X, g.Pacman.Pos.Y)
	g.Grid[g.Pacman.Pos.X][g.Pacman.Pos.Y] = "<"
}

// GenerateManStartPosition informa a posição inicial do PacMan
func (g *GameStatus) GenerateManStartPosition() {
	rand.Seed(time.Now().UnixNano())
	xPos := rand.Intn(g.GridLen.X - 1)
	yPos := rand.Intn(g.GridLen.Y - 1)
	g.Grid[yPos][xPos] = "O"
	g.Pacman.Pos.X = xPos
	g.Pacman.Pos.Y = yPos
}
