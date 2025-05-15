package game

import (
	"maker/common"

	"github.com/hajimehoshi/ebiten/v2"
)

var savePNG common.Button = common.Button{
	Width:  150,
	Height: 50,
	X:      50,
	Y:      float32(common.ScreenHeight - 100),
	Text:   "Save PNG",
	Active: true,
}

var saveButton common.Button = common.Button{
	Width:  150,
	Height: 50,
	X:      float32(common.ScreenWidth/2 - 75),
	Y:      float32(common.ScreenHeight - 200),
	Text:   "Save",
	Active: true,
}
var cancelButton common.Button = common.Button{
	Width:  150,
	Height: 50,
	X:      float32(common.ScreenWidth/2 - 75),
	Y:      float32(common.ScreenHeight - 100),
	Text:   "Cancel",
	Active: true,
}

var saveDialog *common.TextBox = common.NewTextBox(
	int(common.ScreenWidth/2)-200, int(common.ScreenHeight/2)-100,
	200, 400,
	"Save",
)

func drawSaveMenu(screen *ebiten.Image) {
	if State == common.StateSaveDialog {
		screen.Fill(common.DisabledOverlay)
		saveDialog.Draw(screen)
		saveButton.Draw(screen)
		cancelButton.Draw(screen)
	}
}

func saveUpdate() {
	if State == common.StateSaveDialog {
		saveDialog.Update()
	}
}
