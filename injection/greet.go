package injection

import (
	"fmt"
	"io"
)

// Greet means say Hello for a name
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}
