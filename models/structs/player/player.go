package player

import "github.com/hajimehoshi/ebiten"

type Player struct {
	Image                  *ebiten.Image
	XPos                   float64
	YPos                   float64
	Speed                  float64
	UpPressed              bool
	DownPressed            bool
	LeftPressed            bool
	RightPressed           bool
	ResidualSpeedUp        float64
	ResidualSpeedDown      float64
	ResidualSpeedLeft      float64
	ResidualSpeedRight     float64
	ResidualSpeedDecriment float64
}

func (p *Player) MovePlayer() {
	if ebiten.IsKeyPressed(ebiten.KeyUp) || ebiten.IsKeyPressed(ebiten.KeyDown) || ebiten.IsKeyPressed(ebiten.KeyLeft) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			p.YPos -= p.Speed
			p.UpPressed = true
			p.ResidualSpeedUp = p.Speed

			if p.DownPressed && !ebiten.IsKeyPressed(ebiten.KeyDown) {
				p.DownPressed = false
				p.ResidualSpeedDown = 0
			}

			if p.LeftPressed && !ebiten.IsKeyPressed(ebiten.KeyLeft) {
				p.LeftPressed = false
				p.ResidualSpeedLeft = 0
			}

			if p.RightPressed && !ebiten.IsKeyPressed(ebiten.KeyRight) {
				p.RightPressed = false
				p.ResidualSpeedRight = 0
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeyDown) {
			p.YPos += p.Speed
			p.DownPressed = true
			p.ResidualSpeedDown = p.Speed

			if p.UpPressed && !ebiten.IsKeyPressed(ebiten.KeyUp) {
				p.UpPressed = false
				p.ResidualSpeedUp = 0
			}

			if p.LeftPressed && !ebiten.IsKeyPressed(ebiten.KeyLeft) {
				p.LeftPressed = false
				p.ResidualSpeedLeft = 0
			}

			if p.RightPressed && !ebiten.IsKeyPressed(ebiten.KeyRight) {
				p.RightPressed = false
				p.ResidualSpeedRight = 0
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			p.XPos -= p.Speed
			p.LeftPressed = true
			p.ResidualSpeedLeft = p.Speed

			if p.DownPressed && !ebiten.IsKeyPressed(ebiten.KeyDown) {
				p.DownPressed = false
				p.ResidualSpeedDown = 0
			}

			if p.UpPressed && !ebiten.IsKeyPressed(ebiten.KeyUp) {
				p.UpPressed = false
				p.ResidualSpeedUp = 0
			}

			if p.RightPressed && !ebiten.IsKeyPressed(ebiten.KeyRight) {
				p.RightPressed = false
				p.ResidualSpeedRight = 0
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			p.XPos += p.Speed
			p.RightPressed = true
			p.ResidualSpeedRight = p.Speed

			if p.DownPressed && !ebiten.IsKeyPressed(ebiten.KeyDown) {
				p.DownPressed = false
				p.ResidualSpeedDown = 0
			}

			if p.LeftPressed && !ebiten.IsKeyPressed(ebiten.KeyLeft) {
				p.LeftPressed = false
				p.ResidualSpeedLeft = 0
			}

			if p.UpPressed && !ebiten.IsKeyPressed(ebiten.KeyUp) {
				p.UpPressed = false
				p.ResidualSpeedUp = 0
			}
		}
	} else {
		if p.UpPressed {
			if p.ResidualSpeedUp <= 0 {
				p.UpPressed = false
			} else {
				p.ResidualSpeedUp -= p.ResidualSpeedDecriment
				p.YPos -= p.ResidualSpeedUp
			}
		}
		if p.DownPressed {
			if p.ResidualSpeedDown <= 0 {
				p.DownPressed = false
			} else {
				p.ResidualSpeedDown -= p.ResidualSpeedDecriment
				p.YPos += p.ResidualSpeedDown
			}
		}
		if p.LeftPressed {
			if p.ResidualSpeedLeft <= 0 {
				p.LeftPressed = false
			} else {
				p.ResidualSpeedLeft -= p.ResidualSpeedDecriment
				p.XPos -= p.ResidualSpeedLeft
			}
		}
		if p.RightPressed {
			if p.ResidualSpeedRight <= 0 {
				p.RightPressed = false
			} else {
				p.ResidualSpeedRight -= p.ResidualSpeedDecriment
				p.XPos += p.ResidualSpeedRight
			}
		}
	}
}
