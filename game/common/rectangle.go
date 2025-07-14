package common

type Rectangle struct {
	Position      Position
	Width, Height float64
}

func (r *Rectangle) Contains(position Position) bool {
	// fmt.Println("input Position: ", position)
	// fmt.Println("Rectangle Position: ", r.Position, "Rectangle Width: ", r.Width, "Rectangle Height: ", r.Height)
	return position.X >= r.Position.X && position.X <= r.Position.X+r.Width &&
		position.Y >= r.Position.Y && position.Y <= r.Position.Y+r.Height
}

func (r *Rectangle) HorizontalCenter() float64 {
	return r.Position.X + r.Width/2
}

func (r *Rectangle) BottomY() float64 {
	return r.Position.Y + r.Height
}
