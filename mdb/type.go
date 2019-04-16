package mdb

import (
	"fmt"
	"strings"
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

type clnInfo struct {
	Name    string
	Count   int64
	Indexes []string
}

func (c clnInfo) Print() {
	fmt.Println()
	// var s string
	fmt.Printf("Name: %s\nCount: %d\nIndexes:\n", c.Name, c.Count)
	fmt.Println("  ", strings.Join(c.Indexes, ", "))
}
