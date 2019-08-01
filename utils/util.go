package utils

import (
	"os"
	"os/exec"
)

// ClearScreen limpa a tela do terminal
func ClearScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

// HideCursor esconde o cursor no terminal
func HideCursor() {
	c := exec.Command("tput civis  -- invisible")
	c.Stdout = os.Stdout
	c.Run()
}
