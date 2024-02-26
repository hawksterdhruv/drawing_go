package main

import (
	"log"

	mv "hawksterdhruv/drawing_go/mover"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/backends/opengl"
	"github.com/gopxl/pixel/v2/ext/imdraw"
	"golang.org/x/image/colornames"
)

const (
	HEIGHT = 1000
	WIDTH  = 1000
)

func collision(m *mv.Mover) {
	if m.Position.Y <= 0 {
		m.Position.Y = m.Size
		m.Velocity.Y = -m.Velocity.Y
	}
	if m.Position.X <= 0 || m.Position.X > WIDTH {
		m.Velocity.X = -m.Velocity.X
	}
}

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

	// WRITE SETUP CODE HERE
	// TODO: have a method for setup and draw. Figure out a way to pass variables between them.
	m := mv.Mover{
		Position:     pixel.V(WIDTH/2, HEIGHT/2),
		Velocity:     pixel.V(0, 0),
		Acceleration: pixel.V(0, 0),
		Size:         10,
	}
	// END SETUP CODE
	leftCount, rightCount := 0, 0
	for !win.Closed() {
		win.Clear(colornames.Darkslategray)

		// WRITE UPDATE CODE HERE
		m.Update()
		// if win.JustPressed(pixel.MouseButtonLeft) {
		// 	m.ApplyForce(pixel.V(0, -0.1))
		// 	log.Println("pressed : ", m.Acceleration, m.Velocity)

		// }
		if win.Pressed(pixel.KeyRight) {
			m.ApplyForce(pixel.V(0.01, 0))
			leftCount++
			log.Println(leftCount)
		}
		if win.Pressed(pixel.KeyLeft) {
			m.ApplyForce(pixel.V(-0.01, 0))
			rightCount++
			log.Println(rightCount)
		}
		m.ApplyForce(pixel.V(0, -0.1))

		collision(&m)
		m.Draw(imd)
		// END UPDATE CODE
		imd.Draw(win)
		imd.Clear()
		win.Update()
	}
}

func main() {
	log.Println("Hello World")
	opengl.Run(run)
}
