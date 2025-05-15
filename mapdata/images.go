package mapdata

import (
	"image"
	"image/png"
	"maker/common"
	"os"
)

func (m *MapArray) OutputPNG(filename string) {
	upLeft := image.Point{0, 0}
	lowRight := image.Point{m.Width, m.Height}

	img := image.NewRGBA(image.Rectangle{upLeft, lowRight})

	// Set color for each pixel
	for x := 0; x < m.Width; x++ {
		for y := 0; y < m.Height; y++ {
			switch {
			case m.mapArray[x][y] == 1:
				// Draw Ground
				img.Set(x, y, common.GroundColor)
			case m.mapArray[x][y] == 0:
				// Draw Water
				img.Set(x, y, common.WaterColor)
			default:
				// Use zero value.
			}
		}
	}

	// Encode as PNG.
	f, _ := os.Create("saved_maps/" + filename)
	png.Encode(f, img)
}
