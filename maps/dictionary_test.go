package maps

import "testing"

func TestAdd(t *testing.T) {

	t.Run("add new word", func(t *testing.T) {
		dictionary := Dictionary{}
		key := "test"
		value := "this is a test"

		dictionary.Add(key, value)

		assertDefinition(t, dictionary, key, value)
	})

	t.Run("existing word", func(t *testing.T) {
		key := "test"
		value := "this is a test"
		dictionary := Dictionary{key: value}

		err := dictionary.Add(key, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, key, value)
	})
}

func assertDefinition(t *testing.T, d Dictionary, key, value string) {
	t.Helper()

	got, err := d.Search(key)
	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if value != got {
		t.Errorf("got %q want %q", got, value)
	}
}

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("known word", func(t *testing.T) {
		got, err := dictionary.Search("test")
		want := "this is a test"

		assertStrings(t, got, want)
		assertNoError(t, err)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("cat")

		assertError(t, err, ErrNotFound)
	})
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("encountered %q when no error should be found", err)
	}
}
