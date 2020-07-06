package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is a test string"}

	actual := Search(dictionary, "test")
	expect := "this is a test string"
	assertMessageEquals(t, actual, expect)

}

func assertMessageEquals(t *testing.T, actual, expect string) {
	t.Helper()
	if actual != expect {
		t.Errorf("actual '%s', but expect '%s'", actual, expect)
	}
}
