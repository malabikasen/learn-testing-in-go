package shapes

type Rectangle struct {
	Width  float64
	Height float64
}

func Perimeter(height float64, width float64) float64 {
	return (height + width) * 2
}

func Area(height float64, width float64) float64 {
	return height * width
}
