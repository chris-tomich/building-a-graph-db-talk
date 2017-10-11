package simpleimmutability

import "testing"

func TestCoordinates(t *testing.T) {
	a1 := &Coordinates{X: 10, Y: 10}
	a2 := a1.TranslateNorth(1)

	b1 := a2.TranslateEast(2)
	b2 := b1.TranslateNorth(2)

	c1 := a2.TranslateWest(2)
	c2 := c1.TranslateSouth(2)

	if a1.X != 10 && a1.Y != 10 {
		t.Fail()
	}

	if a2.X != 10 && a2.Y != 9 {
		t.Fail()
	}

	if b1.X != 12 && b1.Y != 9 {
		t.Fail()
	}

	if b2.X != 12 && b2.Y != 7 {
		t.Fail()
	}

	if c1.X != 8 && c1.Y != 9 {
		t.Fail()
	}

	if c2.X != 8 && c2.Y != 11 {
		t.Fail()
	}
}
