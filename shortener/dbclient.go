package shortener

import (
	"log"
	"time"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoClient holds master session and other db-related info
type MongoClient struct {
	session *mgo.Session // master session
	uri     string       // mongodb uri
	dbName  string       // database name
}

// NewMongoClient establishes connection to MongoDB database
// and returns new MongoClient object
func NewMongoClient(uri, dbName string) *MongoClient {
	session, err := mgo.Dial(uri)
	if err != nil {
		log.Panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	return &MongoClient{session, uri, dbName}
}

// GetSession returns mgo.Session copied from
// MongoClient's master session
// Be sure to close the session after done
func (mc *MongoClient) GetSession() *mgo.Session {
	return mc.session.Copy()
}

// Register inserts URL into database
func (mc *MongoClient) Register(original, short string) error {
	s := mc.GetSession()
	// TODO: create index??
	err := s.DB(mc.dbName).C("map").Insert(&URIMap{
		original,
		short,
		time.Now().UTC(),
	})
	s.Close()
	return err
}

// FindOriginal searches for Original URL from database
func (mc *MongoClient) FindOriginal(short string) (*URIMap, error) {
	s := mc.GetSession()
	defer s.Close()
	var uriMap URIMap
	err := s.DB(mc.dbName).C("map").Find(bson.M{"short": short}).One(&uriMap)
	if err != nil {
		return nil, err
	}
	return &uriMap, err
}
