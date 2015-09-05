package Godarena

type Coords struct {
	x, y int
}


type Elements struct {
	elements [7]int
	sum int // The sum of all elements but CHAOS.
}

func(e *Elements) AddElement(value int, id int) {
	e.elements[id] += value
	e.sum += value
}

func(e *Elements) AddElements(values [7]int) {
	for i := 0; i < 7; i++ {
		e.elements[i] += values[i]
		e.sum += values[i]
	}
}

func(e *Elements) SetElements(values [7]int) {
	e.sum = 0
	for i := 0; i < 7; i++ {
		e.elements[i] = values[i]
		e.sum += values[i]
	}
}

func(e *Elements) SetElement(value int, id int) {
	e.sum -= e.elements[id]
	e.elements[id] = value
	e.sum += value
}


type Stats struct {
	stats [6]int
}

func(s *Stats) SetStats(values [6]int) {
	for i := 0; i < 6; i++ {
		s.stats[i] = values[i]
	}
}

func(s *Stats) AddStats(value [6]int) {
	for i := 0; i < 6; i++ {
		s.stats[i] += value[i]
	}
}

func(s *Stats) AddStat(value int, id int) {
	s.stats[id] += value
}

func(s* Stats) SetStat(value int, id int) {
	s.stats[id] = value
}