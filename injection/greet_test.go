package injection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("word exist in dict", func(t *testing.T) {
		writer := &bytes.Buffer{}
		Greet(writer, "Ken")
		actual := writer.String()
		expect := "Hello, Ken"
		assertMessageEquals(t, actual, expect)
	})

}

func assertMessageEquals(t *testing.T, actual, expect string) {
	t.Helper()
	if actual != expect {
		t.Errorf("actual '%s', but expect '%s'", actual, expect)
	}
}
