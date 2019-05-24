package main

import (
	"fmt"
	"net/http"
	"sync"
)

type Server struct {
	hub *Hub

	// Buffered channel of outbound messages.
	send chan *postUser

	mu      sync.Mutex
}

func (c *Server) writePump() {
	fmt.Println("2")

	for {
		select {
		case msg, _ := <-c.send:
			teamID := msg.TeamID
			fmt.Println("teamid:", teamID)
			usersID := msg.UserID
			fmt.Println(usersID)
			//var mutex sync.Mutex
			c.mu.Lock()
			defer c.mu.Unlock()
			for _, v := range usersID {
				_, ok := c.hub.usersID[v]
				if ok {
					fmt.Println("writePump:", v)
					reData := & MyTestResponse{
						Data: fmt.Sprintf("team_id:%s,user_id:%s",teamID,v),
					}
					//w, err := c.conn.NextWriter(websocket.TextMessage)
					//if err != nil {
					//	return
					//}
					go func(c *Client) {
						err := c.conn.WriteJSON(reData)
						if err != nil {
							fmt.Println("error writePump")
							return
						}
					}(c.hub.usersID[v])
					//mutex.Lock()
					//err := c.hub.usersID[v].conn.WriteJSON(reData)
					//mutex.Unlock()
					//defer func() {
					//	c.hub.usersID[v].conn.Close()
					//}()
					//if err != nil {
					//	fmt.Println("error writePump")
					//	return
					//}
				}
			}









		}
	}
}

func serveLocal(hub *Hub, w http.ResponseWriter, r *http.Request, userID *postUser) {
	fmt.Println("0")
	server := &Server{hub: hub,send: make(chan *postUser, 256)}
	server.hub.userRegister <- server
	server.hub.userBroadcast <- userID
	fmt.Println("1")
	go server.writePump()
}

