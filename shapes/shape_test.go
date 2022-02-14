package shapes

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{4.0, 5.0}
	expectedResult := (rectangle.Width + rectangle.Height) * 2
	actualResult := Perimeter(rectangle)

	if expectedResult != actualResult {
		t.Errorf("got %.2f, wanted %.2f", actualResult, expectedResult)
	}
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g, wanted %g", got, want)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{4.0, 5.0}
		checkArea(t, rectangle, 20.0)
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.2)
	})

}

func Test_Area(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12.0, 6.0}, 72.0},
		{Circle{10.0}, 314.2},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g, wanted %g", got, tt.want)
		}
	}
}
