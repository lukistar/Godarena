package Godarena

import "math/rand"
import "time"

// Return Id of Tile to be placed into MAP.arrays.
func GenerateField(x, y int) int {
	sum := 0
	for i := 0; i < TilesCount; i++ {
		sum += MAP.chances[x][y][i]
	}
	rand.Seed(time.Now().Unix())
	randomed := RandomInt(0, sum)
	sum2 := 0
	for i := 0; i < TilesCount; i++ {
		sum2 += MAP.chances[x][y][i]
		if randomed < sum2 {
			return i
		}
	}
	return TilesCount // Who knows... just in case. 
}

func RandomInt(min, max int) int {
	return rand.Intn(max - min) + min
}