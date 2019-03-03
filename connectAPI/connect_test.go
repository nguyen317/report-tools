package connectAPI

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"

	"net/http"

	"github.com/sacOO7/gowebsocket"
)

func CheckUpdateDataOnDB(key string, token string, id_board string) {
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
			response, err := http.Get("https://api.trello.com/1/boards/" + id_board + "/cards?key=" + key + "&token=" + token)
			if err != nil {
				fmt.Printf("The HTTP request failed with error %s\n", err)
			} else {
				data, _ := ioutil.ReadAll(response.Body)
				fmt.Println(string(data))
				fmt.Printf("%T", data)
			}
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
	// socket.SendText("data")
	// socket.Close()
	for {
		select {
		case <-interrupt:
			log.Println("interrupt")
			socket.Close()
			return

		}
	}
}
