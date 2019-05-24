// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import "fmt"



// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan *MyTestMSG

	// Register requests from the clients.
	register chan map[string]*Client

	// Unregister requests from clients.
	unregister chan *Client

	//user id
	usersID map[string]*Client


	userBroadcast chan *postUser

	userRegister chan *Server

	servers map[*Server]bool

}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan *MyTestMSG),
		register:   make(chan map[string]*Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		usersID:    make(map[string]*Client),

		userBroadcast:	make(chan *postUser),
		userRegister:	make(chan *Server),
		servers:	make(map[*Server]bool),
	}
}

func (h *Hub) run() {
	var clientUser map[string]*Client
	for {
		select {
		case serverUser := <-h.userRegister:
			fmt.Println("3")
			h.servers[serverUser] = true

		case message := <-h.userBroadcast:
			for server := range h.servers {
				select {
				case server.send <- message:
				default:
					close(server.send)
					delete(h.servers, server)
				}
			}

		case clientUser = <-h.register:
			for k,v := range clientUser {
				h.clients[v] = true
				h.usersID[k] = v
				fmt.Println("client:",k,"-",v)
			}

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				fmt.Println("删除：",h.usersID)
				for k,v := range h.usersID{
					if v == client{
						delete(h.usersID, k)

					}
				}
				fmt.Println("删除：",h.usersID)
				close(client.send)

			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}