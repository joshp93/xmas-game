package main

import (
	"log"
	. "xmas-game/models/structs/game"

	"github.com/hajimehoshi/ebiten"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Xmas game")
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
