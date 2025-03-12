package impl

import (
	"auction_service_api/internal/models"
	"auction_service_api/internal/repositories"
	"auction_service_api/internal/services"
	"context"

	"github.com/google/uuid"
)

type auctionSrviceImpl struct {
	repo repositories.AuctionRepository
}

func NewAuctionService(repo repositories.AuctionRepository) services.AuctionService {
	return &auctionSrviceImpl{repo}
}

func (asi *auctionSrviceImpl) GetAll(ctx context.Context) []models.AuctionModel {
	var auctionModel []models.AuctionModel
	_ = auctionModel
	asi.repo.GetAll(ctx)
	return nil
}
func (as *auctionSrviceImpl) GetItem(ctx context.Context, id uuid.UUID) (*models.AuctionModel, error) {
	return nil, nil
}
func (as *auctionSrviceImpl) Create(ctx context.Context, auction models.CreateAuctionModel) error {
	return nil
}
func (as *auctionSrviceImpl) Update(ctx context.Context, auction models.UpdateAuctionModel) error {
	return nil
}
func (as *auctionSrviceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}
