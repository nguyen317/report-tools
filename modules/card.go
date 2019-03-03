package modules

import (
	"time"

	"github.com/adlio/trello"
)

type MyCard struct {
	ID                   string
	Name                 string
	TimeGuessForDone     int
	TimeRealForDone      int
	DateLastActivity     *time.Time
	Due                  *time.Time
	ChangeDueDate        bool
	HistoryChangeDueDate []*time.Time
}

func (mc MyCard) New(card *trello.Card) (myCard MyCard) {
	myCard.ID = card.ID
	myCard.Name = card.Name
	myCard.TimeGuessForDone = GetTimeGuessForDone(card.Name)
	myCard.TimeRealForDone = GetRealTimeOfDone(card.Name)
	myCard.DateLastActivity = card.DateLastActivity
	myCard.Due = card.Due
	myCard.HistoryChangeDueDate = nil
	return
}
