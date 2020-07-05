package integers

import "testing"

func TestAdd(t *testing.T) {
	//define expect & actual
	actual := Add(1, 2)
	expect := 3
	assertEquals(t, actual, expect)
}

func assertEquals(t *testing.T, actual, expect interface{}) {
	if actual != expect {
		t.Errorf("actual '%d', expect '%d'", actual, expect)
	}
}
