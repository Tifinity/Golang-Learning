package shape

import(
    "testing"
)
// TestPerimeter for sdfsdfdsf
func TestPerimeter(t *testing.T) {
    rectangle := Rect{10.0, 10.0}
    got := Perimeter(rectangle)
    want := 40.0

    if got != want {
        t.Errorf("got %.2f but want %.2f", got, want)
    }
}

// TestArea forsdfs dfds
func TestArea(t *testing.T) {
    rectangle := Rect{12.0, 6.0}
    got := Area(rectangle)
    want := 72.0

    if got != want {
        t.Errorf("got %.2f but want %.2f", got, want)
    }
}
