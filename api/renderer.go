package api

import (
	"strings"

	tm "github.com/buger/goterm"
)

type Drawer interface {
	Draw() [][]string
	Debug() string
}

type Renderer struct {
	W, H   int
	buffer [][]string
}

func NewRenderer(w, h int) *Renderer {
	buffer := make([][]string, h)
	for y := 0; y < h; y++ {
		buffer[y] = make([]string, w)
		for x := 0; x < w; x++ {
			buffer[y][x] = " "
		}
	}
	return &Renderer{w, h, buffer}
}

func (f Renderer) RenderBox() {
	tm.MoveCursor(0, 0)
	tm.Println("┌" + strings.Repeat("─", f.W) + "┐")
	for i := 0; i < f.H; i++ {
		tm.Println("│" + strings.Repeat(" ", f.W) + "│")
	}
	tm.Println("└" + strings.Repeat("─", f.W) + "┘")
}

func (f *Renderer) Render(s Drawer) {
	newbuffer := s.Draw()
	if len(newbuffer) != f.H || len(newbuffer[0]) != f.W {
		panic("buffer size does not match field size")
	}
	for y := 0; y < f.H; y++ {
		for x := 0; x < f.W; x++ {
			if f.buffer[y][x] != newbuffer[y][x] {
				tm.MoveCursor(x+2, y+2)
				tm.Printf(newbuffer[y][x])
				tm.MoveCursor(1, f.H+4)
				dbg := s.Debug()
				if len(dbg) > 0 {
					tm.Printf("DEBUG: %s", dbg)
				}
				tm.MoveCursor(1, f.H+5)
				f.buffer[y][x] = newbuffer[y][x]
			}
		}
	}
	tm.Flush()
}

func (f *Renderer) Reset() {
	tm.Clear()
}
