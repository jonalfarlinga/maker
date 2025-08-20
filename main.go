package main

import (
	"flag"
	"image"
	"maker/assets"
	"maker/common"
	"maker/game"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	var debugFiles string
	flag.StringVar(&debugFiles, "debug", "", "comma-separated list of files to activate debug logging for (or 'all')")
	flag.Parse()

	for _, f := range strings.Split(debugFiles, ",") {
		f = strings.TrimSpace(f)
		if f != "" {
			common.DebugFiles[f] = true
		}
	}
	g := &game.Game{}

	// Set window size and run in fullscreen
	ebiten.SetWindowSize(int(common.ScreenWidth), int(common.ScreenHeight))
	ebiten.SetFullscreen(false)
	ebiten.SetWindowTitle("Maker")
	ebiten.SetWindowIcon([]image.Image{assets.IconImage})

	err := ebiten.RunGame(g)
	if err != nil {
		panic(err)
	}
}
