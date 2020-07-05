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
	circle := Circle{10.0}
	actual := circle.Area()
	expect := 100 * math.Pi
	if actual != expect {
		t.Errorf("actual '%.2f', but expect '%.2f'", actual, expect)
	}
}
