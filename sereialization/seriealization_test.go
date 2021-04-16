package sereialization

import (
	"fmt"
	"testing"
)

func TestSeriealize(t *testing.T) {
	type Student struct {
		name string
		age  int
	}
	t.Run("test", func(t *testing.T) {
		var students = []Student{
			{name: "Alice", age: 23},
			{name: "Bob", age: 25},
		}

		fmt.Println(students)

		str := Seriealize(students)
		fmt.Println(str)
	})
}
