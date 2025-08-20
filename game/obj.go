package game

import (
	c "maker/common"
	"maker/mapdata"
)

var exitButton c.Button = c.Button{
	Width:  3*c.UITileUnit,
	Height: c.UITileUnit,
	X:      float32(c.ScreenWidth - 4*int(c.UITileUnit)),
	Y:      float32(c.ScreenHeight - 2*int(c.UITileUnit)),
	Text:   "Exit",
	Active: true,
}

var generateButton c.Button = c.Button{
	Width:  3*c.UITileUnit,
	Height: c.UITileUnit,
	X:      c.UITileUnit,
	Y:      c.UITileUnit,
	Text:   "Generate",
	Active: true,
}

var terraformLakesButton c.Button = c.Button{
	Width:  3*c.UITileUnit,
	Height: c.UITileUnit,
	X:      float32(c.ScreenWidth) - 4*c.UITileUnit,
	Y:      c.UITileUnit,
	Text:   "Terraform Lakes",
	Active: true,
}

var smoothLandformsButton c.Button = c.Button{
	Width:  3*c.UITileUnit,
	Height: c.UITileUnit,
	X:      float32(c.ScreenWidth) - 4*c.UITileUnit,
	Y:      4*c.UITileUnit,
	Text:   "Smooth Landforms",
	Active: true,
}

var falloffProbBar *mapdata.MapControl = mapdata.NewMapControl(
	c.UITileUnit, 3*c.UITileUnit, 4*c.UITileUnit,
	c.LandmassExpansionMinValue, c.LandmassExpansionMaxValue,
	"Landmass Expansion",
)

var numberOfIslandsBar *mapdata.MapControl = mapdata.NewMapControl(
	c.UITileUnit, 4*c.UITileUnit, 4*c.UITileUnit,
	c.NumberOfIslandsMinValue, c.NumberOfIslandsMaxValue,
	"Number of Islands",
)

var resolutionBarX *mapdata.MapControl = mapdata.NewMapControl(
	c.UITileUnit, 5*c.UITileUnit, 4*c.UITileUnit,
	c.ResolutionMinValue, c.ResolutionMaxValue,
	"Resolution X",
)

var resolutionBarY *mapdata.MapControl = mapdata.NewMapControl(
	c.UITileUnit, 6*c.UITileUnit, 4*c.UITileUnit,
	c.ResolutionMinValue, c.ResolutionMaxValue,
	"Resolution Y",
)

var fillinBar *mapdata.MapControl = mapdata.NewMapControl(
	float32(c.ScreenWidth)-5*c.UITileUnit, 3*c.UITileUnit, 4*c.UITileUnit,
	c.LakeSuppressionMinValue, c.LakeSuppressionMaxValue,
	"Lake Suppresion",
)
