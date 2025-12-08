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
	Name *string `json:"name"`
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

type GetCountryResponce struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name,omitempty"`
	Code      string    `json:"code,omitempty"`
	RegionId  uuid.UUID `json:"region_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at "`
}

type CreateCountryRequest struct {
	Name     *string    `json:"name"`
	Code     *string    `json:"code"`
	RegionId *uuid.UUID `json:"region_id"`
}

type CreateCountryResponce struct {
	Id        uuid.UUID `json:"id"`
	Name      string    `json:"name,omitempty"`
	Code      string    `json:"code,omitempty"`
	RegionId  uuid.UUID `json:"region_id,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at "`
}

type DeleteCountryRequest struct {
	Name string `json:"name"`
}
