package game

import (
	"maker/common"
	"maker/mapdata"
)

var exitButton common.Button = common.Button{
	Width:  150,
	Height: 50,
	X:      float32(common.ScreenWidth - 200),
	Y:      float32(common.ScreenHeight - 100),
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

var terraformLakesButton common.Button = common.Button{
	Width:  150,
	Height: 50,
	X:      float32(common.ScreenWidth - 200),
	Y:      50,
	Text:   "Terraform Lakes",
	Active: true,
}

var savePNG common.Button = common.Button{
	Width:  150,
	Height: 50,
	X:      50,
	Y:      float32(common.ScreenHeight - 200),
	Text:   "Save PNG",
	Active: true,
}

var falloffProbBar *mapdata.MapControl = mapdata.NewMapControl(
	50, 150, 200,
	1.0, 100.0,
	"Landmass Expansion",
)

var numberOfIslandsBar *mapdata.MapControl = mapdata.NewMapControl(
	50, 200, 200,
	1.0, 10.0,
	"Number of Islands",
)

var resolutionBarX *mapdata.MapControl = mapdata.NewMapControl(
	50, 250, 200,
	25.0, 720.0,
	"Resolution X",
)

var resolutionBarY *mapdata.MapControl = mapdata.NewMapControl(
	50, 300, 200,
	25.0, 720.0,
	"Resolution Y",
)

var fillinBar *mapdata.MapControl = mapdata.NewMapControl(
	common.ScreenWidth-250, 150, 200,
	0.0, 15.0,
	"Lake Suppresion",
)
