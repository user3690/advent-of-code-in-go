package util

import "errors"

type Point struct {
	X int // rows
	Y int // cols
}

func (p Point) Move(dir Direction, steps uint32) (Point, error) {
	switch dir {
	case Up:
		return Point{
			X: p.X - int(steps),
			Y: p.Y,
		}, nil
	case Right:
		return Point{
			X: p.X,
			Y: p.Y + int(steps),
		}, nil
	case Down:
		return Point{
			X: p.X + int(steps),
			Y: p.Y,
		}, nil
	case Left:
		return Point{
			X: p.X,
			Y: p.Y - int(steps),
		}, nil
	default:
		return Point{}, errors.New("unknown movement direction given")
	}
}
