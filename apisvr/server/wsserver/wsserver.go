package wsserver

import (
	"time"
	"net/http"
	"log"
	"bytes"
	"encoding/json"	

	"github.com/gorilla/websocket"

	"github.com/juancolamendy/water-jug-riddle/lib-service/service/wjsimulatorsvc"
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
	simulatorSvc *wjsimulatorsvc.SimulatorSvc
}

func NewWsServer() *WsServer {	
	return &WsServer{
		compCh: make(chan bool),
		simulatorSvc: wjsimulatorsvc.NewSimulatorSvc(wjsimulatorsvc.DefaultOpts()),
	}
}

func (ws *WsServer) Serve(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Websocket error: %v\n", err)
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
				log.Println("End of websocket connection")
				return
			case resp := <- ws.simulatorSvc.GetOutChan():
				// Encode simulator response
				b := bytes.NewBuffer(make([]byte, 0, 20))
				err := json.NewEncoder(b).Encode(resp)
				if err != nil {
					log.Printf("Error on encoding resp. Error:[%+v]\n", err)
					continue
				}
				msg := b.Bytes()

				// Forward simulator resp through the ws connection
				ws.conn.SetWriteDeadline(time.Now().Add(writeWait))
				if err := ws.conn.WriteMessage(websocket.TextMessage, msg); err != nil {
					log.Printf("Websocket error: %v\n", err)
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
		// Wait for input
		_, payload, err := ws.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Websocket error: %v\n", err)
			}
			log.Println("Sending end of websocket connection")
			ws.compCh <- true
			close(ws.compCh)
			return
		}
		// Decode input
		req := &wjsimulatorsvc.SimulateReq{}
		b := bytes.NewBuffer(payload)
		err = json.NewDecoder(b).Decode(req)
		if err != nil {
			log.Printf("Error on decoding payload. Payload:[%s] Error:[%+v]\n", string(payload), err)
			continue
		}

		// Send request to simulator
		ws.simulatorSvc.Simulate(req)
	}
}