package mocking

import (
	"fmt"
	"io"
)

// Countdown means count to 0
func Countdown(writer io.Writer) {
	for i := 3; i >= 1; i-- {
		fmt.Fprintln(writer, i)
	}
	fmt.Fprint(writer, "GO!")
}
