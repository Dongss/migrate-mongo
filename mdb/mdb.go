package mdb

import (
	"context"
	"log"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	MongoOpt "go.mongodb.org/mongo-driver/mongo/options"
)

// MDB offers mongodb operation support
type MDB interface {
	// Connect init db connections
	Connect()

	// Disconnect disconnect
	Disconnect()

	// Overview shows source mongo database collection info,
	// includes document count, size, indexes
	Overview(cln ClnOpt)

	// Migrate do mgrations and return result
	Migrate(cln ClnOpt, opt migOpt)
}

type mdb struct {
	srcURI, dstURI       string
	srcClient, dstClient *mongo.Client
	srcDb, dstDb         *mongo.Database
}

// NewMDB create a new dbs
func NewMDB(srcURI string, dstURI string) MDB {
	m := &mdb{}
	m.srcURI = srcURI
	m.dstURI = dstURI
	return m
}

func (m *mdb) Connect() {
	m.srcClient, m.srcDb = conDatabase(m.srcURI)
	m.dstClient, m.dstDb = conDatabase(m.dstURI)
}

func (m *mdb) Disconnect() {
	ctx := context.Background()
	_ = m.srcClient.Disconnect(ctx)
	_ = m.dstClient.Disconnect(ctx)
}

func (m mdb) Overview(cln ClnOpt) {
	log.Printf("Show collection details: %v", cln.ClnNames)
}

func (m mdb) Migrate(cln ClnOpt, opt migOpt) {
}

// get all source db collections
func (m mdb) collections() []string {
	return []string{}
}

// get collection detail
func clnDetail(clbName string) {}

func conDatabase(uri string) (*mongo.Client, *mongo.Database) {
	client, err := mongo.NewClient(MongoOpt.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("connect mongo error, ", err.Error())
	}
	tmp := strings.Split(uri, "/")
	dbstr := tmp[len(tmp)-1]
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("connect mongo error, ", err.Error())
	}
	database := client.Database(dbstr)
	return client, database
}
