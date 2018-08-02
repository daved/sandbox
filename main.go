package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

// Person ...
type Person struct {
	Name string
	Age  int
}

type node struct {
	ug *websocket.Upgrader
}

func (n *node) simpleHandler(w http.ResponseWriter, r *http.Request) {
	c, err := n.ug.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err = c.Close(); err != nil {
			fmt.Println(err)
		}

		fmt.Println("client unsubbed")
	}()

	fmt.Println("client subbed")

	p := Person{
		Name: "Bill",
		Age:  0,
	}

	for i := 0; i <= 20; i++ {
		j, err := json.Marshal(p)
		if err != nil {
			fmt.Println(err)
			return
		}

		if err = c.WriteMessage(websocket.TextMessage, j); err != nil {
			fmt.Println(err)
			break
		}

		p.Age += 2
	}
}

func (n *node) pingpongHandler(w http.ResponseWriter, r *http.Request) {
	c, err := n.ug.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer func() {
		if err = c.Close(); err != nil {
			fmt.Println(err)
		}

		fmt.Println("client unsubbed")
	}()

	fmt.Println("client subbed")

	for {
		msgType, msg, err := c.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}

		if string(msg) == "ping" {
			fmt.Println("ping")
			time.Sleep(time.Second * 2)
			if err = c.WriteMessage(msgType, []byte("pong")); err != nil {
				fmt.Println(err)
				return
			}

			continue
		}
	}
}

func main() {
	n := &node{
		ug: &websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}

	http.HandleFunc("/websocket/simple", n.simpleHandler)
	http.HandleFunc("/websocket/pingpong", n.pingpongHandler)

	if err := http.ListenAndServe(":3000", nil); err != nil {
		fmt.Println(err)
	}
}
