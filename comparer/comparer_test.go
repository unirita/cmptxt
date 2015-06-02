package comparer

import "testing"

func TestCompareLine_Same(t *testing.T) {
	c := New()
	if !c.CompareLine("test", "test") {
		t.Error("Returned false, but lines are same.")
	}
}

func TestCompareLine_Different(t *testing.T) {
	c := New()
	if c.CompareLine("test1", "test2") {
		t.Error("Returned true, but lines are not same.")
	}
}

func TestCompareLine_SameWithoutIgnorePattern(t *testing.T) {
	c := New()
	c.AddIgnorePattern(`[0-9]{3}`)
	if !c.CompareLine("abc123def", "abc456def") {
		t.Error("Returned false, but lines are same without ignore pattern.")
	}
}

func TestCompareLine_DifferentWithoutIgnorePattern(t *testing.T) {
	c := New()
	c.AddIgnorePattern(`[0-9]{3}`)
	if c.CompareLine("abc123def1", "abc456def2") {
		t.Error("Returned true, but lines are not same without ignore pattern.")
	}
}
