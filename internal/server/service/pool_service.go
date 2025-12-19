package service

import (
	"context"
	"net/http"

	"github.com/torchlabssoftware/subnetwork_system/internal/db/repository"
	models "github.com/torchlabssoftware/subnetwork_system/internal/server/models"
)

type PoolService interface {
	GetRegions(ctx context.Context) ([]models.GetRegionResponce, int, string, error)
	CreateRegion(ctx context.Context, req models.CreateRegionRequest) (models.CreateRegionResponce, int, string, error)
	DeleteRegion(ctx context.Context, name string) (int, string, error)
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
