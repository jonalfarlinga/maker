package main

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
	Y:      50,
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

var falloffProbBar *mapdata.MapControl = mapdata.NewMapControl(
	50, 150, 200,
	float32(common.MapRes), float32(common.MapRes*6),
	"Landmass Expansion",
)

var numberOfIslandsBar *mapdata.MapControl = mapdata.NewMapControl(
	50, 200, 200,
	1.0, 10.0,
	"Number of Islands",
)

var prevMousePressed bool = false
var mapData *mapdata.MapArray = mapdata.NewMapArray(common.MapRes, common.MapRes)

func (g *Game) Update() error {
	// Update game logic here
	falloffProbBar.Update()
	numberOfIslandsBar.Update()
	mouseButtonPressed := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)
	if !mouseButtonPressed && prevMousePressed {
		x, y := ebiten.CursorPosition()
		if common.Collide(x, y, &generateButton) {
			// Generate button logic
			falloffProb := falloffProbBar.GetValue()
			mapData.ResetMap()
			border := common.MapRes / 8
			n := int(numberOfIslandsBar.GetValue())
			mapData.GenerateIsland(common.MapRes/2, common.MapRes/2, int(falloffProb))
			for i := 1; i < n; i++ {
				x := border + rand.Intn(border*6)
				y := border + rand.Intn(border*6)
				mapData.GenerateIsland(x, y, int(falloffProb))
			}
		} else if common.Collide(x, y, &exitButton) {
			os.Exit(0)
		}
	}
	prevMousePressed = mouseButtonPressed
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw game graphics here
	screen.Fill(color.Black)
	mapData.RenderMap(screen)
	exitButton.Draw(screen)
	generateButton.Draw(screen)
	falloffProbBar.Draw(screen)
	numberOfIslandsBar.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(common.ScreenWidth), int(common.ScreenHeight)
}

func main() {
	g := &Game{}

	// Set window size and run in fullscreen
	ebiten.SetWindowSize(int(common.ScreenWidth), int(common.ScreenHeight))
	ebiten.SetFullscreen(true)

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
