package mongoconnector

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoConnection struct {
	Configuration Config
}

var mdb MongoConnection

func init() {
	c := SetConfig()
	mdb.Configuration = c
}

// Ping tries to hit the Mongo instance.
func Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Printf("Trying to connect to MongoDB %s...\n", mdb.Configuration.Collection)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mdb.Configuration.GetConnection()))
	if err != nil {
		return err
	}
	defer client.Disconnect(context.Background())

	coll := client.Database(mdb.Configuration.Database).Collection(mdb.Configuration.Collection)
	ctx2, cancel2 := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel2()

	int, err := coll.EstimatedDocumentCount(ctx2)
	if err != nil {
		return err
	}
	log.Printf("Mongo found %d docs\n", int)
	return nil
}

type User struct {
	Username string
	ID       *primitive.ObjectID `json:"ID" bson:"_id,omitempty"`
}

// IsBanned checks if user mail is banned.
func IsBanned(username string) bool {
	client, err := mongo.NewClient(options.Client().ApplyURI(mdb.Configuration.GetConnection()))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer cancel()
	defer client.Disconnect(ctx)

	user := User{}

	coll := client.Database(mdb.Configuration.Database).Collection(mdb.Configuration.Collection)
	err = coll.FindOne(context.TODO(), bson.D{{"username", username}}).Decode(&user)

	return err != mongo.ErrNoDocuments
}
