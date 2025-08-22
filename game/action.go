package game

import (
	c "maker/common"
	"maker/mapdata"
	"maker/settlements"
	"math/rand"
)

func placeNewSettlement(x, y int) {
	gridX, gridY := mapData.PixToGrid(x, y)
	c.DebugPrintln("mouse", "Placing new settlement at:", x, y, " pixels, ", gridX, gridY, " cell")
	settlements.SettlementsList.NewSettlement(gridX, gridY)
}

func handleGenerateIsland() {
	mapData = mapdata.NewMapArray(
		int(resolutionBarX.GetValue()), int(resolutionBarY.GetValue()))
	falloffProb := (falloffProbBar.GetValue() *
		float32(
			max(mapData.Width, mapData.Height)/
				c.WorldFillReductionFactor))
	border := min(mapData.Height, mapData.Width) / c.WorldEdgeFraction
	n := int(numberOfIslandsBar.GetValue())
	mapData.GenerateIsland(mapData.Width/2, mapData.Height/2, int(falloffProb))
	for i := 1; i < n; i++ {
		x := border + rand.Intn(border*(c.WorldEdgeFraction-2))
		y := border + rand.Intn(border*(c.WorldEdgeFraction-2))
		mapData.GenerateIsland(x, y, int(falloffProb))
	}
	c.DebugPrintln("mouse", "GenerateButton clicked")
}
