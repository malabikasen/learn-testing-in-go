package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{4.0, 5.0}
	expectedResult := (rectangle.Width + rectangle.Height) * 2
	actualResult := Perimeter(rectangle)

	if expectedResult != actualResult {
		t.Errorf("got %.2f, wanted %.2f", actualResult, expectedResult)
	}
}

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{4.0, 5.0}
		expectedResult := rectangle.Width * rectangle.Height
		actualResult := rectangle.Area()
		if expectedResult != actualResult {
			t.Errorf("got %.2f, wanted %.2f", actualResult, expectedResult)
		}
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		expectedResult := 314.2
		actualResult := circle.Area()
		if expectedResult != actualResult {
			t.Errorf("got %g, wanted %g", actualResult, expectedResult)
		}
	})

}
