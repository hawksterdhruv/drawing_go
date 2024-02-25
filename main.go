package main

import (
	"log"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/gopxl/pixel/v2/ext/imdraw"
	"golang.org/x/image/colornames"

	wl "hawksterdhruv/drawing_go/walker"
)

const (
	HEIGHT = 200
	WIDTH  = 300
)

func run() {
	cfg := opengl.WindowConfig{
		Title:  "random walker",
		Bounds: pixel.R(0, 0, WIDTH, HEIGHT),
		VSync:  true,
	}

	win, err := opengl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	imd := imdraw.New(nil)
	walker := wl.NewWalker(WIDTH, HEIGHT)

	for !win.Closed() {
		win.Clear(colornames.Darkslategray)

		walker.Update()
		walker.Draw(imd)

		imd.Draw(win)
		win.Update()
	}
}

func main() {
	log.Println("Hello World")
	opengl.Run(run)
}
