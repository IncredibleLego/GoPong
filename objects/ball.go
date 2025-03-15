package objects

type Ball struct {
	*Object
	Dxdt int // x velocity per tick
	Dydt int // y velocity per tick
}

func (b *Ball) Move() { // Move the ball
	b.X += b.Dxdt
	b.Y += b.Dydt
}

func (b *Ball) IncreaseSpeed(increase int) { // Increase the speed of the ball
	b.Dxdt += increase
	b.Dydt += increase
}
