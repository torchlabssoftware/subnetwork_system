package server

type Event struct {
	Type   string `json:"type"`
	Pyload string `json:"payload"`
}

type EventHandler func(event Event, w *Worker) error
