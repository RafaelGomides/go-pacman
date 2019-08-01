package game

import (
	"fmt"
	char "go-pacMan/characters"
	"math/rand"
	"time"

	"github.com/mgutz/ansi"
)

var (
	pacmancolor string = ansi.ColorCode("yellow+hbB:black")
	reset       string = ansi.ColorCode("reset")
)

// Status é a estrutura do jogo em execução
type Status struct {
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

// GenerateManStartPosition informa a posição inicial do PacMan
func (g *Status) GenerateManStartPosition() {
	rand.Seed(time.Now().UnixNano())
	xPos := rand.Intn(g.GridLen.X - 1)
	yPos := rand.Intn(g.GridLen.Y - 1)
	g.Grid[yPos][xPos] = fmt.Sprintf("%vO%v", pacmancolor, reset)
	g.Pacman = char.NewPacman()
	g.Pacman.Move.Pos.X = xPos
	g.Pacman.Move.Pos.Y = yPos
}

// GenerateMonsterStartPosition cria a posição inicial dos monstros no labirinto
func (g *Status) GenerateMonsterStartPosition() {

}

// RefreshGrid atualiza o grid com as posições de todos os personagens
func (g *Status) RefreshGrid() (res bool) {
	defer time.Sleep(time.Millisecond * 500)
	return true
}
