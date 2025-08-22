package game

import (
	"maker/common"
	"maker/settlements"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

var prevMousePressed bool = false

func mouseUpdate() {
	mouseButtonPressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	if !mouseButtonPressed && prevMousePressed {
		x, y := ebiten.CursorPosition()
		if settlements.PlacingSettlements() && common.Collide(x, y, mapData) {
			placeNewSettlement(x, y)
		} else {
			switch {
			case common.Collide(x, y, &smoothLandformsButton):
				mapData.SmoothLandforms()
				common.DebugPrintln("mouse", "SmoothLandformsButton clicked")

			case common.Collide(x, y, &generateButton):
				handleGenerateIsland()

			case common.Collide(x, y, &terraformLakesButton):
				mapData.TerraformLakes(fillinBar.GetValue())
				common.DebugPrintln("mouse", "TerraformLakesButton clicked")

			case common.Collide(x, y, &savePNG):
				State = common.StateSaveDialog
				saveDialog.SetActive(true)
				saveDialog.SetText("map")
				common.DebugPrintln("mouse", "SavePNG clicked, opening save dialog")

			case common.Collide(x, y, &saveButton) && State == common.StateSaveDialog:
				mapData.OutputPNG(saveDialog.GetText() + ".png")
				State = common.StateMain
				common.DebugPrintln("mouse", "SaveButton clicked, saving map as:", saveDialog.GetText()+".png")

			case common.Collide(x, y, &cancelButton) && State == common.StateSaveDialog:
				State = common.StateMain
				saveDialog.SetActive(false)
				common.DebugPrintln("mouse", "CancelButton clicked, closing save dialog")

			case common.Collide(x, y, &exitButton):
				common.DebugPrintln("mouse", "ExitButton clicked, exiting application")
				os.Exit(0)

			case common.Collide(x, y, &settlementsButton):
				common.DebugPrintln("mouse", "SettlementsButton clicked, toggling settlement placement")
				on := settlements.TogglePlacing()
				if on {
					settlementsButton.Color = common.ButtonGlowColor
				} else {
					settlementsButton.Color = common.ButtonColor
				}
			}
		}
	}
	prevMousePressed = mouseButtonPressed
}
