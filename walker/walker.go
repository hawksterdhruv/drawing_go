package walker

import (
	"math/rand"

	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/ext/imdraw"
)

type Walker struct {
	currentPos      pixel.Vec
	width           float64
	height          float64
	walkingStrategy WalkingStrategy
}

func NewWalker(width, height float64, walkingStrtegy WalkingStrategy) *Walker {
	startingPosition := pixel.V(rand.Float64()*width, rand.Float64()*height)
	walker := Walker{
		currentPos:      startingPosition,
		width:           width,
		height:          height,
		walkingStrategy: walkingStrtegy,
	}
	return &walker
}

func (walker *Walker) NextStep() {
	walker.walkingStrategy.nextStep(walker)
}
func (walker Walker) Draw(imd *imdraw.IMDraw) {
	imd.Push(walker.currentPos)
	imd.Circle(1, 0)
}
