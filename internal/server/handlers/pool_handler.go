package server

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/torchlabssoftware/subnetwork_system/internal/db/repository"
	functions "github.com/torchlabssoftware/subnetwork_system/internal/server/functions"
	middleware "github.com/torchlabssoftware/subnetwork_system/internal/server/middleware"
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

	r.Use(middleware.AdminAuthentication)

	r.Get("/region", p.getRegions)
	r.Post("/region", p.createRegion)
	r.Delete("/region", p.DeleteRegion)
	r.Get("/country", p.getcountries)
	r.Post("/country", p.createCountry)
	r.Delete("/country", p.DeleteCountry)
	r.Get("/upstream", p.getUpstreams)
	r.Post("/upstream", p.createUpstream)
	r.Delete("/upstream", p.deleteUpstream)

	r.Post("/", p.createPool)
	r.Get("/", p.getPools)
	return r
}

func (p *PoolHandler) getRegions(w http.ResponseWriter, r *http.Request) {

	regions, err := p.Queries.GetRegions(r.Context())
	if err != nil {
		functions.RespondwithError(w, http.StatusBadRequest, "server error", err)
		return
	}

	res := []models.GetRegionResponce{}

	for _, region := range regions {
		r := models.GetRegionResponce{
			Id:        region.ID,
			Name:      region.Name,
			CreatedAt: region.CreatedAt,
			UpdatedAt: region.UpdatedAt,
		}

		res = append(res, r)
	}

	functions.RespondwithJSON(w, http.StatusOK, res)
}

func (p *PoolHandler) createRegion(w http.ResponseWriter, r *http.Request) {
	var req models.CreateRegionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		functions.RespondwithError(w, http.StatusBadRequest, "err in request body", err)
		return
	}

	if req.Name == nil && *req.Name == "" {
		functions.RespondwithError(w, http.StatusBadRequest, "add Request name", fmt.Errorf("no region name"))
		return
	}

	region, err := p.Queries.AddRegion(r.Context(), *req.Name)
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

func (p *PoolHandler) getcountries(w http.ResponseWriter, r *http.Request) {

	countries, err := p.Queries.GetCountries(r.Context())
	if err != nil {
		functions.RespondwithError(w, http.StatusBadRequest, "server error", err)
		return
	}

	res := []models.GetCountryResponce{}

	for _, country := range countries {
		r := models.GetCountryResponce{
			Id:        country.ID,
			Name:      country.Name,
			Code:      country.Code,
			RegionId:  country.RegionID,
			CreatedAt: country.CreatedAt,
			UpdatedAt: country.UpdatedAt,
		}

		res = append(res, r)
	}

	functions.RespondwithJSON(w, http.StatusCreated, res)
}

func (p *PoolHandler) createCountry(w http.ResponseWriter, r *http.Request) {
	var req models.CreateCountryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		functions.RespondwithError(w, http.StatusBadRequest, "err in request body", err)
		return
	}

	if (req.Name == nil && *req.Name == "") || (req.Code == nil && *req.Code == "") || req.RegionId == nil {
		functions.RespondwithError(w, http.StatusBadRequest, "err in request body", fmt.Errorf("err in request body"))
		return
	}

	args := repository.AddCountryParams{
		Name:     *req.Name,
		Code:     *req.Code,
		RegionID: *req.RegionId,
	}

	country, err := p.Queries.AddCountry(r.Context(), args)
	if err != nil {
		functions.RespondwithError(w, http.StatusBadRequest, "server error", err)
		return
	}

	res := models.CreateCountryResponce{
		Id:        country.ID,
		Name:      country.Name,
		Code:      country.Code,
		RegionId:  country.RegionID,
		CreatedAt: country.CreatedAt,
		UpdatedAt: country.UpdatedAt,
	}

	functions.RespondwithJSON(w, http.StatusCreated, res)
}

func (p *PoolHandler) DeleteCountry(w http.ResponseWriter, r *http.Request) {
	var req models.DeleteCountryRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		functions.RespondwithError(w, http.StatusBadRequest, "err in request body", err)
		return
	}

	err := p.Queries.DeleteCountry(r.Context(), req.Name)
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

func (p *PoolHandler) getUpstreams(w http.ResponseWriter, r *http.Request) {

	upstreams, err := p.Queries.GetUpstreams(r.Context())
	if err != nil {
		functions.RespondwithError(w, http.StatusBadRequest, "server error", err)
		return
	}

	res := []models.GetUpstreamResponce{}

	for _, upstream := range upstreams {
		r := models.GetUpstreamResponce{
			Id:               upstream.ID,
			Tag:              upstream.Tag,
			UpstreamProvider: upstream.UpstreamProvider,
			Format:           upstream.Format,
			Domain:           upstream.Domain,
			Port:             int(upstream.Port),
			CreatedAt:        upstream.CreatedAt,
			UpdatedAt:        upstream.UpdatedAt,
		}

		res = append(res, r)
	}

	functions.RespondwithJSON(w, http.StatusCreated, res)
}

func (p *PoolHandler) createUpstream(w http.ResponseWriter, r *http.Request) {
	var req models.CreateUpstreamRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		functions.RespondwithError(w, http.StatusBadRequest, "err in request body", err)
		return
	}

	if (req.UpstreamProvider == nil && *req.UpstreamProvider == "") ||
		(req.Format == nil && *req.Format == "") ||
		(req.Tag == nil && *req.Tag == "") ||
		req.Port == nil ||
		(req.Domain == nil && *req.Domain == "") {
		functions.RespondwithError(w, http.StatusBadRequest, "err in request body", fmt.Errorf("err in request body"))
		return
	}

	args := repository.AddUpstreamParams{
		Tag:              *req.Tag,
		UpstreamProvider: *req.UpstreamProvider,
		Format:           *req.Format,
		Port:             int32(*req.Port),
		Domain:           *req.Domain,
	}

	upstream, err := p.Queries.AddUpstream(r.Context(), args)
	if err != nil {
		functions.RespondwithError(w, http.StatusBadRequest, "server error", err)
		return
	}

	res := models.CreateUpstreamResponce{
		Id:               upstream.ID,
		Tag:              upstream.Tag,
		UpstreamProvider: upstream.UpstreamProvider,
		Format:           upstream.Format,
		Port:             int(upstream.Port),
		Domain:           upstream.Domain,
		CreatedAt:        upstream.CreatedAt,
		UpdatedAt:        upstream.UpdatedAt,
	}

	functions.RespondwithJSON(w, http.StatusCreated, res)
}

func (p *PoolHandler) deleteUpstream(w http.ResponseWriter, r *http.Request) {
	var req models.DeleteUpstreamRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		functions.RespondwithError(w, http.StatusBadRequest, "err in request body", err)
		return
	}

	err := p.Queries.DeleteUpstream(r.Context(), req.Id)
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

func (p *PoolHandler) createPool(w http.ResponseWriter, r *http.Request) {

	//begin transaction
	ctx, err := p.DB.Begin()
	if err != nil {
		functions.RespondwithError(w, http.StatusInternalServerError, "failed to create user", err)
		return
	}
	defer func() {
		_ = ctx.Rollback()
	}()

	qtx := p.Queries.WithTx(ctx)

	var req models.CreatePoolRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		functions.RespondwithError(w, http.StatusBadRequest, "err in request body", err)
		return
	}

	if (req.Name == nil && *req.Name == "") || (req.Tag == nil && *req.Tag == "") || req.RegionId == nil || (req.Subdomain == nil && *req.Subdomain == "") || (req.Port == nil) || req.UpStreams == nil {
		functions.RespondwithError(w, http.StatusBadRequest, "err in request body", fmt.Errorf("err in request body"))
		return
	}

	args := repository.InsetPoolParams{
		Name:      *req.Name,
		Tag:       *req.Tag,
		RegionID:  *req.RegionId,
		Subdomain: *req.Subdomain,
		Port:      *req.Port,
	}

	pool, err := qtx.InsetPool(r.Context(), args)
	if err != nil {
		functions.RespondwithError(w, http.StatusBadRequest, "server error", err)
		return
	}

	weights := []int32{}
	upstreamTags := []string{}

	for _, upstreams := range *req.UpStreams {
		weights = append(weights, *upstreams.Weight)
		upstreamTags = append(upstreamTags, *upstreams.UpstreamTag)
	}

	weightArgs := repository.InsertPoolUpstreamWeightParams{
		PoolID:  pool.ID,
		Column2: weights,
		Column3: upstreamTags,
	}

	poolUpstreamWeights, err := qtx.InsertPoolUpstreamWeight(r.Context(), weightArgs)
	if err != nil {
		functions.RespondwithError(w, http.StatusBadRequest, "server error", err)
		return
	}

	if err := ctx.Commit(); err != nil {
		functions.RespondwithError(w, http.StatusInternalServerError, "failed to create pool", err)
		return
	}

	upstreamsRes := []models.CreateUpstreamWeightResponce{}

	for i, puw := range poolUpstreamWeights {
		upstreamRes := models.CreateUpstreamWeightResponce{
			UpstreamTag: upstreamTags[i],
			Weight:      puw.Weight,
		}
		upstreamsRes = append(upstreamsRes, upstreamRes)
	}
	res := models.CreatePoolResponce{
		Id:        pool.ID,
		Name:      &pool.Name,
		Tag:       &pool.Tag,
		RegionId:  &pool.RegionID,
		Subdomain: &pool.Subdomain,
		Port:      &pool.Port,
		UpStreams: &upstreamsRes,
		CreatedAt: pool.CreatedAt,
		UpdatedAt: pool.UpdatedAt,
	}

	functions.RespondwithJSON(w, http.StatusCreated, res)

}

func (p *PoolHandler) getPools(w http.ResponseWriter, r *http.Request) {
	rows, err := p.Queries.ListPoolsWithUpstreams(r.Context())
	if err != nil {
		functions.RespondwithError(w, http.StatusInternalServerError, "Failed to fetch pools", err)
		return
	}

	poolMap := make(map[uuid.UUID]*models.GetPoolsResponse)

	// Preserve order
	var orderedPools []*models.GetPoolsResponse

	for _, row := range rows {
		pool, exists := poolMap[row.PoolID]
		if !exists {
			pool = &models.GetPoolsResponse{
				Id:        row.PoolID,
				Name:      row.PoolName,
				Tag:       row.PoolTag,
				Subdomain: row.PoolSubdomain,
				Port:      row.PoolPort,
				Upstreams: []models.PoolUpstream{},
			}
			poolMap[row.PoolID] = pool
			orderedPools = append(orderedPools, pool)
		}

		if row.UpstreamTag.Valid {
			pool.Upstreams = append(pool.Upstreams, models.PoolUpstream{
				Tag:    row.UpstreamTag.String,
				Format: row.UpstreamFormat.String,
				Port:   row.UpstreamPort.Int32,
				Domain: row.UpstreamDomain.String,
			})
		}
	}

	// Create final response from ordered list
	response := make([]models.GetPoolsResponse, 0, len(orderedPools))
	for _, pool := range orderedPools {
		response = append(response, *pool)
	}

	functions.RespondwithJSON(w, http.StatusOK, response)
}
