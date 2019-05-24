package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type postUser struct {
	// 拼团ID
	// required: true
	TeamID  string `json:"team_id"`
	// 用户ID
	// required: true
	UserID  []string `json:"user_id"`
}


func main() {
	router := gin.Default()
	hub := newHub()
	go hub.run()

	router.GET("/websocket/:userID", func(c *gin.Context) {
		//func(w http.ResponseWriter, r *http.Request)
		userID := c.Param("userID")
		serveWs(hub, c.Writer, c.Request, userID)
	})
	router.POST("/websocket", func(c *gin.Context) {
		b := &postUser{}



		if err := c.Bind(b);err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "websocket服务端出错"})
			return
		}

		fmt.Println(b)



		serveLocal(hub, c.Writer, c.Request, b)

	})
	router.Run(":5000")
}
