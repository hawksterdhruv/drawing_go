package walker

import (
	"math/rand"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/ext/imdraw"
)

type Walker struct {
	currentPos pixel.Vec
	width      float64
	height     float64
}

func NewWalker(width, height float64) *Walker {
	startingPosition := pixel.V(rand.Float64()*width, rand.Float64()*height)
	walker := Walker{
		currentPos: startingPosition,
		width:      width,
		height:     height,
	}
	return &walker
}

func (walker *Walker) Update() {
	chosen_direction := rand.Intn(4)
	switch chosen_direction {
	case 0: // UP
		if walker.currentPos.Y < walker.height-1 {
			walker.currentPos.Y += 1
		}
	case 1: // RIGHT
		if walker.currentPos.X < walker.width-1 {
			walker.currentPos.X += 1
		}
	case 2: // DOWN
		if walker.currentPos.Y > 0 {
			walker.currentPos.Y -= 1
		}
	case 3: // LEFT
		if walker.currentPos.X > 0 {
			walker.currentPos.X -= 1
		}
	}
	// log.Println(walker.currentPos)
}
func (walker Walker) Draw(imd *imdraw.IMDraw) {
	imd.Push(walker.currentPos)
	imd.Circle(1, 0)
}
