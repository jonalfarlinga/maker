package mapdata

import (
	"log"
	"maker/common"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type MapArray struct {
	height   int
	width    int
	mapArray [][]int
	visited  [][]bool
}

func NewMapArray(width, height int) *MapArray {
	m := make([][]int, width)
	for i := range m {
		m[i] = make([]int, height)
	}
	v := make([][]bool, width)
	for i := range v {
		v[i] = make([]bool, height)
	}
	return &MapArray{
		mapArray: m,
		visited:  v,
		height:   height,
		width:    width,
	}
}

func (m *MapArray) RenderMap(screen *ebiten.Image) {
	rectSize := float32(common.ScreenHeight) / float32(common.MapRes)
	origin := (common.ScreenWidth - common.ScreenHeight) / 2
	for i := 0; i < common.MapRes; i++ {
		for j := 0; j < common.MapRes; j++ {
			color := common.WaterColor
			if m.mapArray[i][j] == 1 {
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

func (m *MapArray) GenerateIsland(x int, y int, prob int) {
	if x < 0 || y < 0 || x > common.MapRes-1 || y > common.MapRes-1 {
		return
	}
	if m.visited[x][y] {
		return
	}
	m.mapArray[x][y] = 1
	m.visited[x][y] = true
	if x < common.MapRes/6 || x > common.MapRes*5/6 || y < common.MapRes/6 || y > common.MapRes*5/6 {
		log.Printf("prob: %d", prob)
		prob = min(prob, expansionEdgeFalloff(x, y))
		log.Printf("falloff: %d", prob)
	}
	// North
	if rand.Intn(100) < prob {
		m.GenerateIsland(x, y-1, prob-common.MapRes/50)
	}
	// North East
	if rand.Intn(100) < prob {
		m.GenerateIsland(x+1, y-1, prob-common.MapRes/50)
	}
	// East
	if rand.Intn(100) < prob {
		m.GenerateIsland(x+1, y, prob-common.MapRes/50)
	}
	// South East
	if rand.Intn(100) < prob {
		m.GenerateIsland(x+1, y+1, prob-common.MapRes/50)
	}
	// South
	if rand.Intn(100) < prob {
		m.GenerateIsland(x, y+1, prob-common.MapRes/50)
	}
	// South West
	if rand.Intn(100) < prob {
		m.GenerateIsland(x-1, y+1, prob-common.MapRes/50)
	}
	// West
	if rand.Intn(100) < prob {
		m.GenerateIsland(x-1, y, prob-common.MapRes/50)
	}
	// North West
	if rand.Intn(100) < prob {
		m.GenerateIsland(x-1, y-1, prob-common.MapRes/50)
	}
}

func (m *MapArray) ResetMap() {
	for i := 0; i < m.width; i++ {
		for j := 0; j < m.height; j++ {
			m.mapArray[i][j] = 0
			m.visited[i][j] = false
		}
	}
}

func expansionEdgeFalloff(x, y int) int {
	dist := float32(math.Sqrt(float64(
		float32(x-common.MapRes/2)*float32(x-common.MapRes/2) +
			float32(y-common.MapRes/2)*float32(y-common.MapRes/2)),
	))
	falloff := float32(1) - (dist / float32(common.MapRes/2))
	return int(max(30, falloff))
}
