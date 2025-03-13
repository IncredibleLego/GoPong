package objects

type Ball struct {
	*Ojbect
	Dxdt int // x velocity per tick
	Dydt int // y velocity per tick
}

func (b *Ball) Move() { // Move the ball
	b.X += b.Dxdt
	b.Y += b.Dydt
}
