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

func TestCompare_Same_WithIgnorePattern(t *testing.T) {
	c := New()
	c.AddIgnorePattern(`[0-9]{3}`)
	base := "abc123def\nghi456jkl\nmno789pqr\n"
	target := "abc987def\nghi654jkl\nmno321pqr\n"
	if !c.Compare(strings.NewReader(base), strings.NewReader(target)) {
		t.Error("Returned false, but texts are same without ignore pattern.")
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

func TestCompare_Different_DifferentOrder(t *testing.T) {
	c := New()
	base := "test1\ntest2\ntest3\n"
	target := "test2\ntest3\ntest1\n"
	if c.Compare(strings.NewReader(base), strings.NewReader(target)) {
		t.Error("Returned true, but texts are not same.")
	}
}

func TestCompare_Different_WithIgnorePattern(t *testing.T) {
	c := New()
	c.AddIgnorePattern(`[0-9]{3}`)
	base := "abc123def\nghi456jkl\nmno789pqr\n"
	target := "abc987def\nghi654jkl2\nmno321pqr\n"
	if c.Compare(strings.NewReader(base), strings.NewReader(target)) {
		t.Error("Returned false, but texts are not same without ignore pattern.")
	}
}

func TestCompareFreeOrder_Same(t *testing.T) {
	c := New()
	base := "test1\ntest2\ntest3\n"
	target := "test1\ntest2\ntest3\n"
	if !c.CompareFreeOrder(strings.NewReader(base), strings.NewReader(target)) {
		t.Error("Returned false, but texts are same.")
	}
}

func TestCompareFreeOrder_Same_DifferentOrder(t *testing.T) {
	c := New()
	base := "test1\ntest2\ntest3\n"
	target := "test3\ntest1\ntest2\n"
	if !c.CompareFreeOrder(strings.NewReader(base), strings.NewReader(target)) {
		t.Error("Returned false, but texts are same.")
	}
}

func TestCompareFreeOrder_Same_WithIgnorePattern(t *testing.T) {
	c := New()
	c.AddIgnorePattern(`[0-9]{3}`)
	base := "abc123def\nghi456jkl\nmno789pqr\n"
	target := "abc987def\nghi654jkl\nmno321pqr\n"
	if !c.CompareFreeOrder(strings.NewReader(base), strings.NewReader(target)) {
		t.Error("Returned false, but texts are same without ignore pattern.")
	}
}

func TestCompareFreeOrder_Same_WithIgnorePattern_FreeOrder(t *testing.T) {
	c := New()
	c.AddIgnorePattern(`[0-9]{3}`)
	base := "abc123def\nghi456jkl\nmno789pqr\n"
	target := "mno321pqr\nabc987def\nghi654jkl\n"
	if !c.CompareFreeOrder(strings.NewReader(base), strings.NewReader(target)) {
		t.Error("Returned false, but texts are same without ignore pattern.")
	}
}

func TestCompareFreeOrder_Different(t *testing.T) {
	c := New()
	base := "test1\ntest2\ntest3\n"
	target := "test2\ntest2\ntest3\n"
	if c.CompareFreeOrder(strings.NewReader(base), strings.NewReader(target)) {
		t.Error("Returned true, but texts are not same.")
	}
}

func TestCompareFreeOrder_Different_WithIgnorePattern(t *testing.T) {
	c := New()
	c.AddIgnorePattern(`[0-9]{3}`)
	base := "abc123def\nghi456jkl\nmno789pqr\n"
	target := "abc987def\nghi654jkl2\nmno321pqr\n"
	if c.CompareFreeOrder(strings.NewReader(base), strings.NewReader(target)) {
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

func TestCompareLine_Same_WithIgnorePattern(t *testing.T) {
	c := New()
	c.AddIgnorePattern(`[0-9]{3}`)
	if !c.CompareLine("abc123def", "abc456def") {
		t.Error("Returned false, but lines are same without ignore pattern.")
	}
}

func TestCompareLine_Different_WithIgnorePattern(t *testing.T) {
	c := New()
	c.AddIgnorePattern(`[0-9]{3}`)
	if c.CompareLine("abc123def1", "abc456def2") {
		t.Error("Returned true, but lines are not same without ignore pattern.")
	}
}
