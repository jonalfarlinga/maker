package game

import (
	"image/color"
	"maker/common"
	"maker/common/components"
	"maker/mapdata"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	buttons  []common.Interactable
	controls []components.MapControl
}

var mapData *mapdata.MapArray = mapdata.NewMapArray(int(resolutionBarX.GetValue()), int(resolutionBarY.GetValue()))
var State int = common.StateMain

func (g *Game) Update() error {
	// Update game logic here
	falloffProbBar.Update()
	numberOfIslandsBar.Update()
	resolutionBarX.BoundUpdate(resolutionBarY)
	resolutionBarY.Update()
	fillinBar.Update()
	for _, control := range g.controls {
		control.Update()
	}
	mouseUpdate()
	saveUpdate()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	mapData.RenderMap(screen)
	for _, control := range g.controls {
		control.Draw(screen)
	}
	for _, button := range g.buttons {
		button.Draw(screen)
	}
	drawHUD(screen)
	drawSaveMenu(screen)
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
	smoothLandformsButton.Draw(screen)
	settlementsButton.Draw(screen)
	// Draw the control bars
	falloffProbBar.Draw(screen)
	numberOfIslandsBar.Draw(screen)
	resolutionBarX.Draw(screen)
	resolutionBarY.Draw(screen)
	fillinBar.Draw(screen)
}
