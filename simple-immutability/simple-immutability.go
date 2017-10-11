package simpleimmutability

// LateralScalar represents a quantity in the east-west direction.
type LateralScalar int

// LongitudinalScalar represents a quantity in the north-south direction.
type LongitudinalScalar int

// Coordinates represents the position of an object on a 2D cartesian surface with the origin in the top-left.
// The value for X positively increases towards the right and the value for Y positively increases towards the bottom.
type Coordinates struct {
	X LateralScalar
	Y LongitudinalScalar
}

func (c Coordinates) TranslateNorth(distance LongitudinalScalar) *Coordinates {
	c.Y = c.Y - distance

	return &c
}

func (c Coordinates) TranslateEast(distance LateralScalar) *Coordinates {
	c.X = c.X + distance

	return &c
}

func (c Coordinates) TranslateSouth(distance LongitudinalScalar) *Coordinates {
	c.Y = c.Y + distance

	return &c
}

func (c Coordinates) TranslateWest(distance LateralScalar) *Coordinates {
	c.X = c.X - distance

	return &c
}
