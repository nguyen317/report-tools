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
		//Proxy: gowebsocket.BuildProxy("http://example.com"),
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
			// defer socket.SendText("data")
			GetCards(key, token, id_board, func(cards []*trello.Card, err error) {
				if err != nil {
					// notify to admin
				} else {
					myCards := ConverseFromCardToMyCard(cards)
					for k, v := range myCards {
						database.FindOne(v.ID, func(result interface{}, err error) {
							fmt.Println(k)
							fmt.Println("res", result)
							fmt.Println("Err", err)
							// if result.ID != nil {

							// }
							if err != nil {
								// database.InsertData(myCards[i], func(res *mongo.InsertOneResult, err error) {
								// 	if err != nil {
								// 		fmt.Println("Can't Insert")
								// 	}
								// 	fmt.Println("Inserted")
								// })
							} else {
								CompareCards(result, v)
							}
						})

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

//@ Conver cards from trello api to my card
func ConverseFromCardToMyCard(card []*trello.Card) (myCards []modules.MyCard) {
	var myCard modules.MyCard
	for i := 0; i < len(card); i++ {
		myCards = append(myCards, myCard.New(card[i]))
	}
	return
}

func CompareCards(cardOndb, cardOnTrello interface{}) {
	fmt.Println(cardOndb)

	// fmt.Println(cardOnTrello)
}
