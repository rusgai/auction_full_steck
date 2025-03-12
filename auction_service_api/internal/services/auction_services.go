package services

import (
	"auction_service_api/internal/models"
	"context"

	"github.com/google/uuid"
)

type AuctionService interface {
	GetAll(ctx context.Context) []models.AuctionModel
	GetItem(ctx context.Context, id uuid.UUID) (*models.AuctionModel, error)
	Create(ctx context.Context, auction models.CreateAuctionModel) error
	Update(ctx context.Context, auction models.UpdateAuctionModel) error
	Delete(ctx context.Context, id uuid.UUID) error
}
