package comparer

import (
	"strings"
	"testing"
)

func TestCompare_Same(t *testing.T) {
	c := New()
	base := "test1\ntest2\ntest3\n"
	target := "test1\ntest2\ntest3\n"
	if !c.Compare(strings.NewReader(base), strings.NewReader(target)) {
		t.Error("Returned false, but texts are same.")
	}
}

func TestCompare_Different(t *testing.T) {
	c := New()
	base := "test1\ntest2\ntest3\n"
	target := "test2\ntest2\ntest3\n"
	if c.Compare(strings.NewReader(base), strings.NewReader(target)) {
		t.Error("Returned true, but texts are not same.")
	}
}

func TestCompare_SameWithoutIgnorePattern(t *testing.T) {
	c := New()
	c.AddIgnorePattern(`[0-9]{3}`)
	base := "abc123def\nabc456def\nabc789def\n"
	target := "abc987def\nabc654def\nabc321def\n"
	if !c.Compare(strings.NewReader(base), strings.NewReader(target)) {
		t.Error("Returned false, but texts are same without ignore pattern.")
	}
}

func TestCompare_DifferentWithoutIgnorePattern(t *testing.T) {
	c := New()
	c.AddIgnorePattern(`[0-9]{3}`)
	base := "abc123def\nabc456def1\nabc789def\n"
	target := "abc987def\nabc654def2\nabc321def\n"
	if c.Compare(strings.NewReader(base), strings.NewReader(target)) {
		t.Error("Returned false, but texts are not same without ignore pattern.")
	}
}

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
