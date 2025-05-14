package mapdata

import (
	"maker/common"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type MapArray struct {
	height   int
	width    int
	mapArray [][]int
	visited  [][]bool
}

var r *rand.Rand

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
	r = rand.New(rand.NewSource(time.Now().UnixNano()))
	if x < 0 || y < 0 || x > common.MapRes-1 || y > common.MapRes-1 {
		return
	}
	if m.visited[x][y] {
		return
	}
	m.mapArray[x][y] = 1
	m.visited[x][y] = true

	if x < common.MapRes/10 || x > common.MapRes*9/10 ||
		y < common.MapRes/10 || y > common.MapRes*9/10 {
		prob = min(prob, min(x, y, common.MapRes-x, common.MapRes-y)*10)
	}

	// North
	if r.Intn(100) < prob {
		m.GenerateIsland(x, y-1, prob-5)
	}
	// North East
	if r.Intn(100) < prob {
		m.GenerateIsland(x+1, y-1, prob-5)
	}
	// East
	if r.Intn(100) < prob {
		m.GenerateIsland(x+1, y, prob-5)
	}
	// South East
	if r.Intn(100) < prob {
		m.GenerateIsland(x+1, y+1, prob-5)
	}
	// South
	if r.Intn(100) < prob {
		m.GenerateIsland(x, y+1, prob-5)
	}
	// South West
	if r.Intn(100) < prob {
		m.GenerateIsland(x-1, y+1, prob-5)
	}
	// West
	if r.Intn(100) < prob {
		m.GenerateIsland(x-1, y, prob-5)
	}
	// North West
	if r.Intn(100) < prob {
		m.GenerateIsland(x-1, y-1, prob-5)
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
