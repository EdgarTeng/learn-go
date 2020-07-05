package structs

import "testing"

func TestPrimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	actual := Primeter(rectangle)
	expect := 40.0
	if actual != expect {
		t.Errorf("actual '%.2f', but expect '%.2f'", actual, expect)
	}
}
