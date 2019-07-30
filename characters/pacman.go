package characters

type Pacman struct {
	Move *Move
}

// NewPacman retorna um pacman novo e completamente zerado
func NewPacman() *Pacman {
	mv := NewMove()
	return &Pacman{
		Move: mv,
	}
}

// Up movimenta o personagem para cima
func (g *Pacman) Up() {
	g.Move.up()
}

// Down movimenta o personagem para baixo
func (g *Pacman) Down() {
	g.Move.down()
}

// Left movimenta o personagem para esquerda
func (g *Pacman) Left() {
	g.Move.left()
}

// Right movimenta o personagem para direita
func (g *Pacman) Right() {
	g.Move.right()
}
