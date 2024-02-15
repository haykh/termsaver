package pipes

import (
	"math/rand"

	"github.com/haykh/termsaver/api"
)

type Dir int

const (
	North Dir = iota
	East
	South
	West
)

func StringDir(d Dir) string {
	switch d {
	case North:
		return "^"
	case South:
		return "v"
	case East:
		return ">"
	default:
		return "<"
	}
}

func NextDir(d Dir) Dir {
	if d == West {
		return []Dir{North, South}[rand.Intn(2)]
	} else if d == East {
		return []Dir{North, South}[rand.Intn(2)]
	} else if d == North {
		return []Dir{East, West}[rand.Intn(2)]
	} else {
		return []Dir{East, West}[rand.Intn(2)]
	}
}

var pipes = map[Dir]map[Dir]string{
	North: {
		North: "│",
		East:  "┌",
		West:  "┐",
	},
	East: {
		North: "┘",
		East:  "─",
		South: "┐",
	},
	South: {
		East:  "└",
		South: "│",
		West:  "┘",
	},
	West: {
		North: "└",
		South: "┌",
		West:  "─",
	},
}

type Pipes struct {
	api.Animation
	x0, y0 int
	dir    Dir
	prob   float64
}

func New(w, h int) api.Termsaver {
	p := Pipes{}
	p.Animation.Init(w, h)
	p.x0 = p.W / 2
	p.y0 = p.H / 2
	p.dir = []Dir{North, East, South, West}[rand.Intn(4)]
	p.prob = 0.0
	return &p
}

func (p *Pipes) Update() {
	p.Log(StringDir(p.dir))
	prev_dir := p.dir
	if rand.Float64() < p.prob {
		p.dir = NextDir(p.dir)
		p.prob = 0.0
	} else {
		p.prob += 0.2
	}
	p.Buffer[p.y0][p.x0] = pipes[prev_dir][p.dir]

	switch p.dir {
	case North:
		p.y0--
	case East:
		p.x0++
	case South:
		p.y0++
	case West:
		p.x0--
	}
	if p.x0 < 0 {
		p.x0 = p.W - 1
	} else if p.x0 >= p.W {
		p.x0 = 0
	}
	if p.y0 < 0 {
		p.y0 = p.H - 1
	} else if p.y0 >= p.H {
		p.y0 = 0
	}
}
