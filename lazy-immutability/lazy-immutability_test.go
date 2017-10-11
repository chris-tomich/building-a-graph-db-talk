package lazyimmutability

import "testing"

func TestCoordinates(t *testing.T) {
	a1 := &Coordinates{X: 10, Y: 10}
	a2 := a1.TranslateNorth(1)

	b1 := a2.TranslateEast(2)
	b2 := b1.TranslateNorth(2)

	c1 := a2.TranslateWest(2)
	c2 := c1.TranslateSouth(2)

	a2Result := a2.Execute()
	b1Result := b1.Execute()
	b2Result := b2.Execute()
	c2Result := c2.Execute()
	c1Result := c1.Execute()

	if a1.X != 10 && a1.Y != 10 {
		t.Fail()
	}

	if a2Result.X != 10 && a2Result.Y != 9 {
		t.Fail()
	}

	if b1Result.X != 12 && b1Result.Y != 9 {
		t.Fail()
	}

	if b2Result.X != 12 && b2Result.Y != 7 {
		t.Fail()
	}

	if c1Result.X != 8 && c1Result.Y != 9 {
		t.Fail()
	}

	if c2Result.X != 8 && c2Result.Y != 11 {
		t.Fail()
	}
}
