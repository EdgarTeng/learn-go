package integers

import "testing"

func TestAdd(t *testing.T) {
	//define expect & actual
	got := Add(1,2)
	want := 3
	if got != want {
		t.Errorf("got '%q', want '%q'", got, want)
	}
}
