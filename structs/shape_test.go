package structs

import (
	"math"
	"testing"
)

func TestPrimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	actual := rectangle.Primeter()
	expect := 40.0
	if actual != expect {
		t.Errorf("actual '%.2f', but expect '%.2f'", actual, expect)
	}
}

func TestArea(t *testing.T) {
	areaTest := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10.0, 20.0}, 200.0},
		{Circle{10.0}, 100 * math.Pi},
	}

	for _, tt := range areaTest {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %.2f want %.2f", got, tt.want)
		}
	}
}
