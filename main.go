package main

import (
	"time"

	tm "github.com/buger/goterm"
	"github.com/haykh/termsaver/api"
	"github.com/haykh/termsaver/plugins/maze"
)

func main() {
	// api.Tick(pipes.New, 60|tm.PCT, 80|tm.PCT, 50*time.Millisecond)
	api.Tick(maze.New, 60|tm.PCT, 80|tm.PCT, 50*time.Millisecond)
}
