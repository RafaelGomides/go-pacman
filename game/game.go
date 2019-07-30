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

// GenerateManStartPosition informa a posição inicial do PacMan
func (g *GameStatus) GenerateManStartPosition() {
	rand.Seed(time.Now().UnixNano())
	xPos := rand.Intn(g.GridLen.X - 1)
	yPos := rand.Intn(g.GridLen.Y - 1)
	g.Grid[yPos][xPos] = "O"
	g.Pacman = char.NewPacman()
	g.Pacman.Move.Pos.X = xPos
	g.Pacman.Move.Pos.Y = yPos
}

// GenerateMonsterStartPosition cria a posição inicial dos monstros no labirinto
func (g *GameStatus) GenerateMonsterStartPosition() {

}

// RefreshGrid atualiza o grid com as posições de todos os personagens
func (g *GameStatus) RefreshGrid() (res bool) {

	return
}
