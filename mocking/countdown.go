package main

import (
	"fmt"
	"io"
	"os"
)

// Countdown means count to 0
func Countdown(writer io.Writer) {
	for i := 3; i >= 1; i-- {
		fmt.Fprintln(writer, i)
	}
	fmt.Fprint(writer, "GO!")
}

func main() {
	Countdown(os.Stdout)
}
