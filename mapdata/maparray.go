package mapdata

import (
	c "maker/common"
	"maker/settlements"
	r "math/rand"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type MapArray struct {
	Height   int
	Width    int
	mapArray [][]int
	visited  [][]bool
}

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

func (m *MapArray) Draw(screen *ebiten.Image) {
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
	for pos := range settlements.SettlementsList {
		// Draw Settlement
		settSize := max(10, int(rectSize/4))
		vector.DrawFilledCircle(
			screen,
			float32(origin)+float32(pos[0])*rectSize, float32(pos[1])*rectSize,
		float32(settSize)/2, c.SettlementColor, false,
		)
	}
}

func (m *MapArray) SmoothLandforms() {
	window := make([][]int, 3)
	for i := range window {
		window[i] = make([]int, 3)
		for j := range window[i] {
			window[i][j] = m.mapArray[i][j]
		}
	}
	for x := 0; x < m.Width; x++ {
		for y := 0; y < m.Height; y++ {
			// Populate the 3x3 window with the current cell and its neighbors
			for dx := -1; dx <= 1; dx++ {
				for dy := -1; dy <= 1; dy++ {
					nx, ny := x+dx, y+dy
					if nx >= 0 && ny >= 0 && nx < m.Width && ny < m.Height {
						window[dx+1][dy+1] = m.mapArray[nx][ny]
					} else {
						window[dx+1][dy+1] = 0 // Treat out-of-bounds as water
					}
				}
			}
			// Count the number of land cells in the window
			landCount := 0
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					if window[i][j] == 1 {
						landCount++
					}
				}
			}
			// Apply smoothing rules
			if landCount >= 5 {
				m.mapArray[x][y] = 1
			} else {
				m.mapArray[x][y] = 0
			}
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

func (m *MapArray) GetBounds() (float32, float32, float32, float32) {
	// Calculate rectSize and origin as in Draw
	rectSize := float32(c.ScreenHeight) / float32(max(m.Width, m.Height))
	origin := float32(c.ScreenWidth-c.ScreenHeight) / 2

	// The rendered map is always a square of size c.ScreenHeight, centered horizontally
	x0 := origin
	var y0 float32 = 0
	width := rectSize * float32(m.Width)
	height := rectSize * float32(m.Height)
	return x0, y0, width, height
}

func (m *MapArray) PixToGrid(x, y int) (int, int) {
	rectSize := float32(c.ScreenHeight) / float32(max(m.Width, m.Height))
	origin := float32(c.ScreenWidth-c.ScreenHeight) / 2
	gridX := int((float32(x)-origin)/rectSize) + 1
	gridY := int(float32(y)/rectSize) + 1
	return gridX, gridY
}
