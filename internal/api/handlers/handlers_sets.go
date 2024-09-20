package handlers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
	"github.com/gin-gonic/gin"
)

func HandleWebSocker(c *gin.Context) {
	role := c.DefaultQuery("role", "watch")

	conn, err := websocket.Accept(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println("Failed to accept WebSocket connection:", err)
		return
	}
	defer conn.Close(websocket.StatusNormalClosure, "")

	if role == "interact" {
		handleInteractiveConnection(conn)
	} else {
		handleWatchingConnection(conn)
	}
}

func handleWatchingConnection(conn *websocket.Conn) {
	ctx := context.Background()

	for {
		time.Sleep(2 * time.Second)
		err := wsjson.Write(ctx, conn, "match updated")
		if err != nil {
			fmt.Println("Error sending update to watcher:", err)
			return
		}
	}
}

func handleInteractiveConnection(conn *websocket.Conn) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	for {
		var message interface{}
		err := wsjson.Read(ctx, conn, &message)
		if err != nil {
			log.Println("Read error or timeout:", err)
			return
		}

		log.Println("Received interactive message:", message)
	}
}
