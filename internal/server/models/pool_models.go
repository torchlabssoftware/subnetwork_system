package server

import (
	"time"

	"github.com/google/uuid"
)

type GetRegionResponce struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at "`
}

type CreateRegionRequest struct {
	Name string `json:"name"`
}

type CreateRegionResponce struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at "`
}

type DeleteRegionRequest struct {
	Name string `json:"name"`
}
