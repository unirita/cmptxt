// Package comparer implements compare functions.
package comparer

import "regexp"

type comparer struct {
	ignorePatterns []*regexp.Regexp
}

// New creates a new comparer.
func New() *comparer {
	obj := new(comparer)
	obj.ignorePatterns = make([]*regexp.Regexp, 0)
	return obj
}

// AddIgnorePattern compiles regexp pattern string and add it to comparer.
// Comparer will ignore part of string which matches added pattern.
//
// If pattern compile was failed, return error object.
// Otherwise, return nil.
func (c *comparer) AddIgnorePattern(pattern string) error {
	r, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	c.ignorePatterns = append(c.ignorePatterns, r)
	return nil
}

// CompareLine compares two line.
// If they are same, return true. Otherwise, return false.
func (c *comparer) CompareLine(base string, target string) bool {
	for _, ptn := range c.ignorePatterns {
		base = ptn.ReplaceAllString(base, "")
		target = ptn.ReplaceAllString(target, "")
	}
	return base == target
}
