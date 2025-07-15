package mapdata

import (
	c "maker/common"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type MapArray struct {
	Height   int
	Width    int
	mapArray [][]int
	visited  [][]bool
}

var r *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

var cardinal [8][2]int = [8][2]int{
	{0, -1},  // North
	{1, -1},  // North East
	{1, 0},   // East
	{1, 1},   // South East
	{0, 1},   // South
	{-1, 1},  // South West
	{-1, 0},  // West
	{-1, -1}, // North West
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
		Height:   height,
		Width:    width,
	}
}

func (m *MapArray) RenderMap(screen *ebiten.Image) {
	rectSize := float32(c.ScreenHeight) / float32(max(m.Width, m.Height))
	origin := (c.ScreenWidth - c.ScreenHeight) / 2
	for i := 0; i < m.Width; i++ {
		for j := 0; j < m.Height; j++ {
			color := c.WaterColor
			if m.mapArray[i][j] == 1 {
				// Draw Ground
				color = c.GroundColor
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

func (m *MapArray) GenerateIsland(x, y, prob int) {
	m.mapArray[x][y] = 1
	m.visited[x][y] = true
	if x < int(float32(m.Width)*c.EdgeTaperingFactorMax) || x > int(float32(m.Width)*c.EdgeTaperingFactorMin) ||
		y < int(float32(m.Height)*c.EdgeTaperingFactorMax) || y > int(float32(m.Height)*c.EdgeTaperingFactorMin) {
		prob = min(prob, int(float32(min(x, y, m.Width-x, m.Height-y))/c.EdgeTaperingFactorMin))
	}

	offset := r.Intn(8)
	for i := 0; i < 8; i++ {
		dx := x + cardinal[(i+offset)%8][0]
		dy := y + cardinal[(i+offset)%8][1]
		if dx < 0 || dy < 0 || dx > m.Width-1 || dy > m.Height-1 {
			continue
		}
		if m.visited[dx][dy] {
			continue
		}
		if r.Intn(100) < prob {
			m.GenerateIsland(dx, dy, prob-1)
		}
	}
}

func (m *MapArray) ResetMap() {
	for i := 0; i < m.Width; i++ {
		for j := 0; j < m.Height; j++ {
			m.mapArray[i][j] = 0
			m.visited[i][j] = false
		}
	}
}
