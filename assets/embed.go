package assets

// import (
// 	"image"
// 	"embed"
// 	_ "image/png"
// 	"github.com/hajimehoshi/ebiten/v2"
// )

// var assets embed.FS
// var IconImage = mustLoadImage("assets/icon.png")

// func mustLoadImage(path string) *ebiten.Image {
// 	f, err := assets.Open(path)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()
// 	img, _, err := image.Decode(f)
// 	if err != nil {
// 		panic(err)
// 	}
// 	return ebiten.NewImageFromImage(img)
// }
