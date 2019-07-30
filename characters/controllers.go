package characters

// Move estrutura responsavel pela movimentação dos personagens
type Move struct {
	Speed int
	Pos   struct {
		X int
		Y int
	}
}

// NewMove retorna uma estrutura de movimentação padrão
func NewMove() *Move {
	return &Move{
		Speed: 1,
	}
}

func (m *Move) up() {

}

func (m *Move) down() {

}

func (m *Move) left() {

}

func (m *Move) right() {

}
