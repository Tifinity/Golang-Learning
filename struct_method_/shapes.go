package shape

type Rect struct {
	width  float64
	height float64
}

// Permeter is p
func Perimeter(rect Rect) float64 {
	return 2 * (rect.width + rect.height)
}

// Area is a
func Area(rect Rect) float64 {
	return rect.width * rect.height
}
