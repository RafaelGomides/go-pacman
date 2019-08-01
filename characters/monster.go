package characters

// Monster é a estrutura dos mosntros do jogo
type Monster struct {
	Move *Move
}

// NewMonster retorna um pacman novo e completamente zerado
func NewMonster() *Monster {
	mv := NewMove()
	return &Monster{
		Move: mv,
	}
}
