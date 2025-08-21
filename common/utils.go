package common

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Interactable interface {
	GetBounds() (float32, float32, float32, float32)
	Draw(*ebiten.Image)
}

func Collide(x, y int, button Interactable) bool {
	left, top, right, bottom := button.GetBounds()
	X, Y := float32(x), float32(y)
	return X > left && X < right && Y > top && Y < bottom
}

func RGBAEnhance(c color.Color, factor int) color.Color {
	r, g, b, a := c.RGBA()
	r8 := uint8(Clamp(int(r>>8)+factor, 0, 255))
	g8 := uint8(Clamp(int(g>>8)+factor, 0, 255))
	b8 := uint8(Clamp(int(b>>8)+factor, 0, 255))
	a8 := uint8(a >> 8)
	return &color.RGBA{r8, g8, b8, a8}
}

func Clamp(x, min, max int) int {
    if x < min { return min }
    if x > max { return max }
    return x
}
