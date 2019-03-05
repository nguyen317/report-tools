package database

import (
	"../modules"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var collection *mgo.Collection

func init() {
	db, _ := NewDatabase()
	db.Session.SetMode(mgo.Monotonic, true)
	collection = db.Session.DB("report-tools").C("cards")
}

func InsertData(data interface{}, fn func(error)) {
	err := collection.Insert(data)
	if err != nil {
		fn(err)
	}
	fn(nil)
}

func FindOne(id string) (modules.MyCard, error) {
	var card modules.MyCard
	err := collection.Find(bson.M{"id": id}).One(&card)
	if err != nil {
		return card, err
	}
	return card, nil
}

func GetAllCard() ([]modules.MyCard, error) {
	var cards []modules.MyCard
	err := collection.Find(bson.M{}).All(&cards)
	if err != nil {
		return nil, err
	}
	return cards, nil
}

func UpdateCard(id string, data interface{}) error {
	err := collection.Update(bson.M{"id": id}, bson.M{"$set": data})
	if err != nil {
		return err
	}
	return nil
}
