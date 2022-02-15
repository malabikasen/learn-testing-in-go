package area

import "testing"

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g, wanted %g", got, want)
		}
	}
	t.Run("get area for rectangle", func(t *testing.T) {
		r := Rectangle{4.0, 5.0}
		checkArea(t, r, 20.0)
	})

	t.Run("get area for circle", func(t *testing.T) {
		c := Circle{10.0}
		checkArea(t, c, 314.2)
	})

	t.Run("get area for triangle", func(t *testing.T) {
		tri := Triangle{2.0, 2.0}
		checkArea(t, tri, 2.0)
	})

}
