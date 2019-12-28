package dbmongo

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type DBMongo struct {
	Client *mongo.Client
}


func Connect(connUrl string, clientOptions *options.ClientOptions) *DBMongo{

	if clientOptions == nil {
		clientOptions = options.Client().
			SetMinPoolSize(3).
			SetMaxPoolSize(60).
			SetConnectTimeout(15 * time.Second).
			SetMaxConnIdleTime(5 * time.Minute).
			ApplyURI(connUrl)
	}

	db := &DBMongo{}

	var err error
	db.Client, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		logrus.Panicf("db issues: %v", err)
	}

	err = db.Client.Ping(context.Background(), nil)
	if err != nil {
		logrus.Panicf("db issues: %v", err)
	}
	return db
}

func (db *DBMongo) Disconnect(){
	err := db.Client.Disconnect(context.Background())
	if err != nil {
		logrus.Info("disconnect issues: %v", err)
	}
}


func (db *DBMongo) GetDatabase(dbName string) *mongo.Database{
	return db.Client.Database(dbName, options.Database())
}


func (db *DBMongo) GetCollection(dbName, colName string) *mongo.Collection {
	return db.GetDatabase(dbName).Collection(colName, options.Collection())
}
