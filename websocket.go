package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsUpgrader = websocket.Upgrader{}

func wsEcho(w http.ResponseWriter, r *http.Request) {
	c, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("failed to set websocket upgrade: %e\n", err)
		return
	}

	for {
		t, msg, err := c.ReadMessage()
		if err != nil {
			fmt.Printf("failed to read message: %e\n", err)
			break
		}
		fmt.Printf("message: %s\n", msg)
		c.WriteMessage(t, msg)
	}
}

var echoHandler gin.HandlerFunc = func(c *gin.Context) {
	wsEcho(c.Writer, c.Request)
}
