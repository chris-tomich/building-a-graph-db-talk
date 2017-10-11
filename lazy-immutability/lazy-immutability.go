package lazyimmutability

// LateralScalar represents a quantity in the east-west direction.
type LateralScalar int

// LongitudinalScalar represents a quantity in the north-south direction.
type LongitudinalScalar int

type CoordinateOperation func (input *Coordinates)

// Coordinates represents the position of an object on a 2D cartesian surface with the origin in the top-left.
// The value for X positively increases towards the right and the value for Y positively increases towards the bottom.
type Coordinates struct {
	X LateralScalar
	Y LongitudinalScalar
}

func (c Coordinates) TranslateNorth(distance LongitudinalScalar) CoordinateTransform {
	transform := NewCoordinateTransform(&c)
	return transform.TranslateNorth(distance)
}

func (c Coordinates) TranslateEast(distance LateralScalar) CoordinateTransform {
	transform := NewCoordinateTransform(&c)
	return transform.TranslateEast(distance)
}

func (c Coordinates) TranslateSouth(distance LongitudinalScalar) CoordinateTransform {
	transform := NewCoordinateTransform(&c)
	return transform.TranslateSouth(distance)
}

func (c Coordinates) TranslateWest(distance LateralScalar) CoordinateTransform {
	transform := NewCoordinateTransform(&c)
	return transform.TranslateWest(distance)
}

func NewCoordinateTransform(input *Coordinates) CoordinateTransform {
	transform := CoordinateTransform{
		input: input,
		operations: make([]CoordinateOperation, 0, 10),
	}

	return transform
}

type CoordinateTransform struct {
	input      *Coordinates
	operations []CoordinateOperation
}

func (transform CoordinateTransform) Execute() *Coordinates {
	output := *transform.input

	for _, op := range transform.operations {
		op(&output)
	}

	return &output
}

func (transform CoordinateTransform) TranslateNorth(distance LongitudinalScalar) CoordinateTransform {
	newOps := make([]CoordinateOperation, len(transform.operations), len(transform.operations) + 1)

	copy(newOps, transform.operations)
	transform.operations = append(newOps, func (input *Coordinates) {
		input.Y = input.Y - distance
	})

	return transform
}

func (transform CoordinateTransform) TranslateEast(distance LateralScalar) CoordinateTransform {
	newOps := make([]CoordinateOperation, len(transform.operations), len(transform.operations) + 1)

	copy(newOps, transform.operations)
	transform.operations = append(newOps, func (input *Coordinates) {
		input.X = input.X + distance
	})

	return transform
}

func (transform CoordinateTransform) TranslateSouth(distance LongitudinalScalar) CoordinateTransform {
	newOps := make([]CoordinateOperation, len(transform.operations), len(transform.operations) + 1)

	copy(newOps, transform.operations)
	transform.operations = append(newOps, func (input *Coordinates) {
		input.Y = input.Y + distance
	})

	return transform
}

func (transform CoordinateTransform) TranslateWest(distance LateralScalar) CoordinateTransform {
	newOps := make([]CoordinateOperation, len(transform.operations), len(transform.operations) + 1)

	copy(newOps, transform.operations)
	transform.operations = append(newOps, func (input *Coordinates) {
		input.X = input.X - distance
	})

	return transform
}
