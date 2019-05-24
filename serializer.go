package main

// MessageSendByWebsocket : define our message Send by websocket
type MessageSendByWebsocket struct {
	Type    string                  `json:"type"`
	Name    string                  `json:"name"`
	Content DefineRegisterSubscribe `json:"content"`
}

// DefineRegisterSubscribe : define our registre to subscribe with register (is first octet to subscribe) and size (is number octet to subscribe)
type DefineRegisterSubscribe struct {
	Register int `json:"register"`
	Size     int `json:"size"`
}

// ContentResponse : define our Content response
type ContentResponse struct {
	RegisterSubscribe DefineRegisterSubscribe `json:"register"`
	Value             string                  `json:"value"`
}

// MessageReceiveByWebsocket : defines our message to sent by websocket
type MessageReceiveByWebsocket struct {
	Type    string          `json:"type"`
	Name    string          `json:"name"`
	Content ContentResponse `json:"content"`
}

type MyTestMSG struct {
	User    string          `json:"user`
	UserID	string          `json:"user_id`
}

type MyTestResponse struct {
	Data    string          `json:"data`
}
