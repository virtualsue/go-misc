// author: Sue Spence
// Adapted from a go-sdl2 example program by Jacky Boen
// and Rosetta Code

package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"math/cmplx"
	"os"
)

var winTitle string = "Go Mandelbrot"
var winWidth, winHeight int = 800, 600

const (
	maxEsc = 100
	rMin   = -2.
	rMax   = .5
	iMin   = -1.
	iMax   = 1.
	width  = 750
	red    = 230
	green  = 235
	blue   = 255
)

func mandelbrot(a complex128) float64 {
	i := 0
	for z := a; cmplx.Abs(z) < 2 && i < maxEsc; i++ {
		z = z*z + a
	}
	return float64(maxEsc-i) / maxEsc
}

func run() int {
	var window *sdl.Window
	var renderer *sdl.Renderer

	window, err := sdl.CreateWindow(winTitle,
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight,
		sdl.WINDOW_SHOWN)
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
	scale := width / (rMax - rMin)
	//	height := int(scale * (iMax - iMin))

	for x := 0; x < 800; x++ {
		var fy float64 = mandelbrot(x)
		renderer.DrawPoint(x, y)
	}

	//	renderer.SetDrawColor(0, 255, 255, 255)
	//	renderer.DrawPoint(250, 300)

	renderer.Present()

	sdl.Delay(20000)

	return 0
}

func main() {
	os.Exit(run())
}
