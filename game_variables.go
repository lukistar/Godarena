package Godarena

// Everything on the map that collides should be on this slice
// Example: Unit steps on field with fire on it.
// The object representing the fire could be added to the COLLIDERS.
var COLLIDERS []Collider

var MAP Map

var MapX int
var MapY int

// Every tile has a chance to appear.
// When a tile appears, it changes tha chances of its surroundings.
var ChanceTable [][]int
var TilesCount int
var DIE int

// What kind of damage each element causes to others elements.
// These numbers are actually %, i.e 50 means 50 %.
var DamageTable = [7][7]int{
	[7]int{0, 0, 0, 0, 0, 0, 0},         // CHAOS
	[7]int{0, 100, 0, 0, 0, 0, 10},      // LIGHT
	[7]int{0, 0, 100, 200, 10, 50, 200}, // FIRE
	[7]int{0, 0, 200, 100, 90, 0, 10},   // WATER
	[7]int{0, 0, 200, 10, 100, 0, 150},  // EARTH
	[7]int{0, 0, 50, 10, 10, 100, 80},   // WIND
	[7]int{0, 0, 20, 20, 10, 5, 199},    // ORGANIC
}

// What stats gives each point of every element.
// It's flat.
// MASS ENERGY RELEASE CDENERGY(nanos) CDMOVE(nanos) UNABSORB
var StatsTable = [7][6]int{
	[6]int{0, 0, 0, 0, 0, 42},          // CHAOS
	[6]int{0, 1, 1, -10000, -5000, 0},  // LIGHT
	[6]int{0, 6, 3, -1000, -500, -1},   // FIRE
	[6]int{2, 2, 1, -2000, -1000, 0},   // WATER
	[6]int{18, 1, 0, 1000, 1000, 0},    // EARTH
	[6]int{1, 1, 1, -10000, -10000, 0}, // WIND
	[6]int{4, 3, 1, -2000, -2000, 1},   // ORGANIC
}
