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
		// m.Position.Y = m.Size
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
	m1 := mv.Mover{
		Position:     pixel.V(WIDTH/2, HEIGHT/2),
		Velocity:     pixel.V(0, 0),
		Acceleration: pixel.V(0, 0),
		Size:         10,
	}
	m2 := mv.Mover{
		Position:     pixel.V(WIDTH/2+100, HEIGHT/2),
		Velocity:     pixel.V(0, 0),
		Acceleration: pixel.V(0, 0),
		Size:         20,
	}
	// END SETUP CODE

	for !win.Closed() {
		win.Clear(colornames.Darkslategray)

		// WRITE UPDATE CODE HERE
		m1.Update()
		m2.Update()
		// if win.JustPressed(pixel.MouseButtonLeft) {
		// 	m.ApplyForce(pixel.V(0, -0.1))
		// 	log.Println("pressed : ", m.Acceleration, m.Velocity)

		// }
		if win.Pressed(pixel.KeyRight) {
			m1.ApplyForce(pixel.V(0.01, 0))
			m2.ApplyForce(pixel.V(0.01, 0))

		}
		if win.Pressed(pixel.KeyLeft) {
			m1.ApplyForce(pixel.V(-0.01, 0))
			m2.ApplyForce(pixel.V(-0.01, 0))

		}
		gravity := pixel.V(0, -0.025)
		weight1 := gravity.Scaled(m1.Size)
		weight2 := gravity.Scaled(m2.Size)
		m1.ApplyForce(weight1)
		m2.ApplyForce(weight2)

		collision(&m1)
		collision(&m2)

		m1.Draw(imd)
		m2.Draw(imd)
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
