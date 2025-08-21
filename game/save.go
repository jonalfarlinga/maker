package game

import (
	c "maker/common"
	"maker/common/components"

	"github.com/hajimehoshi/ebiten/v2"
)

var savePNG components.Button = components.Button{
	Width:  3 * c.UITileUnit,
	Height: c.UITileUnit,
	X:      c.UITileUnit,
	Y:      float32(c.ScreenHeight) - 2*c.UITileUnit,
	Text:   "Save PNG",
	Active: true,
	Color:  c.ButtonColor,
}

var saveButton components.Button = components.Button{
	Width:  3 * c.UITileUnit,
	Height: c.UITileUnit,
	X:      float32(c.ScreenWidth/2) - 1.5*c.UITileUnit,
	Y:      float32(c.ScreenHeight) - 4*c.UITileUnit,
	Text:   "Save",
	Active: true,
	Color:  c.ButtonColor,
}
var cancelButton components.Button = components.Button{
	Width:  3 * c.UITileUnit,
	Height: c.UITileUnit,
	X:      float32(c.ScreenWidth/2) - 1.5*c.UITileUnit,
	Y:      float32(c.ScreenHeight - 100),
	Text:   "Cancel",
	Active: true,
	Color:  c.ButtonColor,
}

var saveDialog *components.TextBox = components.NewTextBox(
	c.ScreenWidth/2-4*int(c.UITileUnit), c.ScreenHeight/2-2*int(c.UITileUnit),
	int(4*c.UITileUnit), int(8*c.UITileUnit),
	"Save",
)

func drawSaveMenu(screen *ebiten.Image) {
	if State == c.StateSaveDialog {
		screen.Fill(c.DisabledOverlay)
		saveDialog.Draw(screen)
		saveButton.Draw(screen)
		cancelButton.Draw(screen)
	}
}

func saveUpdate() {
	if State == c.StateSaveDialog {
		saveDialog.Update()
	}
}
