package worker

import (
	"github.com/google/uuid"
)

type Event struct {
	Type    string      `json:"type"`
	Payload interface{} `json:"payload"`
}

type EventHandler func(event Event) error

type ErrorPayload struct {
	Success bool        `json:"success"`
	Payload interface{} `json:"payload"`
}

type LoginRequest struct {
	WorkerID string `json:"worker_id"`
}

type LoginResponse struct {
	Otp string `json:"otp"`
}

type ConfigPayload struct {
	PoolID        uuid.UUID        `json:"pool_id"`
	PoolTag       string           `json:"pool_tag"`
	PoolPort      int              `json:"pool_port"`
	PoolSubdomain string           `json:"pool_subdomain"`
	Upstreams     []UpstreamConfig `json:"upstreams"`
}

type UpstreamConfig struct {
	UpstreamID      uuid.UUID `json:"upstream_id"`
	UpstreamTag     string    `json:"upstream_tag"`
	UpstreamAddress string    `json:"upstream_address"`
	UpstreamHost    string    `json:"upstream_host"`
	UpstreamPort    int       `json:"upstream_port"`
	Weight          int       `json:"weight"`
}

type User struct {
	ID          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Status      string    `json:"status"`
	IpWhitelist []string  `json:"ip_whitelist"`
	Pools       []string  `json:"pools"`
}
