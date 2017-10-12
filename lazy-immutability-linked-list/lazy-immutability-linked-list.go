package lazyimmutabilitylinkedlist

type coordinatesLinkedList struct {
	operation CoordinateOperation
	previous  *coordinatesLinkedList
}

func (i *coordinatesLinkedList) execute(input *Coordinates) error {
	if i.previous != nil {
		err := i.previous.execute(input)

		if err != nil {
			return err
		}
	}

	return i.operation(input)
}

// LateralScalar represents a quantity in the east-west direction.
type LateralScalar int

// LongitudinalScalar represents a quantity in the north-south direction.
type LongitudinalScalar int

type CoordinateOperation func (input *Coordinates) error

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
	}

	return transform
}

type CoordinateTransform struct {
	input         *Coordinates
	lastOperation *coordinatesLinkedList
}

func (transform CoordinateTransform) Execute() *Coordinates {
	output := *transform.input

	if transform.lastOperation != nil {
		transform.lastOperation.execute(&output)
	}

	return &output
}

func (transform CoordinateTransform) TranslateNorth(distance LongitudinalScalar) CoordinateTransform {
	newOperation := &coordinatesLinkedList{}
	newOperation.previous = transform.lastOperation
	newOperation.operation = func (input *Coordinates) error {
		input.Y = input.Y - distance

		return nil
	}

	transform.lastOperation = newOperation

	return transform
}

func (transform CoordinateTransform) TranslateEast(distance LateralScalar) CoordinateTransform {
	newOperation := &coordinatesLinkedList{}
	newOperation.previous = transform.lastOperation
	newOperation.operation = func (input *Coordinates) error {
		input.X = input.X + distance

		return nil
	}

	transform.lastOperation = newOperation

	return transform
}

func (transform CoordinateTransform) TranslateSouth(distance LongitudinalScalar) CoordinateTransform {
	newOperation := &coordinatesLinkedList{}
	newOperation.previous = transform.lastOperation
	newOperation.operation = func (input *Coordinates) error {
		input.Y = input.Y + distance

		return nil
	}

	transform.lastOperation = newOperation

	return transform
}

func (transform CoordinateTransform) TranslateWest(distance LateralScalar) CoordinateTransform {
	newOperation := &coordinatesLinkedList{}
	newOperation.previous = transform.lastOperation
	newOperation.operation = func (input *Coordinates) error {
		input.X = input.X - distance

		return nil
	}

	transform.lastOperation = newOperation

	return transform
}
