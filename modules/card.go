package modules

import "time"

type Card struct {
	ID                   string
	Name                 string
	TimeGuessForDone     int
	TimeRealForDone      int
	Date                 *time.Time
	DateLastActivity     *time.Time
	ChangeDueDate        bool
	HistoryChangeDueDate []*time.Time
}

func New() {

}
