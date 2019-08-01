package game

import (
	"fmt"
	char "go-pacMan/characters"
	"go-pacMan/utils"
	"math/rand"
	"time"

	"github.com/mgutz/ansi"
)

var (
	pacmancolor  string = ansi.ColorCode("yellow+hbB:black")
	monstercolor string = ansi.ColorCode("red+hb:black")
	reset        string = ansi.ColorCode("reset")
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
	xPos, yPos := utils.GenerateRandomPosition(g.GridLen.X, g.GridLen.Y)
	g.Grid[yPos][xPos] = fmt.Sprintf("%vO%v", pacmancolor, reset)
	g.Pacman = char.NewPacman()
	g.Pacman.Move.Pos.X, g.Pacman.Move.Pos.Y = xPos, yPos
}

// GenerateMonsterStartPosition cria a posição inicial dos monstros no labirinto
func (g *Status) GenerateMonsterStartPosition() {
	for mon := 0; mon < 5; mon++ {
		xPos, yPos := utils.GenerateRandomPosition(g.GridLen.X, g.GridLen.Y)
		g.Grid[yPos][xPos] = fmt.Sprintf("%vM%v", monstercolor, reset)
		m := char.NewMonster()
		m.Move.Pos.X, m.Move.Pos.Y = xPos, yPos
		g.Monsters = append(g.Monsters, m)
	}
}

// RefreshGrid atualiza o grid com as posições de todos os personagens
func (g *Status) RefreshGrid() (res bool) {
	defer time.Sleep(time.Millisecond * 500)
	return true
}
