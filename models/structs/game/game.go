package game

import (
	"log"

	"xmas-game/models/structs/player"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

var (
	err        error
	background *ebiten.Image
	spaceShip  *ebiten.Image
	playerOne  player.Player
)

type Game struct {
	ScreenWidth  int
	ScreenHeight int
}

func (g *Game) Update(screen *ebiten.Image) error {
	playerOne.MovePlayer()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(0, 0)
	screen.DrawImage(background, op)

	playerOp := &ebiten.DrawImageOptions{}
	playerOp.GeoM.Translate(playerOne.XPos, playerOne.YPos)
	screen.DrawImage(playerOne.Image, playerOp)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.ScreenWidth, g.ScreenHeight
}

func (g *Game) Initialise() {
	ebiten.SetWindowSize(g.ScreenWidth, g.ScreenHeight)
	ebiten.SetWindowTitle("Xmas game")

	background, _, err = ebitenutil.NewImageFromFile("assets/space.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	spaceShip, _, err = ebitenutil.NewImageFromFile("assets/spaceship.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatal(err)
	}

	playerOne = player.Player{
		Image:                  spaceShip,
		XPos:                   float64(g.ScreenWidth / 2.0),
		YPos:                   float64(g.ScreenHeight / 2.0),
		Speed:                  4,
		ResidualSpeedDecriment: 0.1,
	}
}
