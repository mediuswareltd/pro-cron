package cron

import (
	"fmt"
	"strconv"
	"strings"
)

// Field holds the set of matching integers for one cron schedule field.
type Field struct {
	values map[int]bool
}

func (f *Field) Match(v int) bool {
	return f.values[v]
}

func parseField(s string, min, max int) (*Field, error) {
	f := &Field{values: make(map[int]bool)}
	for _, part := range strings.Split(s, ",") {
		if err := parsePart(part, min, max, f.values); err != nil {
			return nil, err
		}
	}
	return f, nil
}

func parsePart(s string, min, max int, values map[int]bool) error {
	// Step syntax: */n  or  a-b/n  or  a/n
	if slash := strings.Index(s, "/"); slash != -1 {
		step, err := strconv.Atoi(s[slash+1:])
		if err != nil || step <= 0 {
			return fmt.Errorf("invalid step in %q", s)
		}
		start, end := min, max
		prefix := s[:slash]
		if prefix != "*" {
			if dash := strings.Index(prefix, "-"); dash != -1 {
				start, err = strconv.Atoi(prefix[:dash])
				if err != nil {
					return fmt.Errorf("invalid range start in %q", s)
				}
				end, err = strconv.Atoi(prefix[dash+1:])
				if err != nil {
					return fmt.Errorf("invalid range end in %q", s)
				}
			} else {
				start, err = strconv.Atoi(prefix)
				if err != nil {
					return fmt.Errorf("invalid value in %q", s)
				}
				end = max
			}
		}
		if start < min || end > max || start > end {
			return fmt.Errorf("range %d-%d out of bounds [%d,%d] in %q", start, end, min, max, s)
		}
		for i := start; i <= end; i += step {
			values[i] = true
		}
		return nil
	}

	// Wildcard: *
	if s == "*" {
		for i := min; i <= max; i++ {
			values[i] = true
		}
		return nil
	}

	// Range: a-b
	if dash := strings.Index(s, "-"); dash != -1 {
		start, err := strconv.Atoi(s[:dash])
		if err != nil {
			return fmt.Errorf("invalid range start in %q", s)
		}
		end, err := strconv.Atoi(s[dash+1:])
		if err != nil {
			return fmt.Errorf("invalid range end in %q", s)
		}
		if start < min || end > max || start > end {
			return fmt.Errorf("range %d-%d out of bounds [%d,%d] in %q", start, end, min, max, s)
		}
		for i := start; i <= end; i++ {
			values[i] = true
		}
		return nil
	}

	// Single value
	n, err := strconv.Atoi(s)
	if err != nil {
		return fmt.Errorf("invalid value %q", s)
	}
	if n < min || n > max {
		return fmt.Errorf("value %d out of bounds [%d,%d]", n, min, max)
	}
	values[n] = true
	return nil
}
