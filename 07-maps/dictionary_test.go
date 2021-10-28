package dictionary

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("search term in dictionary", func(t *testing.T) {
		got, _ := dictionary.Search("test")

		assertStringsEqual(t, got, "this is just a test")
	})

	t.Run("search term not in dictionary", func(t *testing.T) {
		_, err := dictionary.Search("testy")

		if err == nil {
			t.Fatal("wanted error, got no error")
		} else {
			assertStringsEqual(t, err.Error(), "search term not in dictionary")
		}
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add("test", "this is a test")

		assertError(t, err, nil)

		got, _ := dictionary.Search("test")
		assertStringsEqual(t, got, "this is a test")
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"a": "Apple"}
		err := dictionary.Add("a", "Aardvark")

		assertError(t, err, ErrWordExists)
		got, _ := dictionary.Search("a")
		assertStringsEqual(t, got, "Apple")
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"a": "Apple"}
		err := dictionary.Update("a", "Aardvark")

		assertError(t, err, nil)
		got, _ := dictionary.Search("a")
		assertStringsEqual(t, got, "Aardvark")
	})

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Update("test", "this is a test")

		assertError(t, err, ErrWordNotFound)
	})
}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{"a": "Apple"}
	dictionary.Delete("a")

	_, err := dictionary.Search("a")
	assertError(t, err, ErrWordNotFound)
}

func assertStringsEqual(t testing.TB, got, expected string) {
	t.Helper()
	if got != expected {
		t.Errorf("expected %s, got %s", expected, got)
	}
}

func assertError(t testing.TB, got, expected error) {
	t.Helper()
	if got != expected {
		t.Errorf("expected %q, got %q", expected, got)
	}
}
