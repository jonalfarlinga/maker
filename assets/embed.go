package assets

import (
	"embed"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed "*"
var assetLib embed.FS
var IconImage = mustLoadImage("icon.png")

func mustLoadImage(path string) *ebiten.Image {
	f, err := assetLib.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(img)
}

// func loadIcon(path string) *image.Image {
// 	f, err := assetLib.Open(path)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()
// 	img, _, err := image.Decode(f)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return &img
// }
