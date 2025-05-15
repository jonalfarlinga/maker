package game

import (
	"image/color"
	"maker/common"
	"maker/mapdata"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}
var mapData *mapdata.MapArray = mapdata.NewMapArray(int(resolutionBarX.GetValue()), int(resolutionBarY.GetValue()))
var State int = common.StateMain

func (g *Game) Update() error {
	// Update game logic here
	falloffProbBar.Update()
	numberOfIslandsBar.Update()
	resolutionBarX.BoundUpdate(resolutionBarY)
	resolutionBarY.Update()
	fillinBar.Update()
	mouseUpdate()
	if State == common.StateSaveDialog {
		saveDialog.Update()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	mapData.RenderMap(screen)
	drawHUD(screen)
	saveMenu(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(common.ScreenWidth), int(common.ScreenHeight)
}

func drawHUD(screen *ebiten.Image) {
	// Draw buttons
	exitButton.Draw(screen)
	generateButton.Draw(screen)
	savePNG.Draw(screen)
	terraformLakesButton.Draw(screen)
	// Draw the control bars
	falloffProbBar.Draw(screen)
	numberOfIslandsBar.Draw(screen)
	resolutionBarX.Draw(screen)
	resolutionBarY.Draw(screen)
	fillinBar.Draw(screen)
}

func saveMenu(screen *ebiten.Image) {
	if State == common.StateSaveDialog {
		screen.Fill(common.DisabledOverlay)
		saveDialog.Draw(screen)
		saveButton.Draw(screen)
		cancelButton.Draw(screen)
	}
}
