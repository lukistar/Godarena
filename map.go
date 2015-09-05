package Godarena

// Basically an array of ids.
// Every different id represents different Tile.
type Map struct {
	x, y    int
	array   [][]int
	chances [][][]int
}

// Fill an area of the map with values.
// Note: Including xEnd and yEnd!
func (m *Map) Fill(content [][]int, xBegin, xEnd, yBegin, yEnd int) {
	for i := xBegin; i < xEnd; i++ {
		for k := yBegin; k < yEnd; k++ {
			m.array[i][k] = content[i][k]
		}
	}
}

// Set field at position x,y.
func (m *Map) Set(value int, x, y int) {
	m.array[x][y] = value
}

// Should be used as init/restart of the terrain chances
func (m *Map) SetChances(content [][][]int, xBegin, xEnd, yBegin, yEnd int) {
	for i := xBegin; i < xEnd; i++ {
		for k := yBegin; k < yEnd; k++ {
			for z := 0; z < TilesCount; z++ {
				m.chances[i][k][z] = content[i][k][z]
			}
		}
	}
}

// Used to update the terrain appearing chances
func (m *Map) AddChances(content [][][]int, xBegin, xEnd, yBegin, yEnd int) {
	for i := xBegin; i < xEnd; i++ {
		for k := yBegin; k < yEnd; k++ {
			for z := 0; z < TilesCount; z++ {
				m.chances[i][k][z] += content[i][k][z]
			}
		}
	}
}
