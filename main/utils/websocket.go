package utils

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type CustomWSResponse struct {
	User    string
	Message string
}

var connection []*websocket.Conn

func WS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	connection = append(connection, conn)
	fmt.Println("new connection on 8080")
	if err != nil {
		fmt.Println("error upgrade connection")
		return
	}

	pingperiode := 5 * time.Second * 9 / 10
	//pongperiode := 5 * time.Second

	conn.SetPongHandler(func(appData string) error {
		fmt.Println("pong arrived")
		return nil
	})

	go func() {
		ticker := time.NewTicker(pingperiode)
		defer func() {
			ticker.Stop()
		}()

		for {
			var err error
			select {
			case <-ticker.C:
				fmt.Println("ping to client")
				conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
				err = conn.WriteMessage(websocket.PingMessage, []byte("ping ping"))
				if err != nil {
					fmt.Println("error write message")
					break
				}
			}
			if err != nil {
				break
			}
		}
	}()

	go func() {
		defer func() {
			conn.Close()
		}()
		for {
			_, p, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("read message error")
				break
			}

			time.Sleep(1 * time.Second)
			fmt.Println(string(p))

			//for _, con := range connection {
			//	if con != conn {
			//		err = con.WriteMessage(websocket.TextMessage, []byte("tes"))
			//		if err != nil {
			//			fmt.Println("write message err")
			//			return
			//		}
			//	}
			//}

		}
	}()
}
