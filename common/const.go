package common

import (
	"image/color"

	"golang.org/x/image/font/basicfont"
)

const (
	ScreenHeight int = 720
	ScreenWidth  int = 1280
	MapRes       int = 25
	SpreadProb   int = 300
)

var (
	ButtonColor      = &color.RGBA{0xFF, 0x00, 0x00, 0xFF} // Red
	// ButtonHoverColor = &color.RGBA{0xFF, 0x82, 0x00, 0xFF} // Orange
	ButtonGlowColor  = &color.RGBA{0xC0, 0xC0, 0x00, 0xFF} // Gold
	ButtonOffColor   = &color.RGBA{0x99, 0x10, 0x10, 0xFF} // Dark Red
	ButtonTextColor  = &color.RGBA{0x00, 0x00, 0x00, 0xFF} // White
	PanelColor       = &color.RGBA{0x8b, 0x8b, 0x8b, 0xFF} // Grey
	DisabledOverlay  = &color.RGBA{0x00, 0x00, 0x00, 0x30} // Black Transparent
	WaterColor       = &color.RGBA{0x30, 0x30, 0xFF, 0xFF} // Blue
	GroundColor      = &color.RGBA{0xBA, 0x81, 0x45, 0xFF} // Tan
)

var (
	MenuFont = basicfont.Face7x13
)

const (
	StateMain = iota
	StateSaveDialog
)
