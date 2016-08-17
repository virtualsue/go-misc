// author: Jacky Boen
// Messy commenting and so forth done by virtualsue

package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"os"
)

var winTitle string = "Go SDL2 Test"
var winWidth, winHeight int = 800, 600

func run() int {
	var window *sdl.Window
	var renderer *sdl.Renderer

	window, err := sdl.CreateWindow(winTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", err)
		return 1
	}
	defer window.Destroy()

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", err)
		return 2
	}
	defer renderer.Destroy()

	renderer.Clear()

	// RGB black = 0, 0, 0

	renderer.SetDrawColor(0, 0, 255, 255)
	for i := 0; i < 800; i++ {
		var fy float64 = math.Sin(float64(i) * (math.Pi / 180))
		y := (int(fy*100) + 240)

		renderer.DrawPoint(i, y)
	}

//	renderer.SetDrawColor(0, 255, 255, 255)
//	renderer.DrawPoint(250, 300)

	renderer.DrawPoint(350, 300)

	renderer.Present()

	sdl.Delay(20000)

	return 0
}

func main() {
	os.Exit(run())
}
