package server

import "github.com/google/uuid"

type Event struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type EventHandler func(event Event, w *Worker) error

type loginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type successPayload struct {
	Success bool        `json:"success"`
	Payload interface{} `json:"payload"`
}

type errorPayload struct {
	Success bool        `json:"success"`
	Payload interface{} `json:"payload"`
}

type UpstreamConfig struct {
	UpstreamID      uuid.UUID `json:"upstream_id"`
	UpstreamTag     string    `json:"upstream_tag"`
	UpstreamAddress string    `json:"upstream_address"`
	UpstreamHost    string    `json:"upstream_host"`
	UpstreamPort    int32     `json:"upstream_port"`
	Weight          int32     `json:"weight"`
}

type ConfigPayload struct {
	PoolID        uuid.UUID        `json:"pool_id"`
	PoolTag       string           `json:"pool_tag"`
	PoolPort      int32            `json:"pool_port"`
	PoolSubdomain string           `json:"pool_subdomain"`
	Upstreams     []UpstreamConfig `json:"upstreams"`
}
