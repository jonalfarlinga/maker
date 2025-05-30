package mapdata

func (m *MapArray) TerraformLakes(fillin int) {
	scope := max(min(m.Width, m.Height)/25, 5) // 5 or 1/25 of the map size
	x1, y1 := 0, 0
	x2, y2 := scope, scope
	// log.Printf("Terraforming lakes with scope %d\n", scope)
	perimeterSize := 2*(scope-1) + 2*(scope-1)
	lakelets := 0
	expandedLakes := 0
	// Scanning function with a sliding window
	scanRow := func(x1, y1, x2, y2 int) {
		// log.Printf("Scanning row %d %d %d %d\n", x1, y1, x2, y2)
		// Calculate the number of perimeter land cells in the first window
		perimeterLand := 0
		for i := x1; i < x2; i++ {
			if m.mapArray[i][y1] == 1 {
				perimeterLand++
			}
			if m.mapArray[i][y2] == 1 {
				perimeterLand++
			}
		}
		for j := y1 + 1; j < y2-1; j++ {
			if m.mapArray[x1][j] == 1 {
				perimeterLand++
			}
			if m.mapArray[x2][j] == 1 {
				perimeterLand++
			}
		}

		// Calculate the number of contained water cells in first window
		containedWater := 0
		for i := 1; i < scope; i++ {
			for j := 1; j < scope; j++ {
				if m.mapArray[i][j] == 0 {
					containedWater++
				}
			}
		}

		for x2 < m.Width {
			// Discriminate between lakes and rivers
			if perimeterLand == perimeterSize && containedWater > 0 {
				// It's a lake
				lakelets++
				disc := r.Intn(scope*scope) * fillin
				if disc > containedWater*containedWater {
					// Fill the lake with land
					for i := x1; i < x2; i++ {
						for j := y1; j < y2; j++ {
							m.mapArray[i][j] = 1
						}
					}
				} else {
					expandedLakes++
					// make a lake
					// Floodfill helper function to create lakes
					makeALake := func(x1, y1, x2, y2 int) {
						visited := make([][]bool, m.Width)
						for i := range visited {
							visited[i] = make([]bool, m.Height)
						}
						// list water cells
						waterCells := make([][2]int, 0)
						for i := x1; i < x2; i++ {
							for j := y1; j < y2; j++ {
								if m.mapArray[i][j] == 0 {
									waterCells = append(waterCells, [2]int{i, j})
									visited[i][j] = true
								}
							}
						}
						// choose a random water cell
						if len(waterCells) == 0 {
							return
						}
						c := r.Intn(len(waterCells))
						x, y := waterCells[c][0], waterCells[c][1]
						visited[x][y] = false

						// flood fill the lake
						queue := make([][2]int, 0)
						queue = append(queue, [2]int{x, y})
						floodFill := func(x, y int) {
							if x < 0 || x >= m.Width || y < 0 || y >= m.Height {
								return
							}
							if visited[x][y] {
								return
							}
							visited[x][y] = true
							m.mapArray[x][y] = 0
							if x > x1 && x < x2-1 && y > y1 && y < y2-1 {
								containedWater--
							}
							if x < x1 || x >= x2 || y < y1 || y >= y2 {
								if r.Intn(10) > 0 {
									return
								}
							} else {
								if r.Intn(6) > 0 {
									return
								}
							}
							queue = append(queue, [2]int{x, y})
							queue = append(queue, [2]int{x - 1, y})
							queue = append(queue, [2]int{x + 1, y})
							queue = append(queue, [2]int{x, y - 1})
							queue = append(queue, [2]int{x, y + 1})
							queue = append(queue, [2]int{x - 1, y - 1})
							queue = append(queue, [2]int{x + 1, y - 1})
							queue = append(queue, [2]int{x - 1, y + 1})
						}
						for len(queue) > 0 {
							floodFill(queue[0][0], queue[0][1])
							queue = queue[1:]
						}
					}
					makeALake(x1, y1, x2, y2)
				}
			}
			if x2 >= m.Width-1 {
				break
			}
			// Move the window
			// Handle corners
			if m.mapArray[x1][y1] == 1 {
				perimeterLand--
			}
			if m.mapArray[x1][y2] == 1 {
				perimeterLand--
			}
			if m.mapArray[x2][y1] == 1 {
				perimeterLand++
			}
			if m.mapArray[x2][y2] == 1 {
				perimeterLand++
			}
			// Remove left column
			for j := y1 + 1; j < y2-1; j++ {
				if m.mapArray[x1][j] == 1 {
					perimeterLand--
				}
			}
			// Move the left interior to left column
			for j := y1 + 1; j < y2-1; j++ {
				if m.mapArray[x1+1][j] == 0 {
					containedWater--
				} else {
					perimeterLand++
				}
			}
			// Move right column to right interior
			for j := y1 + 1; j < y2-1; j++ {
				if m.mapArray[x2][j] == 0 {
					containedWater++
				} else {
					perimeterLand--
				}
			}
			// Add right column
			for j := y1 + 1; j < y2-1; j++ {
				if m.mapArray[x2+1][j] == 1 {
					perimeterLand++
				}
			}
			x1++
			x2++
		}
	}

	// Perform the Terraform
	for y2 < m.Height {
		scanRow(x1, y1, x2, y2)
		x1, x2 = 0, scope
		y1++
		y2++
	}
	// log.Printf("Lakelets: %d, Expanded Lakes: %d\n", lakelets, expandedLakes)
}
