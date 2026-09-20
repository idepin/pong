package object

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func DrawCube(screen *ebiten.Image) {
	vector.FillRect(screen, 10, 10, 100, 10, color.White, true)
}
