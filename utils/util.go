package utils

import (
	"math/rand"
	"os"
	"os/exec"
	"time"
)

// ClearScreen limpa a tela do terminal
func ClearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

// HideCursor esconde o cursor no terminal
// TODO: Essa é uma função que precisa ser re-visitada para analisar melhor seu
// comportamento
func HideCursor() {
	c := exec.Command("tput civis  -- invisible")
	c.Stdout = os.Stdout
	c.Run()
}

// GenerateRandomPosition retorna uma posição aleatória
func GenerateRandomPosition(limX, limY int) (int, int) {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(limX - 1), rand.Intn(limY - 1)
}
