package strs_test

import (
	"testing"

	"github.com/pdk/gotil/strs"
)

func TestSetList(t *testing.T) {

	var x strs.SetList

	if x.Len() != 0 {
		t.Errorf("expected size of brand new SetList to be 0, got %d", x.Len())
	}

	x.Append("foo")

	if x.Len() != 1 {
		t.Errorf("expected size of SetList to be 1, got %d", x.Len())
	}

	x.Remove("bar")

	if x.Len() != 1 {
		t.Errorf("expected size of SetList to be 1, got %d", x.Len())
	}

	x.Append("bar", "cat")

	if x.Len() != 3 {
		t.Errorf("expected size of SetList to be 3, got %d", x.Len())
	}

	x.Remove("bar")

	if x.Len() != 2 {
		t.Errorf("expected size of SetList to be 2, got %d", x.Len())
	}

	var y strs.SetList

	y.Append("foo", "cat")

	if !x.Equal(y) {
		t.Errorf("expected x (%v) and y (%v) to be equal", x, y)
	}

	g, ok := y.Pull()

	if !ok {
		t.Errorf("expected Pull operation to succeed, got false")
	}

	if g != "cat" {
		t.Errorf("expected to pull cat, but got %s", g)
	}

	if x.Equal(y) {
		t.Errorf("expected x to NOT equal y")
	}

	y.Pull()

	y.Append("cat", "foo")

	if x.Equal(y) {
		t.Errorf("expected x to not equal y")
	}

	if !x.Match(y) {
		t.Errorf("expected x to Match y")
	}
}
