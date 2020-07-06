package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {

	t.Run("count down", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer)

		actual := buffer.String()
		expect := `3
2
1
GO!`

		assertMessageEquals(t, actual, expect)
	})

}

func assertMessageEquals(t *testing.T, actual, expect string) {
	t.Helper()
	if actual != expect {
		t.Errorf("actual '%s', but expect '%s'", actual, expect)
	}
}
