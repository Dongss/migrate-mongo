package mdb

import (
	"fmt"
	"strings"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ClnOpt is collection options
type ClnOpt struct {
	IfAll    bool     // if all collections
	ClnNames []string // collection names
}

// MigOpt options for migration
type MigOpt struct {
	Interval int64
	IfIndex  bool // if include indexes
	FBatch   int32
}

type clnIndex struct {
	Key  primitive.D `bson:"key"`
	Name string      `bson:"name"`
	Ns   string      `bson:"ns"`
}

type clnInfo struct {
	Name    string
	Count   int64
	Indexes []clnIndex
}

func (c clnInfo) Print() {
	fmt.Println()
	fmt.Printf("Name: %s\nCount: %d\nIndexes:\n", c.Name, c.Count)
	var s []string
	for _, v := range c.Indexes {
		s = append(s, v.Name)
	}
	fmt.Println("  ", strings.Join(s, ", "))
}
