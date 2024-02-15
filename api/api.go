package api

import (
	"time"
)

type Termsaver interface {
	Init(w, h int)
	Update()
	Draw() [][]string
	Width() int
	Height() int
	Debug() string
}

func Tick(New func(int, int) Termsaver, w, h int, dt time.Duration) {
	ts := New(w, h)
	renderer := NewRenderer(ts.Width(), ts.Height())
	renderer.Reset()
	renderer.RenderBox()
	for {
		ts.Update()
		renderer.Render(ts)
		time.Sleep(dt)
	}
}
