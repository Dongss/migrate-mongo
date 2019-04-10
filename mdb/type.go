package mdb

import (
	"time"
)

// collection options
type clnOpt struct {
	ifAll    bool     // if all collections
	ifIndex  bool     // if include indexes
	clnNames []string // collection names
}

// migrate otions
type migOpt struct{}

// migrate goal
type migGoal struct {
}

// migrate result
type migResult struct {
	totalTime time.Duration
}

func (m migGoal) String() string {
	return ""
}

func (m migResult) String() string {
	return ""
}
