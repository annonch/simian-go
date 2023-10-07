package simian

import (
	"testing"
)

func TestNewEngine(t *testing.T) {
	got := NewEngine("test")
	want := &Engine{Name: "test", minDelay: 1}

	if got.Name != want.Name {
		t.Errorf("got %+v, wanted %+v", got, want)

	}
	if got.minDelay != want.minDelay {
		t.Errorf("got %+v, wanted %+v", got, want)
	}
}
