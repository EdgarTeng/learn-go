package iteration

import "testing"

func TestRepeat(t *testing.T) {
	actual := Repeat("a", 5)
	expect := "aaaaa"
	assertEqual(t, actual, expect)
}

func BenchmarkRepeat(t *testing.B) {
	for i := 0; i < t.N; i++ {
		Repeat("a", 5)
	}
}

func assertEqual(t *testing.T, actual, expect interface{}) {
	if actual != expect {
		t.Errorf("actual '%d', expect '%d'", actual, expect)
	}
}
