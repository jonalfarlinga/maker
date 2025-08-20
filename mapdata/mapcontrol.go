package mapdata

import (
	"image/color"
	"maker/common"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type MapControl struct {
	name     string
	originX  int
	originY  int
	length   int
	pos      int
	minValue float32
	maxValue float32
	floating bool
}

func NewMapControl(originX, originY, length, minValue, maxValue float32, name string) *MapControl {
	return &MapControl{
		name:     name,
		originX:  int(originX),
		originY:  int(originY),
		length:   int(length),
		pos:      0,
		minValue: minValue,
		maxValue: maxValue,
	}
}

func (mc *MapControl) Update() {
	// Update the position of the map control
	mouseButtonPressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	if !mouseButtonPressed {
		if mc.floating {
			mc.floating = false
			common.DebugPrintln("mapcontrol", "MapControl", mc.name, "released at position:", mc.pos)
		}
		return
	}
	x, y := ebiten.CursorPosition()
	if common.Collide(x, y, mc) {
		mc.floating = true
	}
	if !mc.floating {
		return
	}
	if x > mc.originX+mc.length {
		mc.pos = mc.length
	} else if x < mc.originX {
		mc.pos = 0
	} else {
		mc.pos = x - mc.originX
	}
}

func (mc *MapControl) BoundUpdate(twin *MapControl) {
	if mc.floating {
		twin.pos = mc.pos
	}
	mc.Update()
}

func (mc *MapControl) GetValue() float32 {
	return float32(mc.pos)/float32(mc.length)*float32(mc.maxValue-mc.minValue) + float32(mc.minValue)
}

func (mc *MapControl) Draw(screen *ebiten.Image) {
	// Draw the control bar
	vector.DrawFilledRect(screen, float32(mc.originX), float32(mc.originY), float32(mc.length+15), 15, color.RGBA{0xFF, 0x00, 0x00, 0xFF}, false)
	// Draw the control knob
	vector.DrawFilledRect(screen, float32(mc.originX+mc.pos)+1, float32(mc.originY)+1, 13, 13, color.RGBA{0x99, 0x99, 0x99, 0xFF}, false)
	// Draw the Name
	text.Draw(screen, mc.name, common.MenuFont, mc.originX+5, mc.originY-3, color.White)
	// Draw the value
	value := strconv.Itoa(int(mc.GetValue()))
	text.Draw(screen, value, common.MenuFont, mc.originX+mc.length/2, mc.originY+12, color.White)
}

func (mc *MapControl) GetBounds() (float32, float32, float32, float32) {
	return float32(mc.originX), float32(mc.originY), float32(mc.originX + mc.length), float32(mc.originY + 15)
}
