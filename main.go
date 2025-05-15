package main

import (
	"maker/common"
	"maker/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := &game.Game{}

	// Set window size and run in fullscreen
	ebiten.SetWindowSize(int(common.ScreenWidth), int(common.ScreenHeight))
	ebiten.SetFullscreen(false)

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
