package misc

type Rectangle struct {
	Position      Position
	Width, Height int
}

func (r *Rectangle) contains(position Position) bool {
	return position.X >= r.Position.X && position.X <= r.Position.X+r.Width && position.Y >= r.Position.Y && position.Y <= r.Position.Y+r.Height
}
