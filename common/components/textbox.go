package components
import (
	"image/color"
	c "maker/common"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type TextBox struct {
	name   string
	x      int
	y      int
	height int
	width  int
	text   string
	active bool
}

var prevBackPressed bool = true

func NewTextBox(x, y, height, width int, name string) *TextBox {
	return &TextBox{
		name:   name,
		x:      x,
		y:      y,
		height: height,
		width:  width,
		text:   "",
		active: true,
	}
}

func (tb *TextBox) Update() {
	if !prevBackPressed && ebiten.IsKeyPressed(ebiten.KeyBackspace) && len(tb.text) > 0 {
		tb.text = tb.text[:len(tb.text)-1]
	}
	for _, r := range ebiten.AppendInputChars(nil) {
		if r > 32 && r < 127 { // all ASCII characters
			tb.text += string(r)
		} else {
			// log.Println("Invalid character:", r)
		}
	}
	prevBackPressed = ebiten.IsKeyPressed(ebiten.KeyBackspace)
}

func (tb *TextBox) Draw(screen *ebiten.Image) {
	if !tb.active {
		return
	}
	vector.DrawFilledRect(screen, float32(tb.x), float32(tb.y), float32(tb.width), float32(tb.height), c.PanelColor, false)
	vector.DrawFilledRect(screen, float32(tb.x+5), float32(tb.y+15), float32(tb.width-10), float32(tb.height-20), color.White, false)
	text.Draw(screen, tb.name, c.MenuFont, int(tb.x+7), int(tb.y+12), color.Black)
	text.Draw(screen, tb.text, c.MenuFont, int(tb.x+7), int(tb.y+30), color.Black)
}

func (tb *TextBox) GetBounds() (float32, float32, float32, float32) {
	return float32(tb.x), float32(tb.y), float32(tb.x + tb.width), float32(tb.y + tb.height)
}

func (tb *TextBox) SetActive(active bool) {
	tb.active = active
}

func (tb *TextBox) SetText(text string) {
	tb.text = text
}

func (tb *TextBox) GetText() string {
	return tb.text
}
