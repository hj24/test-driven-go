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
	areaTestCase := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "circle", shape: Circle{radius: 10}, want: 314.1592653589793},
		{name: "rectangle", shape: Rectangle{height: 10.0, width: 10.0}, want: 100.0},
		{name: "triangle", shape: Triangle{base: 10.0, height: 4.0}, want: 20.0},
	}

	for _, testCase := range areaTestCase {
		t.Run(testCase.name, func(t *testing.T) {
			got := testCase.shape.Area()
			if got != testCase.want {
				t.Errorf("%#v got %.2f want %.2f", testCase.shape, got, testCase.want)
			}
		})
	}
}
