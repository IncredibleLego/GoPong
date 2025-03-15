package objects

type Object struct {
	X, Y, W, H int // X, Y, Width, Height. (0,0) is top-left.
}

func (o *Object) HitboxHit(other *Object) bool {
	return o.X < other.X+other.W &&
		o.X+o.W > other.X &&
		o.Y < other.Y+other.H &&
		o.Y+o.H > other.Y
}
