package main

import (
	"log"
	. "xmas-game/models/structs/game"

	"github.com/hajimehoshi/ebiten"
)

var game *Game

func initGame() *Game {
	return &Game{
		ScreenWidth:  680,
		ScreenHeight: 480,
	}
}

func init() {
	game := initGame()
	game.Initialise()
}

func main() {
	game := initGame()
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
