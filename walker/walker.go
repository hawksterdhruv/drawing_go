package walker

import (
	"github.com/gopxl/pixel/v2"
	"github.com/gopxl/pixel/v2/ext/imdraw"
)

type Walker struct {
	CurrentPos      pixel.Vec
	width           float64
	height          float64
	walkingStrategy WalkingStrategy
}

func NewWalker(width, height float64, walkingStrtegy WalkingStrategy) *Walker {
	// startingPosition := pixel.V(rand.Float64()*width, rand.Float64()*height)
	startingPosition := pixel.V(width/2, height/2)
	walker := Walker{
		CurrentPos:      startingPosition,
		width:           width,
		height:          height,
		walkingStrategy: walkingStrtegy,
	}
	return &walker
}

func (walker *Walker) NextStep() {
	walker.walkingStrategy.nextStep(walker)
}
func (walker Walker) Draw(imd *imdraw.IMDraw, prevPos pixel.Vec) {
	imd.Push(walker.CurrentPos, prevPos)
	imd.Line(2)
	// imd.Circle(1, 0)
}
