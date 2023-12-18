package player

import "github.com/hajimehoshi/ebiten"

type Player struct {
	Image *ebiten.Image
	XPos  float64
	YPos  float64
	Speed float64
}

func (p *Player) MovePlayer() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.YPos -= p.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.YPos += p.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.XPos -= p.Speed
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.XPos += p.Speed
	}
}
