package greeting

import "testing"

func TestHello(t *testing.T) {
	cases := map[string]string{
		"":       "Hello, world!",
		"Gopher": "Hello, Gopher!",
	}

	for name, want := range cases {
		if got := Hello(name); got != want {
			t.Errorf("Hello(%q) = %q, want %q", name, got, want)
		}
	}
}
