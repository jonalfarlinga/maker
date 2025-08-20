package game

import (
	"maker/common"
	"maker/mapdata"
	"math/rand"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

var prevMousePressed bool = false

func mouseUpdate() {
	mouseButtonPressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	if !mouseButtonPressed && prevMousePressed {
		x, y := ebiten.CursorPosition()
		switch {
		case common.Collide(x, y, &smoothLandformsButton):
			mapData.SmoothLandforms()
			common.DebugPrintln("mouse", "SmoothLandformsButton clicked")

		case common.Collide(x, y, &generateButton):
			mapData = mapdata.NewMapArray(
				int(resolutionBarX.GetValue()), int(resolutionBarY.GetValue()))
			falloffProb := (falloffProbBar.GetValue() *
				float32(
					max(mapData.Width, mapData.Height)/
						common.WorldFillReductionFactor))
			border := min(mapData.Height, mapData.Width) / common.WorldEdgeFraction
			n := int(numberOfIslandsBar.GetValue())
			mapData.GenerateIsland(mapData.Width/2, mapData.Height/2, int(falloffProb))
			for i := 1; i < n; i++ {
				x := border + rand.Intn(border*(common.WorldEdgeFraction-2))
				y := border + rand.Intn(border*(common.WorldEdgeFraction-2))
				mapData.GenerateIsland(x, y, int(falloffProb))
			}
			common.DebugPrintln("mouse", "GenerateButton clicked")

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
		}
	}
	prevMousePressed = mouseButtonPressed
}
