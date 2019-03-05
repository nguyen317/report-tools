package connectAPI

import (
	"fmt"
	"sync"

	"../database"
	"../modules"
	"github.com/adlio/trello"
)

var wg sync.WaitGroup

//@ Update data on database real time
func UpdateDataOnDB(key string, token string, id_board string) {
	wg.Add(1)
	chanCard := make(chan []*trello.Card, 3)
	go GiveData(chanCard, key, token, id_board)
	go HandelData(chanCard, key, token)
	wg.Wait()
}

// @ Write data on chanel card
func GiveData(chanCard chan []*trello.Card, key, token, id string) {
	for {
		cards, err := modules.GetCardsOnTrelloAPI(key, token, id)
		if err != nil {
			chanCard <- nil
		} else {
			chanCard <- cards
		}
	}
}

//@ Conver cards from trello api to my card
func ConverseFromCardToMyCard(key, token string, card []*trello.Card) (myCards []modules.MyCard) {
	var myCard modules.MyCard

	for i := 0; i < len(card); i++ {
		list, err := modules.GetListbByIdOnTrelloAPI(key, token, card[i].IDList)
		if err != nil {

		}
		myCards = append(myCards, myCard.New(card[i], list.Name))
	}
	return
}

//@ Handel data on chanel card
func HandelData(chanCard chan []*trello.Card, key, token string) {
	for {
		cards := <-chanCard
		if cards != nil {
			var myCard modules.MyCard
			myCards := modules.MapFromTrelloCardToMyCard(cards, func(item *trello.Card) modules.MyCard {
				return func() modules.MyCard {
					list, err := modules.GetListbByIdOnTrelloAPI(key, token, item.IDList)
					if err != nil {

					}
					return myCard.New(item, list.Name)
				}()
			})
			for _, v := range myCards {
				result, err := database.FindOne(v.ID)
				if err != nil {
					database.InsertData(v, func(err error) {
						if err != nil {
							fmt.Println("Can't insert")
						}
						fmt.Println("Inserted !")
					})
				} else {
					card := result.CompareTwoCards(v)
					err := database.UpdateCard(card.ID, card)
					if err != nil {
						fmt.Println("Can't Update")
					} else {
						fmt.Println("Updated !")
					}
				}
			}
		} else {
			fmt.Println("false")
		}
	}
}
