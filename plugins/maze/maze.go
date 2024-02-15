package maze

import (
	"math/rand"

	"github.com/haykh/termsaver/api"
)

type Maze struct {
	api.Animation
}

func Generate(buff *([][]string), x0, y0, w, h int) {
	if (w <= 3) || (h <= 3) {
		return
	}
	dxc := rand.Intn(w)
	dyc := rand.Intn(h)
	// fmt.Printf("w = %d; h = %d; dxc = %d; dyc = %d; h - dyc = %d; w - dxc = %d\n", w, h, dxc, dyc, h-dyc, w-dxc)
	for y := y0; y < y0+h; y++ {
		(*buff)[y][x0+dxc] = "*"
	}
	for x := x0; x < x0+w; x++ {
		(*buff)[y0+dyc][x] = "*"
	}
	uncut_wall := rand.Intn(4)
	for wall := 0; wall < 4; wall++ {
		if wall == uncut_wall {
			continue
		}
		var xh, yh int
		switch wall {
		case 0:
			xh = x0 + dxc
			yh = y0 + dyc + 1 + rand.Intn(h-dyc-1)
		case 1:
			xh = x0 + dxc + 1 + rand.Intn(w-dxc-1)
			yh = y0 + dyc
		case 2:
			xh = x0 + dxc
			yh = y0 + rand.Intn(dyc)
		case 3:
			xh = x0 + rand.Intn(dxc)
			yh = y0 + dyc
		default:
			panic("why am i here")
		}
		(*buff)[yh][xh] = " "
	}
	Generate(buff, x0, y0, dxc, dyc)
	Generate(buff, x0+dxc+1, y0, w-dxc-1, dyc)
	Generate(buff, x0, y0+dyc+1, dxc, h-dyc-1)
	Generate(buff, x0+dxc+1, y0+dyc+1, w-dxc-1, h-dyc-1)
}

func New(w, h int) api.Termsaver {
	m := Maze{}
	m.Animation.Init(w, h)
	Generate(&m.Buffer, 0, 0, m.Width(), m.Height())
	return &m
}

func (m *Maze) Update() {
}
