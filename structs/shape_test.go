package structs

import "testing"

func TestPrimeter(t *testing.T) {
	actual := Primeter(10.0, 10.0)
	expect := 40.0
	if actual != expect {
		t.Errorf("actual '%.2f', but expect '%.2f'", actual, expect)
	}
}
