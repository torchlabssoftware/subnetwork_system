package server

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/torchlabssoftware/subnetwork_system/internal/db/repository"
	functions "github.com/torchlabssoftware/subnetwork_system/internal/server/functions"
	models "github.com/torchlabssoftware/subnetwork_system/internal/server/models"
)

type PoolHandler struct {
	Queries *repository.Queries
	DB      *sql.DB
}

func NewPoolHandler(queries *repository.Queries, db *sql.DB) *PoolHandler {
	return &PoolHandler{
		Queries: queries,
		DB:      db,
	}
}

func (p *PoolHandler) Routes() http.Handler {
	r := chi.NewRouter()
	r.Get("/region", p.getRegions)
	r.Post("/region", p.createRegion)
	r.Delete("/region", p.DeleteRegion)
	return r
}

func (p *PoolHandler) getRegions(w http.ResponseWriter, r *http.Request) {

	regions, err := p.Queries.GetRegions(r.Context())
	if err != nil {
		functions.RespondwithError(w, http.StatusBadRequest, "server error", err)
		return
	}

	res := make([]models.GetRegionResponce, len(regions))

	for _, region := range regions {
		r := models.GetRegionResponce{
			Id:        region.ID,
			Name:      region.Name,
			CreatedAt: region.CreatedAt,
			UpdatedAt: region.UpdatedAt,
		}

		res = append(res, r)
	}

	functions.RespondwithJSON(w, http.StatusCreated, res)
}

func (p *PoolHandler) createRegion(w http.ResponseWriter, r *http.Request) {
	var req models.CreateRegionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		functions.RespondwithError(w, http.StatusBadRequest, "err in request body", err)
		return
	}

	region, err := p.Queries.AddRegion(r.Context(), req.Name)
	if err != nil {
		functions.RespondwithError(w, http.StatusBadRequest, "server error", err)
		return
	}

	res := models.CreateRegionResponce{
		Id:        region.ID,
		Name:      region.Name,
		CreatedAt: region.CreatedAt,
		UpdatedAt: region.UpdatedAt,
	}

	functions.RespondwithJSON(w, http.StatusCreated, res)
}

func (p *PoolHandler) DeleteRegion(w http.ResponseWriter, r *http.Request) {
	var req models.DeleteRegionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		functions.RespondwithError(w, http.StatusBadRequest, "err in request body", err)
		return
	}

	err := p.Queries.DeleteRegion(r.Context(), req.Name)
	if err != nil {
		functions.RespondwithError(w, http.StatusBadRequest, "server error", err)
		return
	}

	res := struct {
		Message string `json:"message"`
	}{
		Message: "deleted",
	}

	functions.RespondwithJSON(w, http.StatusOK, res)
}
