package database

import (
	"../config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Config = config.ReadConfig()

func NewDatabase(db interface{}) error {
	var err error
	if Config.Database.Name == "mongodb" {
		db, err = mongodb(Config.Database.Username, Config.Database.Password)
		if err != nil {
			return err
		} else {
			return nil
		}
	}
	return nil
}

func mongodb(username, password string) (*mongo.Client, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://" + username + ":" + password + "@ds213255.mlab.com:13255/report-tools"))

	if err != nil {
		return nil, err
	}

	return client, nil

}
