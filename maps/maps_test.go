package maps

import "testing"

func TestSearch(t *testing.T) {
	assertStrings := func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got %q want %q given, %q", got, want, "test")
		}
	}
	dictionary := map[string]string{"test": "this is just a test"}

	got := Search(dictionary, "test")
	want := "this is just a test"
	assertStrings(t, got, want)
}
