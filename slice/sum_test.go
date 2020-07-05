package slice

import "testing"

func TestSum(t *testing.T) {

	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		actual := Sum(numbers)
		expect := 15
		assertEquals(t, expect, actual)
	})

	t.Run("Collection of any numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
		actual := Sum(numbers)
		expect := 45
		assertEquals(t, expect, actual)
	})
}

func assertEquals(t *testing.T, expect, actual int) {
	if actual != expect {
		t.Errorf("expect '%d', but actual '%d'", expect, actual)
	}
}

func BenchmarkSum(t *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < t.N; i++ {
		Sum(numbers)
	}
}
