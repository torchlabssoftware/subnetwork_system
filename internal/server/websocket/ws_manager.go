package server

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	functions "github.com/torchlabssoftware/subnetwork_system/internal/server/functions"
)

var (
	websocketUpgrader = &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type WebsocketManager struct {
	Workers WorkerList
	sync.RWMutex
	Handlers map[string]EventHandler
}

func NewWebsocketManager() *WebsocketManager {
	w := &WebsocketManager{
		Workers:  make(WorkerList),
		Handlers: make(map[string]EventHandler),
	}
	w.setupEventHandlers()
	return w
}

func (ws *WebsocketManager) setupEventHandlers() {

}

func (ws *WebsocketManager) RouteEvent(event Event, w *Worker) error {
	if handler, ok := ws.Handlers[event.Type]; ok {
		if err := handler(event, w); err != nil {
			return err
		}
		return nil
	} else {
		return fmt.Errorf("no such event type")
	}
}

func (ws *WebsocketManager) ServeWS(w http.ResponseWriter, r *http.Request) {
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		functions.RespondwithError(w, http.StatusBadRequest, "Could not open websocket connection", err)
		return
	}
	defer conn.Close()

	log.Println("Worker connected via WebSocket.")

	worker := NewWorker(conn, ws)
	ws.AddWorker(worker)
	go worker.ReadMessage()
	go worker.WriteMessage()
}

func (ws *WebsocketManager) AddWorker(w *Worker) {
	ws.Lock()
	defer ws.Unlock()
	ws.Workers[w] = true
}

func (ws *WebsocketManager) RemoveWorker(w *Worker) {
	ws.Lock()
	defer ws.Unlock()
	if _, ok := ws.Workers[w]; ok {
		w.Connection.Close()
		delete(ws.Workers, w)
	}
}
