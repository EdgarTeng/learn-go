package maps

import "testing"

func TestSearch(t *testing.T) {
	t.Run("word exist in dict", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a test string"}
		actual := dictionary.Search("test")
		expect := "this is a test string"
		assertMessageEquals(t, actual, expect)
	})

	t.Run("word not exist in dict", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a test string"}
		actual := dictionary.Search("test123")
		expect := ""
		assertMessageEquals(t, actual, expect)
	})
}

func TestAdd(t *testing.T) {
	t.Run("word exist in dict", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a test string"}
		err := dictionary.Add("test", "this is another test string")
		expect := ErrAlreadyExists
		assertError(t, err, expect)
	})

	t.Run("word not exist in dict", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a test string"}
		err := dictionary.Add("test123", "hhhh")
		assertNoError(t, err)
	})
}

func assertMessageEquals(t *testing.T, actual, expect string) {
	t.Helper()
	if actual != expect {
		t.Errorf("actual '%s', but expect '%s'", actual, expect)
	}
}

func assertError(t *testing.T, actual, expect error) {
	t.Helper()
	if actual != expect {
		t.Errorf("actual '%s', but expect '%s'", actual, expect)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("have error '%s'", err)
	}
}
