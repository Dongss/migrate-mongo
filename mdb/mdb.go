package mdb

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
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
	Migrate(cln ClnOpt)
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
	log.Printf("Collection details: %v", cln.ClnNames)
	clns := m.collections()
	// check if collections exist
	for _, n := range cln.ClnNames {
		if con := contains(clns, n); !con {
			log.Fatalf("Collection not found: %v", n)
		}
		info := clnDetail(m.srcDb, n)
		info.Print()
	}
	fmt.Println()
}

func (m mdb) Migrate(cln ClnOpt) {
	log.Println("Start migration:")
	fmt.Println()
	ctx := context.Background()
	for _, n := range cln.ClnNames {
		fmt.Printf("start: %s\n", n)
		start := time.Now()
		var count int64
		c := m.srcDb.Collection(n)
		cur, err := c.Find(ctx, bson.M{})
		defer cur.Close(ctx)
		if err != nil {
			log.Fatal("Get collection error: ", err.Error())
		}
		for cur.Next(ctx) {
			elem := &bson.D{}
			if err := cur.Decode(elem); err != nil {
				log.Fatal(err)
			}
			dCln := m.dstDb.Collection(n)
			// currently, one by one
			// next, batch
			_, err := dCln.InsertOne(ctx, elem)
			if err != nil {
				log.Fatal("Insert data error: ", err.Error())
			}
			count++

		}
		t := time.Now()
		elapsed := t.Sub(start)
		fmt.Printf("done: %s, count: %d, elapsed: %v\n", n, count, elapsed)
		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}
	}

}

// get all source db collections
func (m mdb) collections() []string {
	ctx := context.Background()
	cur, err := m.srcDb.ListCollections(ctx, bson.M{})
	defer cur.Close(ctx)
	if err != nil {
		log.Fatal("Get collections error: ", err.Error())
	}
	var r []string
	for cur.Next(ctx) {
		var elem struct {
			Name string `bson:"name"`
		}
		if err := cur.Decode(&elem); err != nil {
			log.Fatal(err)
		}
		r = append(r, elem.Name)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return r
}

// get collection detail
func clnDetail(db *mongo.Database, clnName string) *clnInfo {
	ctx := context.Background()
	cln := db.Collection(clnName)
	c, err := cln.CountDocuments(ctx, bson.D{})
	if err != nil {
		log.Fatal("Get collection indexes error: ", err.Error())
	}

	indexes := cln.Indexes()
	cur, err := indexes.List(ctx)
	defer cur.Close(ctx)
	if err != nil {
		log.Fatal("Get collections error: ", err.Error())
	}

	var index []string

	for cur.Next(ctx) {
		var elem struct {
			Key  interface{} `bson:"key"`
			Name string      `bson:"name"`
			Ns   string      `bson:"ns"`
		}
		if err := cur.Decode(&elem); err != nil {
			log.Fatal("decode error: ", err.Error())
		} else {
			index = append(index, elem.Name)
		}
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	return &clnInfo{
		Name:    clnName,
		Count:   c,
		Indexes: index,
	}
}

func conDatabase(uri string) (*mongo.Client, *mongo.Database) {
	client, err := mongo.NewClient(MongoOpt.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Connect mongo error, ", err.Error())
	}
	tmp := strings.Split(uri, "/")
	dbstr := tmp[len(tmp)-1]
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Connect mongo error, ", err.Error())
	}
	database := client.Database(dbstr)
	return client, database
}
