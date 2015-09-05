package Godarena

type God struct {
	Unit

	unintegrated    [7]int
	unintegratedSum int
}

func (g *God) Absorb() {
	if g.unintegratedSum > g.stats[UNABSORB] {
		g.Integrate()
	}
}

func (g *God) Update() {
	g.Absorb()
}

func (g *God) Integrate() {
	if g.sum <= 0 {
		g.Destroy()
	}
	for i := 0; i < 7; i++ {
		for k := 0; k < 6; k++ {
			g.stats[k] = StatsTable[i][k] * g.elements[i]
		}
	}
}
