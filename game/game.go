package game

import (
	"image/color"
	"maker/common"
	"maker/mapdata"
	"math/rand"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct{}

var exitButton common.Button = common.Button{
	Width:  150,
	Height: 50,
	X:      float32(common.ScreenWidth - 200),
	Y:      float32(common.ScreenHeight - 100),
	Text:   "Exit",
	Active: true,
}

var generateButton common.Button = common.Button{
	Width:  150,
	Height: 50,
	X:      50,
	Y:      50,
	Text:   "Generate",
	Active: true,
}

var terraformLakesButton common.Button = common.Button{
	Width:  150,
	Height: 50,
	X:      float32(common.ScreenWidth - 200),
	Y:      50,
	Text:   "Terraform Lakes",
	Active: true,
}

var falloffProbBar *mapdata.MapControl = mapdata.NewMapControl(
	50, 150, 200,
	1.0, 100.0,
	"Landmass Expansion",
)

var numberOfIslandsBar *mapdata.MapControl = mapdata.NewMapControl(
	50, 200, 200,
	1.0, 10.0,
	"Number of Islands",
)

var resolutionBarX *mapdata.MapControl = mapdata.NewMapControl(
	50, 250, 200,
	25.0, 720.0,
	"Resolution X",
)

var resolutionBarY *mapdata.MapControl = mapdata.NewMapControl(
	50, 300, 200,
	25.0, 720.0,
	"Resolution Y",
)

var fillinBar *mapdata.MapControl = mapdata.NewMapControl(
	common.ScreenWidth - 250, 150, 200,
	0.0, 15.0,
	"Lake Suppresion",
)

var prevMousePressed bool = false
var mapData *mapdata.MapArray = mapdata.NewMapArray(int(resolutionBarX.GetValue()), int(resolutionBarY.GetValue()))

func (g *Game) Update() error {
	// Update game logic here
	falloffProbBar.Update()
	numberOfIslandsBar.Update()
	resolutionBarX.BoundUpdate(resolutionBarY)
	resolutionBarY.Update()
	fillinBar.Update()
	mouseButtonPressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	if !mouseButtonPressed && prevMousePressed {
		x, y := ebiten.CursorPosition()
		if common.Collide(x, y, &generateButton) {
			// Generate button logic
			falloffProb := falloffProbBar.GetValue() * float32(max(mapData.Width, mapData.Height))/2
			mapData = mapdata.NewMapArray(int(resolutionBarX.GetValue()), int(resolutionBarY.GetValue()))
			border := min(mapData.Height, mapData.Width) / 8
			n := int(numberOfIslandsBar.GetValue())
			mapData.GenerateIsland(mapData.Width/2, mapData.Height/2, int(falloffProb))
			for i := 1; i < n; i++ {
				x := border + rand.Intn(border*6)
				y := border + rand.Intn(border*6)
				mapData.GenerateIsland(x, y, int(falloffProb))
			}
		} else if common.Collide(x, y, &terraformLakesButton) {
			mapData.TerraformLakes(int(fillinBar.GetValue()))
		} else if common.Collide(x, y, &exitButton) {
			os.Exit(0)
		}
	}
	prevMousePressed = mouseButtonPressed
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)
	mapData.RenderMap(screen)
	// Draw buttons
	exitButton.Draw(screen)
	generateButton.Draw(screen)
	terraformLakesButton.Draw(screen)
	// Draw the control bars
	falloffProbBar.Draw(screen)
	numberOfIslandsBar.Draw(screen)
	resolutionBarX.Draw(screen)
	resolutionBarY.Draw(screen)
	fillinBar.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(common.ScreenWidth), int(common.ScreenHeight)
}
