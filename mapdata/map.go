package mapdata

import (
	"maker/common"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

var mapArray [common.MapRes][common.MapRes]int = [common.MapRes][common.MapRes]int{}
var visited [common.MapRes][common.MapRes]bool = [common.MapRes][common.MapRes]bool{}

func RenderMap(screen *ebiten.Image) {
	rectSize := float32(common.ScreenHeight) / float32(common.MapRes)
	origin := (common.ScreenWidth - common.ScreenHeight) / 2
	for i := 0; i < common.MapRes; i++ {
		for j := 0; j < common.MapRes; j++ {
			color := common.WaterColor
			if mapArray[i][j] == 1 {
				// Draw Ground
				color = common.GroundColor
				// } else if mapArray[i][j] == 2 {
				// 	// Draw tree
			}
			vector.DrawFilledRect(
				screen,
				float32(origin)+float32(i)*rectSize, float32(j)*rectSize,
				rectSize, rectSize,
				color, false,
			)
		}
	}
}

func GenerateIsland(x int, y int, prob int) {
	if x < 0 || y < 0 || x > common.MapRes-1 || y > common.MapRes-1 {
		return
	}
	if visited[x][y] {
		return
	}
	if x < common.MapRes/10 || x > common.MapRes*9/10 || y < common.MapRes/10 || y > common.MapRes*9/10 {
		prob -= 5
	}
	mapArray[x][y] = 1
	visited[x][y] = true
	// North
	if rand.Intn(100) < prob {
		GenerateIsland(x, y-1, prob-1)
	}
	// North East
	if rand.Intn(100) < prob {
		GenerateIsland(x+1, y-1, prob-1)
	}
	// East
	if rand.Intn(100) < prob {
		GenerateIsland(x+1, y, prob-1)
	}
	// South East
	if rand.Intn(100) < prob {
		GenerateIsland(x+1, y+1, prob-1)
	}
	// South
	if rand.Intn(100) < prob {
		GenerateIsland(x, y+1, prob-1)
	}
	// South West
	if rand.Intn(100) < prob {
		GenerateIsland(x-1, y+1, prob-1)
	}
	// West
	if rand.Intn(100) < prob {
		GenerateIsland(x-1, y, prob-1)
	}
	// North West
	if rand.Intn(100) < prob {
		GenerateIsland(x-1, y-1, prob-1)
	}
}

func ResetMap() {
	mapArray = [common.MapRes][common.MapRes]int{}
	visited = [common.MapRes][common.MapRes]bool{}
}
