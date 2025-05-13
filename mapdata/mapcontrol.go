package mapdata

import (
	"image/color"
	"maker/common"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type MapControl struct {
	originX  int
	originY  int
	length   int
	pos      int
	minValue float32
	maxValue float32
	floating bool
}

var prevMousePressed bool = false

func NewMapControl(originX int, originY int, length int, minValue, maxValue float32) *MapControl {
	return &MapControl{
		originX:  originX,
		originY:  originY,
		length:   length,
		pos:      0,
		minValue: minValue,
		maxValue: maxValue,
	}
}

func (mc *MapControl) Update() {
	// Update the position of the map control
	mouseButtonPressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	if !mouseButtonPressed {
		mc.floating = false
		prevMousePressed = false
		return
	}
	x, y := ebiten.CursorPosition()
	if common.Collide(x, y, mc) {
		mc.floating = true
	}
	if !mc.floating {
		return
	}
	if y > mc.originY+mc.length {
		mc.pos = mc.length
	} else if y < mc.originY {
		mc.pos = 0
	} else {
		mc.pos = y - mc.originY
	}
	prevMousePressed = mouseButtonPressed
}

func (mc *MapControl) GetValue() float32 {
	return float32(mc.pos)/float32(mc.length)*float32(mc.maxValue-mc.minValue) + float32(mc.minValue)
}

func (mc *MapControl) Draw(screen *ebiten.Image) {
	// Draw the map control
	// Draw the control bar
	vector.DrawFilledRect(screen, float32(mc.originX), float32(mc.originY), 10, float32(mc.length), color.RGBA{0xFF, 0x00, 0x00, 0xFF}, false)
	// Draw the control knob
	vector.DrawFilledRect(screen, float32(mc.originX), float32(mc.originY+mc.pos), 10, 10, color.RGBA{0x99, 0x99, 0x99, 0xFF}, false)
}

func (mc *MapControl) GetBounds() (float32, float32, float32, float32) {
	return float32(mc.originX), float32(mc.originY), float32(mc.originX + mc.length), float32(mc.originY + mc.length)
}
