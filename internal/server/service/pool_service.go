package service

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/torchlabssoftware/subnetwork_system/internal/db/repository"
	models "github.com/torchlabssoftware/subnetwork_system/internal/server/models"
)

type PoolService interface {
	GetRegions(ctx context.Context) ([]models.GetRegionResponce, int, string, error)
	CreateRegion(ctx context.Context, req models.CreateRegionRequest) (models.CreateRegionResponce, int, string, error)
	DeleteRegion(ctx context.Context, name string) (int, string, error)
	GetCountries(ctx context.Context) ([]models.GetCountryResponce, int, string, error)
	CreateCountry(ctx context.Context, req models.CreateCountryRequest) (models.CreateCountryResponce, int, string, error)
	DeleteCountry(ctx context.Context, name string) (int, string, error)
	GetUpstreams(ctx context.Context) ([]models.GetUpstreamResponce, int, string, error)
	CreateUpstream(ctx context.Context, req models.CreateUpstreamRequest) (models.CreateUpstreamResponce, int, string, error)
	DeleteUpstream(ctx context.Context, tag string) (int, string, error)
}

type PoolServiceImpl struct {
	Queries *repository.Queries
}

func NewPoolService(queries *repository.Queries) PoolService {
	return &PoolServiceImpl{
		Queries: queries,
	}
}

func (s *PoolServiceImpl) GetRegions(ctx context.Context) ([]models.GetRegionResponce, int, string, error) {
	regions, err := s.Queries.GetRegions(ctx)
	if err != nil {
		return nil, http.StatusInternalServerError, "failed to get regions", err
	}

	res := []models.GetRegionResponce{}

	for _, region := range regions {
		r := models.GetRegionResponce{
			Id:        region.ID,
			Name:      region.Name,
			CreatedAt: region.CreatedAt,
		}

		res = append(res, r)
	}

	return res, http.StatusOK, "", nil
}

func (s *PoolServiceImpl) CreateRegion(ctx context.Context, req models.CreateRegionRequest) (models.CreateRegionResponce, int, string, error) {
	region, err := s.Queries.AddRegion(ctx, *req.Name)
	if err != nil {
		return models.CreateRegionResponce{}, http.StatusInternalServerError, "failed to create region", err
	}

	res := models.CreateRegionResponce{
		Id:        region.ID,
		Name:      region.Name,
		CreatedAt: region.CreatedAt,
	}

	return res, http.StatusCreated, "region created", nil
}

func (s *PoolServiceImpl) DeleteRegion(ctx context.Context, name string) (int, string, error) {
	if err := s.Queries.DeleteRegion(ctx, name); err != nil {
		return http.StatusInternalServerError, "failed to delete region", err
	}
	return http.StatusOK, "region deleted", nil
}

func (s *PoolServiceImpl) GetCountries(ctx context.Context) ([]models.GetCountryResponce, int, string, error) {
	countries, err := s.Queries.GetCountries(ctx)
	if err != nil {
		return nil, http.StatusInternalServerError, "failed to get countries", err
	}

	res := []models.GetCountryResponce{}

	for _, country := range countries {
		r := models.GetCountryResponce{
			Id:        country.ID,
			Name:      country.Name,
			Code:      country.Code,
			RegionId:  country.RegionID,
			CreatedAt: country.CreatedAt,
		}

		res = append(res, r)
	}

	return res, http.StatusOK, "", nil
}

func (s *PoolServiceImpl) CreateCountry(ctx context.Context, req models.CreateCountryRequest) (models.CreateCountryResponce, int, string, error) {
	args := repository.AddCountryParams{
		Name:     *req.Name,
		Code:     *req.Code,
		RegionID: *req.RegionId,
	}

	country, err := s.Queries.AddCountry(ctx, args)
	if err != nil {
		return models.CreateCountryResponce{}, http.StatusInternalServerError, "failed to create country", err
	}

	res := models.CreateCountryResponce{
		Id:        country.ID,
		Name:      country.Name,
		Code:      country.Code,
		RegionId:  country.RegionID,
		CreatedAt: country.CreatedAt,
	}

	return res, http.StatusCreated, "country created", nil
}

func (s *PoolServiceImpl) DeleteCountry(ctx context.Context, name string) (int, string, error) {
	if err := s.Queries.DeleteCountry(ctx, name); err != nil {
		return http.StatusInternalServerError, "failed to delete country", err
	}
	return http.StatusOK, "country deleted", nil
}

func (s *PoolServiceImpl) GetUpstreams(ctx context.Context) ([]models.GetUpstreamResponce, int, string, error) {
	upstreams, err := s.Queries.GetUpstreams(ctx)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, http.StatusNotFound, "upstreams not found", nil
		}
		return nil, http.StatusInternalServerError, "failed to get upstreams", err
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
		}

		res = append(res, r)
	}

	return res, http.StatusOK, "", nil
}

func (s *PoolServiceImpl) CreateUpstream(ctx context.Context, req models.CreateUpstreamRequest) (models.CreateUpstreamResponce, int, string, error) {

	args := repository.AddUpstreamParams{
		Tag:              *req.Tag,
		UpstreamProvider: *req.UpstreamProvider,
		Format:           *req.Format,
		Port:             int32(*req.Port),
		Domain:           *req.Domain,
	}

	upstream, err := s.Queries.AddUpstream(ctx, args)
	if err != nil {
		return models.CreateUpstreamResponce{}, http.StatusInternalServerError, "failed to create upstream", err
	}

	res := models.CreateUpstreamResponce{
		Id:               upstream.ID,
		Tag:              upstream.Tag,
		UpstreamProvider: upstream.UpstreamProvider,
		Format:           upstream.Format,
		Port:             int(upstream.Port),
		Domain:           upstream.Domain,
		CreatedAt:        upstream.CreatedAt,
	}

	return res, http.StatusCreated, "upstream created", nil
}

func (s *PoolServiceImpl) DeleteUpstream(ctx context.Context, tag string) (int, string, error) {
	if err := s.Queries.DeleteUpstreamByTag(ctx, tag); err != nil {
		return http.StatusInternalServerError, "failed to delete upstream", err
	}
	return http.StatusOK, "upstream deleted", nil
}
