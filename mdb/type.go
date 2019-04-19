package mdb

import (
	"fmt"
	"strings"
)

// ClnOpt is collection options
type ClnOpt struct {
	IfAll    bool     // if all collections
	IfIndex  bool     // if include indexes
	ClnNames []string // collection names
}

// MigOpt options for migration
type MigOpt struct {
	Interval int64
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
