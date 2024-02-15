package api

import (
	tm "github.com/buger/goterm"
)

type Animation struct {
	W, H   int
	Buffer [][]string
	debug  string
}

func (a *Animation) Init(w, h int) {
	a.W, a.H = tm.GetXY(w, h)
	a.Buffer = make([][]string, a.H)
	for y := 0; y < a.H; y++ {
		a.Buffer[y] = make([]string, a.W)
		for x := 0; x < a.W; x++ {
			a.Buffer[y][x] = " "
		}
	}
}

func (a *Animation) Width() int {
	return a.W
}

func (a *Animation) Height() int {
	return a.H
}

func (a Animation) Draw() [][]string {
	return a.Buffer
}

func (a *Animation) Log(msg string) {
	a.debug = msg
}

func (a *Animation) Debug() string {
	return a.debug
}
