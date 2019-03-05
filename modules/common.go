package modules

import (
	"strconv"
	"time"

	"github.com/adlio/trello"
)

func init() {

}

func GetRealTimeOfDone(name string) int {
	l := len(name)
	time := ""
	for i := l - 1; i > 0; i-- {
		if string(name[i]) == "]" {
			i--
			for ; string(name[i]) != "["; i-- {
				time = string(name[i]) + string(time)
			}
			break
		}
	}
	ret, _ := strconv.Atoi(time)
	return ret
}

func GetTimeGuessForDone(name string) int {
	l := len(name)
	time := ""
	for i := l - 1; i > 0; i-- {
		if string(name[i]) == ")" {
			i--
			for ; string(name[i]) != "("; i-- {
				time = string(name[i]) + string(time)
			}
			break
		}
	}
	ret, _ := strconv.Atoi(time)
	return ret
}

func HandelHistory(data []*time.Time, due *time.Time) []*time.Time {
	if due != nil {
		return append(data, due)
	}
	return data
}

func CompareTwoTime(a, b *time.Time) bool {
	if a != nil && b != nil && a.Local().Format("2006-01-02") != b.Local().Format("2006-01-02") {
		return false
	}
	return true
}

//@ Filter []modules.MyCard
func Filter(vs []MyCard, f func(MyCard) bool) (vsf []MyCard) {
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return
}

//@ Map from *trello.Card to Mycard []modules.MyCard
func MapFromTrelloCardToMyCard(vs []*trello.Card, f func(*trello.Card) MyCard) []MyCard {
	vsm := make([]MyCard, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
