package mover

import (
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/ext/imdraw"
)

type Mover struct {
	Position     pixel.Vec
	Velocity     pixel.Vec
	Acceleration pixel.Vec
	Size         float64
}

func (m *Mover) Update() {
	m.Velocity = m.Velocity.Add(m.Acceleration)
	m.Position = m.Position.Add(m.Velocity)
	m.Acceleration = pixel.Vec{X: 0, Y: 0}
}

func (m Mover) Draw(imd *imdraw.IMDraw) {
	imd.Push(m.Position)
	imd.Circle(m.Size, 1)
}

func (m *Mover) ApplyForce(force pixel.Vec) {
	m.Acceleration = m.Acceleration.Add(force.Scaled(1 / m.Size))
}
