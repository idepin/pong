package player

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Player struct {
	X     float32
	Y     float32
	Speed float32
	Size  float32
	Color color.Color
}

func (p *Player) Update() {
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.Move(1, 0)
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.Move(-1, 0)
	}

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.Move(0, -1)
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.Move(0, 1)
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	vector.FillRect(screen, p.X, p.Y, p.Size, p.Size, p.Color, true)
}

func (p *Player) Move(addX float32, addY float32) {
	p.X += addX
	p.Y += addY
}
