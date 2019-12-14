package mongodb

import (
	"context"
	"fmt"
	"time"

	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

func MongoConnect(host, database string) (*mongo.Database, error) {
	uri := fmt.Sprintf("mongodb://%s:27017", host)
	opt := options.Client().ApplyURI(uri)
	client, err := mongo.NewClient(opt)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}
	db := client.Database(database)
	m := migrate.NewMigrate(db, migrate.Migration{
		Version:     1,
		Description: "add my-index",
		Up: func(db *mongo.Database) error {
			opt := options.Index().SetName("my-index")
			keys := bson.D{{"my-key", 1}}
			model := mongo.IndexModel{Keys: keys, Options: opt}
			_, err := db.Collection("my-coll").Indexes().CreateOne(context.TODO(), model)
			if err != nil {
				return err
			}

			return nil
		},
		Down: func(db *mongo.Database) error {
			_, err := db.Collection("my-coll").Indexes().DropOne(context.TODO(), "my-index")
			if err != nil {
				return err
			}
			return nil
		},
	})
	if err := m.Up(migrate.AllAvailable); err != nil {
		return nil, err
	}
	return db, nil
}