package p14681

import "testing"

func Test12And5QuadrantIs1(t *testing.T) {
	quadrant := getQuadrant(12, 5)

	if quadrant != 1 {
		t.Error("12 And 5 Is 1 Quadrant")
	}
}

func TestU12And5QuadrantIs1(t *testing.T) {
	quadrant := getQuadrant(-12, 5)

	if quadrant != 2 {
		t.Error("-12 And 5 Is 2 Quadrant")
	}
}

func TestU12AndU5QuadrantIs1(t *testing.T) {
	quadrant := getQuadrant(-12, -5)

	if quadrant != 3 {
		t.Error("-12 And -5 Is 3 Quadrant")
	}
}

func Test12AndU5QuadrantIs1(t *testing.T) {
	quadrant := getQuadrant(12, -5)

	if quadrant != 4 {
		t.Error("12 And -5 Is 4 Quadrant")
	}
}