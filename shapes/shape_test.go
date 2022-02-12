package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{4.0, 5.0}
	expectedResult := (rectangle.Width + rectangle.Height) * 2
	actualResult := Perimeter(rectangle.Width, rectangle.Height)

	if expectedResult != actualResult {
		t.Errorf("got %.2f, wanted %.2f", actualResult, expectedResult)
	}
}

func TestArea(t *testing.T) {
	rectangle := Rectangle{4.0, 5.0}
	expectedResult := rectangle.Width * rectangle.Height
	actualResult := Area(rectangle.Width, rectangle.Height)

	if expectedResult != actualResult {
		t.Errorf("got %.2f, wanted %.2f", actualResult, expectedResult)
	}
}
