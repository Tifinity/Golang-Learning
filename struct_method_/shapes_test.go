package shape

import (
	"testing"
	"math"
)

func TestPerimeter(t *testing.T) {

	checkPerimeter := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
            t.Errorf("got %.2f want %.2f", got, want)
        }
	}

    t.Run("rectangles", func(t *testing.T) {
        rectangle := Rectangle{12, 6}
        want := 36.0
        checkPerimeter(t, rectangle, want)   
    })

    t.Run("circles", func(t *testing.T) {
        circle := Circle{10}
        want := 2 * math.Pi * 10
        checkPerimeter(t, circle, want)  
    })
}

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
            t.Errorf("got %.2f want %.2f", got, want)
        }
	}

    t.Run("rectangles", func(t *testing.T) {
        rectangle := Rectangle{12, 6}
        want := 72.0
        checkArea(t, rectangle, want)   
    })

    t.Run("circles", func(t *testing.T) {
        circle := Circle{10}
        want := 314.1592653589793
        checkArea(t, circle, want)  
    })

}
