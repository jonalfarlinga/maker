package common

import (
	"image/color"

	"golang.org/x/image/font/basicfont"
)

const (
	ScreenHeight int = 720
	ScreenWidth  int = 1280
	MapRes int = 250
	SpreadProb int = 300
)

var (
	ButtonColor      = &color.RGBA{0xFF, 0x00, 0x00, 0xFF} // Red
	ButtonHoverColor = &color.RGBA{0xFF, 0x30, 0x00, 0xFF} // Orange
	ButtonOffColor   = &color.RGBA{0x99, 0x10, 0x10, 0xFF} // Dark Red
	ButtonTextColor  = &color.RGBA{0x00, 0x00, 0x00, 0xFF} // White
	WaterColor	   = &color.RGBA{0x00, 0x00, 0xFF, 0xFF} // Blue
	GroundColor	   = &color.RGBA{0x00, 0xFF, 0x00, 0xFF} // Green
)

var (
	MenuFont = basicfont.Face7x13
)
