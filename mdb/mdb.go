package mdb

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// MDB offers mongodb operation support
type MDB interface {
	// Overview shows source mongo database collection info,
	// includes document count, size, indexes
	Overview(cln clnOpt) (migGoal, error)

	// Migrate do mgrations and return result
	Migrate(cln clnOpt, opt migOpt) (migResult, error)
}

type mdb struct {
	srcClient, dstClient *mongo.Client
	srcDb, dstDb         *mongo.Database
}

// NewMDB create a new dbs
func NewMDB(srcURI string, dstURI string) MDB {
	return nil
}

func (m mdb) Overview(cln clnOpt) (migGoal, error) {
	return migGoal{}, nil
}

func (m mdb) Migrate(cln clnOpt, opt migOpt) (migResult, error) {
	return migResult{}, nil
}
