package ws

type Server struct {
	Pattern   string
	Messages  []*Message
	Clients   map[int]*Client
	AddCh     chan *Client
	DelCh     chan *Client
	SendAllCh chan *Message
}

type Message struct {
	Token string `json:"token"`
	Body  string `json:"body"`
}
