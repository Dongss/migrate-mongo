package mdb

import (
	"time"
)

// ClnOpt is collection options
type ClnOpt struct {
	IfAll    bool     // if all collections
	IfIndex  bool     // if include indexes
	ClnNames []string // collection names
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
