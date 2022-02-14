package shapes

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Shape interface {
	Area() float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func Perimeter(rectangle Rectangle) float64 {
	return (rectangle.Height + rectangle.Width) * 2
}

func (c Circle) Area() float64 {
	return 3.142 * c.Radius * c.Radius
}
