package structs

import "testing"

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("长方形", func(t *testing.T) {
		r := Rectangle{10.0, 10.0}
		checkPerimeter(t, r, 40.0)
	})

	t.Run("圆", func(t *testing.T) {
		c := Circle{10.0}
		checkPerimeter(t, c, 62.83185307179586)
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

	t.Run("长方形", func(t *testing.T) {
		r := Rectangle{10.0, 10.0}
		checkArea(t, r, 100.0)
	})

	t.Run("圆", func(t *testing.T) {
		c := Circle{10.0}
		checkArea(t, c, 314.1592653589793)
	})
}
