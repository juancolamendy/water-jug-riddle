package wsserver

import (
	"time"
	"net/http"
	"log"

	"github.com/gorilla/websocket"
)

const (
	writeWait = 10 * time.Second
	pongWait = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
	maxMessageSize = 512
	wsBufferSize = 1024
)

var upgrader = websocket.Upgrader {
	ReadBufferSize:  wsBufferSize,
	WriteBufferSize: wsBufferSize,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},	
}

type WsServer struct {
	conn *websocket.Conn
	compCh chan bool

	notificationCh chan string
}

func NewWsServer() *WsServer {
	return &WsServer{
		compCh: make(chan bool),

		notificationCh: make(chan string),
	}
}

func (ws *WsServer) Serve(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("--- Websocket error: %v\n", err)
		return
	}
	ws.conn = conn

	ws.write()
	ws.read()
}

func (ws *WsServer) write() {
	go func() {
		defer func() {
			if(ws.conn != nil) {
				ws.conn.Close()
				ws.conn = nil
			}
		}()

		for {
			select {
			case <- ws.compCh:
				log.Println("--- End of websocket connection")
				return
			case msg := <- ws.notificationCh:
				log.Printf("--- Received and forwarding message: %s\n", msg)
				ws.conn.SetWriteDeadline(time.Now().Add(writeWait))
				if err := ws.conn.WriteMessage(websocket.TextMessage, []byte(msg)); err != nil {
					log.Printf("--- Websocket error: %v\n", err)
					continue
				}
			}
		}
	}()	
}

func (ws *WsServer) read() {
	defer func() {
		if(ws.conn != nil) {
			ws.conn.Close()
			ws.conn = nil
		}
	}()

	ws.conn.SetReadLimit(maxMessageSize)
	ws.conn.SetReadDeadline(time.Now().Add(pongWait))
	ws.conn.SetPongHandler(func(string) error { ws.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	for {
		_, payload, err := ws.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("--- Websocket error: %v\n", err)
			}
			log.Println("--- Sending end of websocket connection")
			ws.compCh <- true
			close(ws.compCh)
			return
		}
		log.Printf("--- %+v\n", string(payload))
		ws.notificationCh <- string(payload)
	}
}