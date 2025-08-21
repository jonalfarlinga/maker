package game

import (
	c "maker/common"
	"maker/common/components"
)

var settlementsButton components.Button = components.Button{
	Width:  3 * c.UITileUnit,
	Height: c.UITileUnit,
	X:      c.UITileUnit,
	Y:      7 * c.UITileUnit,
	Text:   "Settlements",
	Active: true,
	Color:  c.ButtonColor,
}

var exitButton components.Button = components.Button{
	Width:  3 * c.UITileUnit,
	Height: c.UITileUnit,
	X:      float32(c.ScreenWidth - 4*int(c.UITileUnit)),
	Y:      float32(c.ScreenHeight - 2*int(c.UITileUnit)),
	Text:   "Exit",
	Active: true,
	Color:  c.ButtonColor,
}

var generateButton components.Button = components.Button{
	Width:  3 * c.UITileUnit,
	Height: c.UITileUnit,
	X:      c.UITileUnit,
	Y:      c.UITileUnit,
	Text:   "Generate",
	Active: true,
	Color:  c.ButtonColor,
}

var terraformLakesButton components.Button = components.Button{
	Width:  3 * c.UITileUnit,
	Height: c.UITileUnit,
	X:      float32(c.ScreenWidth) - 4*c.UITileUnit,
	Y:      c.UITileUnit,
	Text:   "Terraform Lakes",
	Active: true,
	Color:  c.ButtonColor,
}

var smoothLandformsButton components.Button = components.Button{
	Width:  3 * c.UITileUnit,
	Height: c.UITileUnit,
	X:      float32(c.ScreenWidth) - 4*c.UITileUnit,
	Y:      4 * c.UITileUnit,
	Text:   "Smooth Landforms",
	Active: true,
	Color:  c.ButtonColor,
}

var falloffProbBar *components.MapControl = components.NewMapControl(
	c.UITileUnit, 3*c.UITileUnit, 4*c.UITileUnit,
	c.LandmassExpansionMinValue, c.LandmassExpansionMaxValue,
	"Landmass Expansion",
)

var numberOfIslandsBar *components.MapControl = components.NewMapControl(
	c.UITileUnit, 4*c.UITileUnit, 4*c.UITileUnit,
	c.NumberOfIslandsMinValue, c.NumberOfIslandsMaxValue,
	"Number of Islands",
)

var resolutionBarX *components.MapControl = components.NewMapControl(
	c.UITileUnit, 5*c.UITileUnit, 4*c.UITileUnit,
	c.ResolutionMinValue, c.ResolutionMaxValue,
	"Resolution X",
)

var resolutionBarY *components.MapControl = components.NewMapControl(
	c.UITileUnit, 6*c.UITileUnit, 4*c.UITileUnit,
	c.ResolutionMinValue, c.ResolutionMaxValue,
	"Resolution Y",
)

var fillinBar *components.MapControl = components.NewMapControl(
	float32(c.ScreenWidth)-5*c.UITileUnit, 3*c.UITileUnit, 4*c.UITileUnit,
	c.LakeSuppressionMinValue, c.LakeSuppressionMaxValue,
	"Lake Suppresion",
)
