package walker

import (
	"math"
	"math/rand"

	"github.com/gopxl/pixel/v2"
)

type WalkingStrategy interface {
	nextStep(*Walker)
}

type LevyFlightWalkingStrategy struct {
}

func (l *LevyFlightWalkingStrategy) nextStep(walker *Walker) {
	chosen_angle := 2 * math.Pi * rand.Float64()
	prob := 100 * rand.Float32() // Scaling to 100%
	var chosen_direction pixel.Vec
	step_size := 2.0
	if prob < 1 { // 1% chance
		step_size = 25 + float64(rand.Intn(25))
	}

	chosen_direction = pixel.Unit(chosen_angle).Scaled(step_size)
	walker.CurrentPos = walker.CurrentPos.Add(chosen_direction)
}

type CardinalWalkingStrategy struct {
}

func (l *CardinalWalkingStrategy) nextStep(walker *Walker) {
	chosen_direction := rand.Intn(4)
	switch chosen_direction {
	case 0: // UP
		if walker.CurrentPos.Y < walker.height-1 {
			walker.CurrentPos.Y += 1
		}
	case 1: // RIGHT
		if walker.CurrentPos.X < walker.width-1 {
			walker.CurrentPos.X += 1
		}
	case 2: // DOWN
		if walker.CurrentPos.Y > 0 {
			walker.CurrentPos.Y -= 1
		}
	case 3: // LEFT
		if walker.CurrentPos.X > 0 {
			walker.CurrentPos.X -= 1
		}
	}
}

type OrdinalWalkingStrategy struct {
}

func (l *OrdinalWalkingStrategy) nextStep(walker *Walker) {
	chosen_direction := rand.Intn(8)
	switch chosen_direction {
	case 0: // TOP
		if walker.CurrentPos.Y < walker.height-1 {
			walker.CurrentPos.Y += 1
		}
	case 1: // TOP-RIGHT
		if walker.CurrentPos.Y < walker.height-1 && walker.CurrentPos.X < walker.width-1 {
			walker.CurrentPos.X += 1
			walker.CurrentPos.Y += 1
		}
	case 2: // RIGHT
		if walker.CurrentPos.X < walker.width-1 {
			walker.CurrentPos.X += 1
		}
	case 3: // BOTTOM-RIGHT
		if walker.CurrentPos.X < walker.width-1 && walker.CurrentPos.Y > 0 {
			walker.CurrentPos.X += 1
			walker.CurrentPos.Y -= 1
		}
	case 4: // BOTTOM
		if walker.CurrentPos.Y > 0 {
			walker.CurrentPos.Y -= 1
		}
	case 5: // BOTTOM-LEFT
		if walker.CurrentPos.X > 0 && walker.CurrentPos.Y > 0 {
			walker.CurrentPos.X -= 1
			walker.CurrentPos.Y -= 1
		}
	case 6: // LEFT
		if walker.CurrentPos.X > 0 {
			walker.CurrentPos.X -= 1
		}
	case 7: // TOP-LEFT
		if walker.CurrentPos.X > 0 && walker.CurrentPos.Y < walker.height-1 {
			walker.CurrentPos.X -= 1
			walker.CurrentPos.Y += 1
		}
	}
}
