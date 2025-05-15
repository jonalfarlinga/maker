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
		case common.Collide(x, y, &generateButton):
			falloffProb := falloffProbBar.GetValue() * float32(max(mapData.Width, mapData.Height)) / 2
			mapData = mapdata.NewMapArray(int(resolutionBarX.GetValue()), int(resolutionBarY.GetValue()))
			border := min(mapData.Height, mapData.Width) / 8
			n := int(numberOfIslandsBar.GetValue())
			mapData.GenerateIsland(mapData.Width/2, mapData.Height/2, int(falloffProb))
			for i := 1; i < n; i++ {
				x := border + rand.Intn(border*6)
				y := border + rand.Intn(border*6)
				mapData.GenerateIsland(x, y, int(falloffProb))
			}
		case common.Collide(x, y, &terraformLakesButton):
			mapData.TerraformLakes(int(fillinBar.GetValue()))
		case common.Collide(x, y, &savePNG):
			State = common.StateSaveDialog
			saveDialog.SetActive(true)
			saveDialog.SetText("map")
		case common.Collide(x, y, &saveButton) && State == common.StateSaveDialog:
			mapData.OutputPNG(saveDialog.GetText() + ".png")
			State = common.StateMain
		case common.Collide(x, y, &cancelButton) && State == common.StateSaveDialog:
			State = common.StateMain
			saveDialog.SetActive(false)
		case common.Collide(x, y, &cancelButton) && State == common.StateSaveDialog:
			State = common.StateMain
			saveDialog.SetActive(false)
		case common.Collide(x, y, &exitButton):
			os.Exit(0)
		}
	}
	prevMousePressed = mouseButtonPressed
}
