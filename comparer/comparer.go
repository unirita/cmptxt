// Package comparer implements compare functions.
package comparer

import (
	"bufio"
	"container/list"
	"io"
	"regexp"
)

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

// Compare compares two texts read from base and target.
func (c *comparer) Compare(base io.Reader, target io.Reader) bool {
	baseScanner := bufio.NewScanner(base)
	targetScanner := bufio.NewScanner(target)
	for {
		isBaseEOF := !baseScanner.Scan()
		isTargetEOF := !targetScanner.Scan()
		if isBaseEOF && isTargetEOF {
			return true
		} else if isBaseEOF || isTargetEOF {
			return false
		}

		if !c.CompareLine(baseScanner.Text(), targetScanner.Text()) {
			return false
		}
	}
}

// CompareFreeOrder compares two texts read from base and target.
// This function ignores line order.
func (c *comparer) CompareFreeOrder(base io.Reader, target io.Reader) bool {
	baseLines := list.New()
	targetLines := list.New()
	defer baseLines.Init()
	defer targetLines.Init()
	c.readAllLines(baseLines, base)
	c.readAllLines(targetLines, target)

	if baseLines.Len() != targetLines.Len() {
		return false
	}

	// Remove lines from targetLineList which exeits in baseLineList
	for b := baseLines.Front(); b != nil; b = b.Next() {
		for t := targetLines.Front(); t != nil; t = t.Next() {
			baseLine, ok1 := b.Value.(string)
			targetLine, ok2 := t.Value.(string)
			if !ok1 || !ok2 {
				panic("Logic error!")
			}
			if c.CompareLine(baseLine, targetLine) {
				targetLines.Remove(t)
				break
			}
		}
	}

	if targetLines.Len() > 0 {
		return false
	}

	return true
}

// CompareLine compares baseLine and targetLine.
func (c *comparer) CompareLine(baseLine string, targetLine string) bool {
	for _, ptn := range c.ignorePatterns {
		baseLine = ptn.ReplaceAllString(baseLine, "")
		targetLine = ptn.ReplaceAllString(targetLine, "")
	}
	return baseLine == targetLine
}

func (c *comparer) readAllLines(l *list.List, r io.Reader) {
	s := bufio.NewScanner(r)
	for isEOF := !s.Scan(); !isEOF; isEOF = !s.Scan() {
		l.PushBack(s.Text())
	}
}
