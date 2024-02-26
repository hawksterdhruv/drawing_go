package walker

import "math/rand"

type WalkingStrategy interface {
	nextStep(*Walker)
}

type CardinalWalkingStrategy struct {
}

func (l *CardinalWalkingStrategy) nextStep(walker *Walker) {
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
}

type OrdinalWalkingStrategy struct {
}

func (l *OrdinalWalkingStrategy) nextStep(walker *Walker) {
	chosen_direction := rand.Intn(8)
	switch chosen_direction {
	case 0: // TOP
		if walker.currentPos.Y < walker.height-1 {
			walker.currentPos.Y += 1
		}
	case 1: // TOP-RIGHT
		if walker.currentPos.Y < walker.height-1 && walker.currentPos.X < walker.width-1 {
			walker.currentPos.X += 1
			walker.currentPos.Y += 1
		}
	case 2: // RIGHT
		if walker.currentPos.X < walker.width-1 {
			walker.currentPos.X += 1
		}
	case 3: // BOTTOM-RIGHT
		if walker.currentPos.X < walker.width-1 && walker.currentPos.Y > 0 {
			walker.currentPos.X += 1
			walker.currentPos.Y -= 1
		}
	case 4: // BOTTOM
		if walker.currentPos.Y > 0 {
			walker.currentPos.Y -= 1
		}
	case 5: // BOTTOM-LEFT
		if walker.currentPos.X > 0 && walker.currentPos.Y > 0 {
			walker.currentPos.X -= 1
			walker.currentPos.Y -= 1
		}
	case 6: // LEFT
		if walker.currentPos.X > 0 {
			walker.currentPos.X -= 1
		}
	case 7: // TOP-LEFT
		if walker.currentPos.X > 0 && walker.currentPos.Y < walker.height-1 {
			walker.currentPos.X -= 1
			walker.currentPos.Y += 1
		}
	}
}
