package checks

import (
	"iter"
	"time"
)

type Check struct {
	Name     string `hcl:"name,label"`
	Target   string `hcl:"target"`
	PushURL  string
	Interval time.Duration
	Timeout  time.Duration
}

type Checks struct {
	items []Check
}

func (cs *Checks) All() iter.Seq[Check] {
	return func(yield func(Check) bool) {
		for _, c := range cs.items {
			if !yield(c) {
				return
			}
		}
	}
}
