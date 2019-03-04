package database

import (
	"../config"
	"gopkg.in/mgo.v2"
)

var Config = config.ReadConfig()

func NewDatabase() (interface{}, error) {
	if Config.Database.Name == "mongodb" {
		db, err := mongoDB(Config.Database.Username, Config.Database.Password)
		if err != nil {
			return nil, err
		} else {
			return db, nil
		}
	}
	return nil, nil
}

func mongoDB(username, password string) (*mgo.Session, error) {
	session, err := mgo.Dial("mongodb://roger:roger123@ds213255.mlab.com:13255/report-tools")
	if err != nil {
		return nil, err
	}
	session.SetMode(mgo.Monotonic, true)
	return session, nil
}
