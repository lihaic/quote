package quote

import "testing"

func TestHello(t *testing.T) {
	want := "hello"
	if got := Hello(); got != want {
		t.Errorf("got != want, got=%s, want=%s", got, want)
	}
}

func TestHi(t *testing.T) {
	want := "hi"
	if got := Hi(); got != want {
		t.Errorf("got != want, got=%s, want=%s", got, want)
	}
}

func TestDream(t *testing.T) {
	want := "dream is dream"
	if got := Dream("dream"); got != want {
		t.Errorf("got != want, got=%s, want=%s", got, want)
	}
}
