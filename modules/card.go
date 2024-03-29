package modules

import (
	"time"

	"github.com/adlio/trello"
)

type MyCard struct {
	ID                   string
	Name                 string
	ListName             string
	IdList               string
	TimeGuessForDone     int
	TimeRealForDone      int
	DateLastActivity     *time.Time
	Due                  *time.Time
	ChangeDueDate        bool
	HistoryChangeDueDate []*time.Time
}

func (mc MyCard) New(card *trello.Card, listName string) (myCard MyCard) {
	myCard.ID = card.ID
	myCard.Name = card.Name
	myCard.TimeGuessForDone = GetTimeGuessForDone(card.Name)
	myCard.TimeRealForDone = GetRealTimeOfDone(card.Name)
	myCard.DateLastActivity = card.DateLastActivity
	myCard.ListName = listName
	myCard.IdList = card.IDList
	myCard.Due = card.Due
	myCard.ChangeDueDate = false
	myCard.HistoryChangeDueDate = HandelHistory(myCard.HistoryChangeDueDate, card.Due)
	return
}

//@ Compare two card and return new card
func (c MyCard) CompareTwoCards(cardOnTrello MyCard) MyCard {
	if CompareTwoTime(c.DateLastActivity, cardOnTrello.DateLastActivity) == false {
		c.DateLastActivity = cardOnTrello.DateLastActivity
	}
	if CompareTwoTime(c.Due, cardOnTrello.Due) == false {
		c.ChangeDueDate = true
		c.Due = cardOnTrello.Due
		c.HistoryChangeDueDate = HandelHistory(c.HistoryChangeDueDate, cardOnTrello.Due)
	}
	return c
}
