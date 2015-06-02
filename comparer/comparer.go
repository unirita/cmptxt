package comparer

import "regexp"

type comparer struct {
	ignorePatterns []*regexp.Regexp
}

func New() *comparer {
	obj := new(comparer)
	obj.ignorePatterns = make([]*regexp.Regexp, 0)
	return obj
}

func (c *comparer) AddIgnorePattern(pattern string) error {
	r, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	c.ignorePatterns = append(c.ignorePatterns, r)
	return nil
}

func (c *comparer) CompareLine(base string, target string) bool {
	for _, ptn := range c.ignorePatterns {
		base = ptn.ReplaceAllString(base, "")
		target = ptn.ReplaceAllString(target, "")
	}
	return base == target
}
