package common

type Rectangle struct {
	Position      Position
	Width, Height int
}

func (r *Rectangle) Contains(position Position) bool {
	// fmt.Println("input Position: ", position)
	// fmt.Println("Rectangle Position: ", r.Position, "Rectangle Width: ", r.Width, "Rectangle Height: ", r.Height)
	return position.X >= r.Position.X && position.X <= r.Position.X+r.Width &&
		position.Y >= r.Position.Y && position.Y <= r.Position.Y+r.Height
}
