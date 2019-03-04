package connectAPI

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"../database"
	"../modules"
	"github.com/adlio/trello"
	"github.com/sacOO7/gowebsocket"
)

func UpdateDataOnDB(key string, token string, id_board string) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	socket := gowebsocket.New("ws://echo.websocket.org/")
	socket.ConnectionOptions = gowebsocket.ConnectionOptions{
		UseSSL:         false,
		UseCompression: false,
		Subprotocols:   []string{"chat", "superchat"},
	}

	socket.RequestHeader.Set("Accept-Encoding", "gzip, deflate, sdch")
	socket.RequestHeader.Set("Accept-Language", "en-US,en;q=0.8")
	socket.RequestHeader.Set("Pragma", "no-cache")
	socket.RequestHeader.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.87 Safari/537.36")

	socket.OnConnectError = func(err error, socket gowebsocket.Socket) {
		log.Fatal("Recieved connect error ", err)
	}
	socket.OnConnected = func(socket gowebsocket.Socket) {
		log.Println("Connected to server")
		socket.SendText("data")
	}

	socket.OnTextMessage = func(message string, socket gowebsocket.Socket) {
		log.Println("Recieved message server " + message)
		if message == "data" {
			socket.SendText("Give data for me")
		} else {
			defer socket.SendText("data")
			GetCards(key, token, id_board, func(cards []*trello.Card, err error) {
				if err != nil {
					// notify to admin
				} else {
					myCards := ConverseFromCardToMyCard(key, token, cards)

					for _, v := range myCards {
						// fmt.Println(v)
						result, err := database.FindOne(v.ID)
						if err != nil {
							database.InsertData(v, func(err error) {
								if err != nil {
									fmt.Println("Can't insert")
								}
								fmt.Println("Inserted !")
							})
						} else {
							card := CompareCards(result, v)
							err := database.UpdateCard(card.ID, card)
							if err != nil {
								fmt.Println("Can't Update")
							} else {
								fmt.Println("Updated !")
							}
						}
					}
				}
			})
		}
	}

	socket.OnPingReceived = func(data string, socket gowebsocket.Socket) {
		log.Println("Recieved ping " + data)
	}
	socket.OnDisconnected = func(err error, socket gowebsocket.Socket) {
		log.Println("Disconnected from server ")
		return
	}

	socket.Connect()
	for {
		select {
		case <-interrupt:
			log.Println("interrupt")
			socket.Close()
			return

		}
	}
}

//@ Get all cardon board from trello api
func GetCards(key, token, id string, fn func([]*trello.Card, error)) {
	client := trello.NewClient(key, token)
	board, err := client.GetBoard(id, trello.Defaults())
	if err != nil {
		fn(nil, err)
	}
	cards, err := board.GetCards(trello.Defaults())
	if err != nil {
		fn(nil, err)
	}
	fn(cards, nil)
}

func GetListbById(appKey string, token string, ID string) (*trello.List, error) {
	client := trello.NewClient(appKey, token)
	list, err := client.GetList(ID, trello.Defaults())
	if err != nil {
		return nil, err
	}
	return list, nil
}

//@ Conver cards from trello api to my card
func ConverseFromCardToMyCard(key, token string, card []*trello.Card) (myCards []modules.MyCard) {
	var myCard modules.MyCard

	for i := 0; i < len(card); i++ {
		list, err := GetListbById(key, token, card[i].IDList)
		if err != nil {

		}
		myCards = append(myCards, myCard.New(card[i], list.Name))
	}
	return
}

//@ Compare two card and return new card
func CompareCards(cardOndb, cardOnTrello modules.MyCard) modules.MyCard {
	if modules.CompareTwoTime(cardOndb.DateLastActivity, cardOnTrello.DateLastActivity) == false {
		cardOndb.DateLastActivity = cardOnTrello.DateLastActivity
	}
	if modules.CompareTwoTime(cardOndb.Due, cardOnTrello.Due) == false {
		cardOndb.ChangeDueDate = true
		cardOndb.Due = cardOnTrello.Due
		cardOndb.HistoryChangeDueDate = modules.HandelHistory(cardOndb.HistoryChangeDueDate, cardOnTrello.Due)
	}
	return cardOndb
}
